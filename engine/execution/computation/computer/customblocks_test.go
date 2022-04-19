package computer_test

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go/crypto"
	"github.com/onflow/flow-go/crypto/hash"
	"github.com/onflow/flow-go/engine/execution"
	"github.com/onflow/flow-go/engine/execution/computation/committer"
	"github.com/onflow/flow-go/engine/execution/computation/computer"
	executionState "github.com/onflow/flow-go/engine/execution/state"
	bootstrapexec "github.com/onflow/flow-go/engine/execution/state/bootstrap"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/engine/execution/testutil"
	"github.com/onflow/flow-go/fvm/programs"
	completeLedger "github.com/onflow/flow-go/ledger/complete"
	"github.com/onflow/flow-go/ledger/complete/wal/fixtures"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"testing"

	"github.com/onflow/flow-go/fvm"
	"github.com/rs/zerolog"
)

const (
	// AddressLength is the size of an account address in bytes.
	// (n) is the size of an account address in bits.
	AddressLength = (64 + 7) >> 3
	// addressIndexLength is the size of an account address state in bytes.
	// (k) is the size of an account address in bits.
	addressIndexLength = (45 + 7) >> 3
	maxIndex           = (1 << 45) - 1
)

//var EmptyAddress = flow.Address{}

type vmTest struct {
	bootstrapOptions []fvm.BootstrapProcedureOption
	contextOptions   []fvm.Option
}

func newVMTest() vmTest {
	return vmTest{}
}

func (vmt vmTest) withBootstrapProcedureOptions(opts ...fvm.BootstrapProcedureOption) vmTest {
	vmt.bootstrapOptions = append(vmt.bootstrapOptions, opts...)
	return vmt
}

func (vmt vmTest) withContextOptions(opts ...fvm.Option) vmTest {
	vmt.contextOptions = append(vmt.contextOptions, opts...)
	return vmt
}
func TestPrograms(t *testing.T) {
	t.Run("CustomBlock Testing w/ no mocks",
		newVMTest().run(
			func(t *testing.T, chain flow.Chain) {
			},
		),
	)
}

func (vmt vmTest) run(
	f func(t *testing.T, chain flow.Chain),
) func(t *testing.T) {
	return func(t *testing.T) {
		// preliminaries
		//numTxPerCol := 2
		//numCol := 2
		numAccount := 4
		seeds := []uint64{1, 2, 3, 4}
		/*
			//getSignerReceiver := func(chain flow.Chain) *flow.TransactionBody {
			//	return flow.NewTransactionBody().SetScript([]byte(fmt.Sprintf(`
			//			import FungibleToken from 0x%s
			//			import FlowToken from 0x%s
			//
			//			transaction(recipient: Address) {
			//				prepare(signer: AuthAccount) {
			//					let vaultRef = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVaultUnittest0)
			//						?? panic("failed to borrow reference to sender vault")
			//				}
			//				post {
			//					getAccount(recipient)
			//						.getCapability<&FlowToken.Vault{FungibleToken.Receiver}>(/public/flowTokenReceiverUnittest0)
			//						.check():
			//						"Vault recv ref was not created correctly."
			//				}
			//			}`, fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain))))
			//}
		*/

		// execution objects
		chain := flow.Mainnet.Chain()
		baseOpts := []fvm.Option{fvm.WithChain(chain)}
		opts := append(baseOpts, vmt.contextOptions...)
		// custom address generator used to deterministically generate account addresses from a seed (instead of purely randomly)
		cag := new(CustomAddressGenerator)

		// TODO: I think the issue is in the account creation and setup.
		// this should return the address of newly created accounts using the CustomAddressGenerator and seeds
		accountAddresses, err := cag.AddressAtIndexes(seeds)
		require.NoError(t, err)

		// Account creation requires private keys.
		// Generate private keys and the appropriate flow transactions for creation.
		privateKeys, createAccountTxs := createAccountCreationTransactions(t, chain, numAccount, seeds)

		// account creation transactions need to be signed by the chain service account
		// TODO: should these funding transactions be signed by the account owner?
		signTransactionsAsServiceAccount(createAccountTxs, 0, chain)
		require.NoError(t, err)

		// initialize a new Fund Account flow transaction, fund with 10 tokens
		// transaction arguments : (amount: UFix64, recipient: Address)
		initialAmount := 10
		recipient := accountAddresses[0]
		fundAccountTx := flow.NewTransactionBody().
			SetScript(getFundAccountTransactionScript(chain)).
			AddAuthorizer(chain.ServiceAddress()).
			AddArgument(jsoncdc.MustEncode(cadence.UFix64(initialAmount))).
			AddArgument(jsoncdc.MustEncode(cadence.NewAddress(recipient))).
			// set the proposal key and sequence number for this transaction:
			SetProposalKey(chain.ServiceAddress(), 0, uint64(len(createAccountTxs))).
			// service account is the payer
			SetPayer(chain.ServiceAddress())
		// sign the tx envelope
		err = testutil.SignEnvelope(
			fundAccountTx,
			chain.ServiceAddress(),
			unittest.ServiceAccountPrivateKey,
		)
		require.NoError(t, err)

		// initialize a Transfer From Account transaction, transfer of 4 tokens from address 0 to address 1
		// transaction arguments: (amount: UFix64, to: Address)
		amount := 4
		to := accountAddresses[1]
		transferFromAccountTx := flow.NewTransactionBody().
			SetScript(getTransferTokenTransactionScript(chain)).
			// authorized by address 0
			AddAuthorizer(accountAddresses[0]).
			AddArgument(jsoncdc.MustEncode(cadence.UFix64(amount))).
			AddArgument(jsoncdc.MustEncode(cadence.NewAddress(to))).
			SetProposalKey(accountAddresses[0], 0, uint64(len(createAccountTxs)+1)).
			// paid by address 0
			SetPayer(accountAddresses[0])

		err = testutil.SignEnvelope(
			transferFromAccountTx,
			accountAddresses[0],
			privateKeys[0],
		)
		require.NoError(t, err)

		// TODO: FlowEpoch contract missing is one of the errors
		//epochConfig := epochs.DefaultEpochConfig()
		//epochConfig.NumCollectorClusters = 0
		bootstrpOpts := []fvm.BootstrapProcedureOption{
			fvm.WithInitialTokenSupply(unittest.GenesisTokenSupply),
			fvm.WithAccountCreationFee(fvm.DefaultAccountCreationFee),
			fvm.WithMinimumStorageReservation(fvm.DefaultMinimumStorageReservation),
			fvm.WithTransactionFee(fvm.DefaultTransactionFees),
			fvm.WithStorageMBPerFLOW(fvm.DefaultStorageMBPerFLOW),
			//fvm.WithEpochConfig(epochConfig),
		}

		collectionResult := executeBlockAndNotVerify(t, [][]*flow.TransactionBody{
			{&createAccountTxs[0]},
			{&createAccountTxs[1]},
			{&createAccountTxs[2]},
			{&createAccountTxs[3]},
			{fundAccountTx},
		}, chain, opts, bootstrpOpts)

		fmt.Sprint(collectionResult.ComputationUsed)
		f(t, chain)

	}
}

func createAccountCreationTransactions(t *testing.T, chain flow.Chain, numberOfPrivateKeys int, seedArr []uint64) ([]flow.AccountPrivateKey, []flow.TransactionBody) {
	accountKeys, err := generateAccountPrivateKeys(numberOfPrivateKeys, seedArr)
	require.NoError(t, err)

	var txs []flow.TransactionBody
	for i := 0; i < len(accountKeys); i++ {
		// create the transaction to create the account
		tx := flow.NewTransactionBody().
			SetScript(getSetupAccountTransactionScript(chain)).
			AddAuthorizer(chain.ServiceAddress())

		txs = append(txs, *tx)
	}
	return accountKeys, txs
}

// generateAccountPrivateKeys generates a number of private keys.
func generateAccountPrivateKeys(numberOfPrivateKeys int, seedArr []uint64) ([]flow.AccountPrivateKey, error) {
	var privateKeys []flow.AccountPrivateKey
	for i := 0; i < numberOfPrivateKeys; i++ {
		pk, err := generateAccountPrivateKey(seedArr[i])
		if err != nil {
			return nil, err
		}
		privateKeys = append(privateKeys, pk)
	}

	return privateKeys, nil
}

// generateAccountPrivateKey generates a private key.
func generateAccountPrivateKey(seedInt uint64) (flow.AccountPrivateKey, error) {
	seed := uint64ToPrivKeyBytes(seedInt)
	privateKey, err := crypto.GeneratePrivateKey(crypto.ECDSAP256, seed)
	if err != nil {
		return flow.AccountPrivateKey{}, err
	}
	pk := flow.AccountPrivateKey{
		PrivateKey: privateKey,
		SignAlgo:   crypto.ECDSAP256,
		HashAlgo:   hash.SHA2_256,
	}
	return pk, nil
}

func signTransactionsAsServiceAccount(txs []flow.TransactionBody, seqNum uint64, chain flow.Chain) {
	for i := 0; i < len(txs); i++ {
		txs[i].SetProposalKey(chain.ServiceAddress(), 0, seqNum)
		txs[i].SetPayer(chain.ServiceAddress())
		err := testutil.SignEnvelope(&txs[i], chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
		if err != nil {
			panic(fmt.Errorf("cannot sign envelope: %w", err))
		}
		seqNum++
	}
}

func executeBlockAndNotVerify(t *testing.T,
	txs [][]*flow.TransactionBody,
	chain flow.Chain,
	opts []fvm.Option,
	bootstrapOpts []fvm.BootstrapProcedureOption) *execution.ComputationResult {
	rt := fvm.NewInterpreterRuntime()
	vm := fvm.NewVirtualMachine(rt)

	logger := zerolog.New(log.Writer())
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)

	opts = append(opts, fvm.WithChain(chain))

	fvmContext :=
		fvm.NewContext(
			logger,
			opts...,
		)

	collector := metrics.NewNoopCollector()
	tracer := trace.NewNoopTracer()

	wal := &fixtures.NoopWAL{}

	ledger, err := completeLedger.NewLedger(wal, 100, collector, logger, completeLedger.DefaultPathFinderVersion)
	require.NoError(t, err)

	bootstrapper := bootstrapexec.NewBootstrapper(logger)

	initialCommit, err := bootstrapper.BootstrapLedger(
		ledger,
		unittest.ServiceAccountPublicKey,
		chain,
		bootstrapOpts...,
	)

	require.NoError(t, err)

	ledgerCommiter := committer.NewLedgerViewCommitter(ledger, tracer)

	blockComputer, err := computer.NewBlockComputer(vm, fvmContext, collector, tracer, logger, ledgerCommiter)
	require.NoError(t, err)

	view := delta.NewView(executionState.LedgerGetRegister(ledger, initialCommit))

	executableBlock := unittest.ExecutableBlockFromTransactions(txs)
	executableBlock.StartState = &initialCommit

	computationResult, err := blockComputer.ExecuteBlock(context.Background(), executableBlock, view, programs.NewEmptyPrograms())
	require.NoError(t, err)
	return computationResult
}

// ********************************************
// Custom Address Generator Struct and Methods
// ********************************************
type CustomAddressGenerator struct {
	index   uint64
	curSeed uint64
	addrArr []flow.Address
	seeds   []uint64
}

func (e *CustomAddressGenerator) init(size uint64, seeds []uint64) {
	// init addr generator
	if uint64(len(seeds)) > 0 && uint64(len(seeds)) < size {
		panic(fmt.Errorf("size mismatch: with seed slice length, size: %x, seeds_len: %x", size, len(seeds)))
	}
	e.index = 0
	e.curSeed = 0
	e.addrArr = make([]flow.Address, size)
	e.seeds = seeds
	e.pushRandAddr(size)
}

func (e *CustomAddressGenerator) pushRandAddr(size uint64) {
	if len(e.seeds) == 0 {
		for i := 0; i < int(size); i++ {
			r := uint64(rand.Intn(maxIndex)) + 1
			custAddr := e.initAtIndex(r)
			addr := custAddr.CurrentAddress()
			e.addrArr[i] = addr
			//fmt.Printf("generated address: %+v\n", addr)
		}
	} else {
		for i := 0; i < int(size); i++ {
			cag := e.initAtIndex(e.seeds[i])
			addr := cag.CurrentAddress()
			e.addrArr[i] = addr
			//fmt.Printf("generated address: %+v\n", addr)
		}
	}
}

func (e *CustomAddressGenerator) initAtIndex(rand uint64) CustomAddressGenerator {
	e.curSeed = rand
	return *e
}

func (e *CustomAddressGenerator) NextAddress() (flow.Address, error) {
	if e.index == 0 {
		e.index++
		return e.addrArr[e.index-1], nil
	}
	if (e.index) > maxIndex {
		return flow.Address{}, fmt.Errorf("the new index value is not valid, it must be less or equal to %x", maxIndex)
	}
	e.index++
	return e.addrArr[e.index-1], nil
}

func (e *CustomAddressGenerator) CurrentAddress() flow.Address {
	return uint64ToAddress(e.curSeed)
}

func (e *CustomAddressGenerator) Bytes() []byte {
	panic("not implemented")
}

func (e *CustomAddressGenerator) AddressCount() uint64 {
	panic("not implemented")
}

func (e *CustomAddressGenerator) AddressAtIndexes(seeds []uint64) ([]flow.Address, error) {
	max := findSlicesMax(seeds)
	if max > uint64(maxIndex) {
		return make([]flow.Address, len(seeds)), fmt.Errorf("index must be less or equal to %x", maxIndex)
	}
	var addresses []flow.Address
	for i := 0; i < len(seeds); i++ {
		cust := e.initAtIndex(seeds[i])
		addresses = append(addresses, cust.CurrentAddress())
	}
	return addresses, nil
}

func uint64ToAddress(v uint64) flow.Address {
	var b [AddressLength]byte
	binary.BigEndian.PutUint64(b[:], v)
	return b
}

func findSlicesMax(arr []uint64) uint64 {
	var biggest uint64
	if len(arr) == 0 {
		fmt.Println("Array must be non empty")
	} else {
		biggest = arr[0]
		for _, v := range arr {
			if v > biggest {
				biggest = v
			}
		}
	}
	return biggest
}

func uint64ToPrivKeyBytes(v uint64) []byte {
	var b [crypto.KeyGenSeedMinLenECDSAP256]byte
	binary.BigEndian.PutUint64(b[:], v)
	return b[:]
}

// ---------------------------------
// Script template functions
// ----------------------------------
func getTransferTokenTransactionScript(chain flow.Chain) []byte {
	// transfer token script
	return []byte(fmt.Sprintf(transferTokenTemplate, fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
}

func getFundAccountTransactionScript(chain flow.Chain) []byte {
	txScript := []byte(fmt.Sprintf(fundAccountTemplate, fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
	return txScript
}

func getSetupAccountTransactionScript(chain flow.Chain) []byte {
	return []byte(fmt.Sprintf(setupAccountTemplate, fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
}

const setupAccountTemplate = `
// This transaction is a template for a transaction
// to add a Vault resource to their account
// so that they can use the flowToken

import FungibleToken from 0x%s
import FlowToken from 0x%s

transaction {

    prepare(signer: AuthAccount) {

        if signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault) == nil {
            // Create a new flowToken Vault and put it in storage
            signer.save(<-FlowToken.createEmptyVault(), to: /storage/flowTokenVault)

            // Create a public capability to the Vault that only exposes
            // the deposit function through the Receiver interface
            signer.link<&FlowToken.Vault{FungibleToken.Receiver}>(
                /public/flowTokenReceiver,
                target: /storage/flowTokenVault
            )

            // Create a public capability to the Vault that only exposes
            // the balance field through the Balance interface
            signer.link<&FlowToken.Vault{FungibleToken.Balance}>(
                /public/flowTokenBalance,
                target: /storage/flowTokenVault
            )
        }
    }

}
`

// TODO: put post conditions in to check that capability was created properly, for testing?
//post {
//	// Check that the capabilities were created correctly
//	// by getting the public capability and checking
//	// that it points to a valid Vault object
//	// that implements the Receiver interface
//	getAccount(0x%s).getCapability<&FlowToken.Vault{FungibleToken.Receiver}>(/public/flowTokenReceiver)
//					.check():
//					"Vault Receiver Reference was not created correctly"
//}

// TODO: I think the accounts need to be created and public keys registered?
const accountCreationTemplate = `
transaction(publicKey: [UInt8]) {
	prepare(signer: AuthAccount) {
		let acct = AuthAccount(payer: signer)
		acct.addPublicKey(publicKey)
	}
}
`

const transferTokenTemplate = `
// This transaction is a template for a transaction that
// could be used by anyone to send tokens to another account
// that has been set up to receive tokens.
//
// The withdraw amount and the account from getAccount
// would be the parameters to the transaction

import FungibleToken from 0x%s
import FlowToken from 0x%s

transaction(amount: UFix64, to: Address) {

	// The Vault resource that holds the tokens that are being transferred
	let sentVault: @FungibleToken.Vault

	prepare(signer: AuthAccount) {

		// Get a reference to the signer's stored vault
		let vaultRef = signer.borrow<&FungibleToken.Provider>(from: /storage/flowTokenVault)
			?? panic("Could not borrow reference to the owner's Vault!")

		// Withdraw tokens from the signer's stored vault
		self.sentVault <- vaultRef.withdraw(amount: amount)
	}

	execute {

		// Get the recipient's public account object
		let recipient = getAccount(to)

		// Get a reference to the recipient's Receiver
		let receiverRef = recipient.getCapability(/public/flowTokenReceiver)
			.borrow<&{FungibleToken.Receiver}>()
			?? panic("Could not borrow receiver reference to the recipient's Vault")

		// Deposit the withdrawn tokens in the recipient's receiver
		receiverRef.deposit(from: <-self.sentVault)
	}
}`

/* TODO: this is the error we are trying to solve
error: panic: failed to borrow reference to recipient vault
--> a53a42cdbab6a80bebbea0b643d80e75a9f74f2011fe4f2f40924fd2929d3bee:16:5
|
16 | 		?? panic("failed to borrow reference to recipient vault")
|      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^*/

const fundAccountTemplate = `
import FungibleToken from 0x%s
import FlowToken from 0x%s

transaction(amount: UFix64, recipient: Address) {
	let sentVault: @FungibleToken.Vault
	prepare(signer: AuthAccount) {
	let vaultRef = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)
		?? panic("failed to borrow reference to sender vault")
	self.sentVault <- vaultRef.withdraw(amount: amount)
	}
	execute {
	let receiverRef =  getAccount(recipient)
		.getCapability(/public/flowTokenReceiver)
		.borrow<&{FungibleToken.Receiver}>()
		?? panic("failed to borrow reference to recipient vault")
	receiverRef.deposit(from: <-self.sentVault)
	}
}
`

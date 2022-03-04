package computer_test

import (
	"bufio"
	"context"
	"encoding/binary"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	"github.com/onflow/flow-go/module/mempool/entity"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"reflect"
	"strings"
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

var EmptyAddress = flow.Address{}

type CustomAddressGenerator struct {
	index   uint64
	curSeed uint64
	addrArr []flow.Address
	seeds   []uint64
}

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

func (vmt vmTest) run(
	f func(t *testing.T, chain flow.Chain),
) func(t *testing.T) {
	return func(t *testing.T) {
		numTxPerCol := 2
		numCol := 2
		seeds := []uint64{1, 2, 3, 4}

		chain := flow.Mainnet.Chain()
		baseOpts := []fvm.Option{
			fvm.WithChain(chain),
		}
		opts := append(baseOpts, vmt.contextOptions...)

		//transferTokensTx := func(chain flow.Chain) *flow.TransactionBody {
		//	return flow.NewTransactionBody().
		//		SetScript([]byte(fmt.Sprintf(`
		//					// This transaction is a template for a transaction that
		//					// could be used by anyone to send tokens to another account
		//					// that has been set up to receive tokens.
		//					//
		//					// The withdraw amount and the account from getAccount
		//					// would be the parameters to the transaction
		//
		//					import FungibleToken from 0x%s
		//					import FlowToken from 0x%s
		//
		//					transaction(amount: UFix64, to: Address) {
		//
		//						// The Vault resource that holds the tokens that are being transferred
		//						let sentVault: @FungibleToken.Vault
		//
		//						prepare(signer: AuthAccount) {
		//
		//							// Get a reference to the signer's stored vault
		//							let vaultRef = signer.borrow<&FlowToken.Vault>(from: /storage/flowTokenVault)
		//								?? panic("Could not borrow reference to the owner's Vault!")
		//
		//							// Withdraw tokens from the signer's stored vault
		//							self.sentVault <- vaultRef.withdraw(amount: amount)
		//						}
		//
		//						execute {
		//
		//							// Get the recipient's public account object
		//							let recipient = getAccount(to)
		//
		//							// Get a reference to the recipient's Receiver
		//							let receiverRef = recipient.getCapability(/public/flowTokenReceiver)
		//								.borrow<&{FungibleToken.Receiver}>()
		//								?? panic("Could not borrow receiver reference to the recipient's Vault")
		//
		//							// Deposit the withdrawn tokens in the recipient's receiver
		//							receiverRef.deposit(from: <-self.sentVault)
		//						}
		//					}`, fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain))),
		//		)
		//}

		// ==== Create an account ====
		privateKeys, createAccountTxs := CreateAccountCreationTransactions(t, chain, numCol*numTxPerCol, seeds)
		fmt.Sprint(len(privateKeys))
		// this should return the address of newly created account
		//address, err := chain.AddressAtIndex(5)
		//require.NoError(t, err)

		SignTransactionAsServiceAccounts(createAccountTxs, 0, chain)
		//require.NoError(t, err)

		//// ==== Transfer tokens to new account ====
		//transferTx := transferTokensTx(chain).
		//	AddAuthorizer(chain.ServiceAddress()).
		//	AddArgument(jsoncdc.MustEncode(cadence.UFix64(1000))).
		//	AddArgument(jsoncdc.MustEncode(cadence.NewAddress(address)))
		//
		//transferTx.SetProposalKey(chain.ServiceAddress(), 0, 1)
		//transferTx.SetPayer(chain.ServiceAddress())
		//
		//err = testutil.SignEnvelope(
		//	transferTx,
		//	chain.ServiceAddress(),
		//	unittest.ServiceAccountPrivateKey,
		//)
		//require.NoError(t, err)
		//
		//// ==== Transfer tokens from new account ====
		//
		//transferTx2 := transferTokensTx(chain).
		//	AddAuthorizer(address).
		//	AddArgument(jsoncdc.MustEncode(cadence.UFix64(tc.tryToTransfer))).
		//	AddArgument(jsoncdc.MustEncode(cadence.NewAddress(chain.ServiceAddress())))
		//
		//transferTx2.SetProposalKey(address, 0, 0)
		//transferTx2.SetPayer(address)
		//
		//err = testutil.SignEnvelope(
		//	transferTx2,
		//	address,
		//	privateKey,
		//)
		//require.NoError(t, err)

		bootstrpOpts := []fvm.BootstrapProcedureOption{
			fvm.WithInitialTokenSupply(unittest.GenesisTokenSupply),
			fvm.WithTransactionFee(fvm.DefaultTransactionFees)}

		cr := executeBlockAndNotVerify(t, [][]*flow.TransactionBody{
			{&createAccountTxs[0]},
			{&createAccountTxs[1]},
			{&createAccountTxs[2]},
			{&createAccountTxs[3]},
		}, chain, opts, bootstrpOpts)

		fmt.Sprint(cr.ComputationUsed)

		//logFilename := "customBlockTest.log"
		//csvFilename := "customBlockLogOutput.csv"
		//file, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		//logger := zerolog.New(file)
		//// NewTracer takes a logger, service name (string), a chainID (string), and a trace sensitivity (int))
		//tracer, err := trace.NewTracer(logger, "CustomBlockTrace", "test", trace.SensitivityCaptureAll)
		//// NewExecutionCollector generates a metrics object, taking a tracer and a Registerer object as input.
		//metrics := metrics.NewExecutionCollector(tracer, prometheus.DefaultRegisterer)
		//
		//bootstrapper := bootstrapexec.NewBootstrapper(logger)
		//
		//privateKeys, err := GenerateAccountPrivateKeys(numCol*numTxPerCol, seeds)
		//require.NoError(t, err)
		//
		//ledger, err := completeLedger.NewLedger(wal, 100, metrics, logger, completeLedger.DefaultPathFinderVersion)
		//require.NoError(t, err)
		//ledgerCommiter := committer.NewLedgerViewCommitter(ledger, tracer)
		//
		//bootstrpOpts := []fvm.BootstrapProcedureOption{
		//	fvm.WithInitialTokenSupply(unittest.GenesisTokenSupply),
		//	fvm.WithTransactionFee(fvm.DefaultTransactionFees)}
		//
		//initialCommit, err := bootstrapper.BootstrapLedger(
		//	ledger,
		//	unittest.ServiceAccountPublicKey,
		//	chain,
		//	bootstrpOpts...
		//)
		//
		//// todo: Unmock ledger once we have valid transactions
		//// code snippet to unmock ledger
		////wal := &fixtures.NoopWAL{}
		////ledger, err := complete.NewLedger(wal, 100, &metrics.NoopCollector{}, zerolog.Logger{}, complete.DefaultPathFinderVersion)
		////curS := ledger.InitialState()
		////q := utils.QueryFixture()
		////q.SetState(curS)
		////require.NoError(t, err)
		////retProof, err := ledger.Prove(q)
		////require.NoError(t, err)
		////trieProof, err := encoding.DecodeTrieBatchProof(retProof)
		////assert.True(t, proof.VerifyTrieBatchProof(trieProof, curS))
		//
		////// unmocking the ledger requires to provide proofs of state commitment.
		////// However dummy transactions w/o signatures do not yield valid proofs.
		////ledger := new(ledgermock.Ledger)
		////var expectedStateCommitment led.State
		////copy(expectedStateCommitment[:], []byte{1, 2, 3})
		////ledger.On("Set", mock.Anything).
		////	Return(expectedStateCommitment, nil, nil).
		////	Times(numTxPerCol*numCol + 1)
		////expectedProof := led.Proof([]byte{2, 3, 4})
		////ledger.On("Prove", mock.Anything).
		////	Return(expectedProof, nil).
		////	Times(numTxPerCol*numCol + 1)
		//
		////blockcommitter := committer.NewLedgerViewCommitter(ledger, tracer)
		//
		//exe, err := computer.NewBlockComputer(vm, execCtx, metrics, tracer, logger, committer.NewNoopViewCommitter())
		//require.NoError(t, err)
		//
		//// Generate your own block with 2 collections and 4 txs in total
		////block := generateBlock(numCol, numTxPerCol, rag)
		////customAddr := new(CustomAddressGenerator)
		////customAddr.init(uint64(numCol*numTxPerCol), seeds)
		//block := generateCustomBlock(numCol, numTxPerCol, accounts, privateKeys, chain)
		//
		//// returns nill register value
		//view := delta.NewView(func(owner, controller, key string) (flow.RegisterValue, error) {
		//	return nil, nil
		//})
		//
		//result, err := exe.ExecuteBlock(context.Background(), block, view, programs.NewEmptyPrograms())
		//assert.NoError(t, err)
		//assert.Len(t, result.StateSnapshots, numCol+1)
		//assert.Len(t, result.TrieUpdates, numCol+1)
		//
		//assertEventHashesMatch(t, numCol+1, result)
		////vm.AssertExpectations(t)
		//
		//// open file containing logs in JSON format
		//sourceFile, err := os.Open(logFilename)
		//if err != nil {
		//	// do nothing?
		//}
		//// create csv file
		//outputFile, err := os.Create(csvFilename)
		//if err != nil {
		//	// do nothing?
		//}
		//// convert the JSON logs to CSV file
		//lineswritten, err := convertJSONToCSV(sourceFile, outputFile)
		//if err != nil {
		//	// do nothing
		//}
		//if lineswritten == 0 {
		//	panic("Unexpected logs, most likely block execution contained errors. See " + logFilename)
		//}
		//
		//outputFile.Close()
		//sourceFile.Close()
		//// remove original json log file
		//os.Remove(logFilename)
		//expectedLines := numCol*numTxPerCol + numCol + 2 // +1 for system tx, +1 for block execution log
		//assert.Equal(t, lineswritten, expectedLines)

		f(t, chain)

	}
}

func TestPrograms(t *testing.T) {
	t.Run("CustomBlock Testing w/ no mocks",
		newVMTest().run(
			func(t *testing.T, chain flow.Chain) {
			},
		),
	)
}

func generateCustomBlock(numberOfCol, numOfTxs int, accounts []flow.Address, privKeys []flow.AccountPrivateKey, chain flow.Chain) *entity.ExecutableBlock {
	collections := make([]*entity.CompleteCollection, numberOfCol)
	guarantees := make([]*flow.CollectionGuarantee, numberOfCol)
	completeCollections := make(map[flow.Identifier]*entity.CompleteCollection)

	for i := 0; i < numberOfCol; i++ {
		// func generateCustomCollection(numOfTxs int, accounts []flow.Address, privKeys []flow.AccountPrivateKey, chain flow.Chain, visitor func(body *flow.TransactionBody)) *entity.CompleteCollection {
		collection := generateCustomCollection(numOfTxs, accounts[i*numOfTxs:(i*numOfTxs)+numOfTxs], privKeys, chain, nil)
		collections[i] = collection
		guarantees[i] = collection.Guarantee
		completeCollections[collection.Guarantee.ID()] = collection
	}

	block := flow.Block{
		Header: &flow.Header{
			View: 42,
		},
		Payload: &flow.Payload{
			Guarantees: guarantees,
		},
	}

	return &entity.ExecutableBlock{
		Block:               &block,
		CompleteCollections: completeCollections,
		StartState:          unittest.StateCommitmentPointerFixture(),
	}
}

func generateCustomCollection(numOfTxs int, accounts []flow.Address, privKeys []flow.AccountPrivateKey, chain flow.Chain, visitor func(body *flow.TransactionBody)) *entity.CompleteCollection {
	transactions := make([]*flow.TransactionBody, numOfTxs)
	accontCreationScript := `
		transaction(publicKey: [UInt8]) {
			prepare(signer: AuthAccount) {
				let acct = AuthAccount(payer: signer)
				acct.addPublicKey(publicKey)
			}
		}
		`
	for i := 0; i < numOfTxs; i++ {
		// todo generate valid transactions.
		accountKey := privKeys[i].PublicKey(fvm.AccountKeyWeightThreshold)
		encAccountKey, _ := flow.EncodeRuntimeAccountPublicKey(accountKey)
		cadAccountKey := testutil.BytesToCadenceArray(encAccountKey)
		encCadAccountKey, _ := jsoncdc.Encode(cadAccountKey)

		txBody := flow.NewTransactionBody().
			SetScript([]byte(accontCreationScript)).
			AddArgument(encCadAccountKey).
			AddAuthorizer(chain.ServiceAddress()).
			SetProposalKey(chain.ServiceAddress(), 0, uint64(i)).
			SetPayer(chain.ServiceAddress())

		err := testutil.SignPayload(txBody, accounts[i], privKeys[i])
		if err != nil {
			panic(fmt.Errorf("cannot sign payload: %w", err))
		}
		err = testutil.SignEnvelope(txBody, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
		if err != nil {
			panic(fmt.Errorf("cannot sign envelope: %w", err))
		}
		if visitor != nil {
			visitor(txBody)
		}
		transactions[i] = txBody
		//fmt.Printf("tx data: %s for address: %+v\n", string(txBody.Script), txBody.Payer)

	}

	collection := flow.Collection{Transactions: transactions}

	guarantee := &flow.CollectionGuarantee{
		CollectionID: collection.ID(),
		Signature:    nil,
	}

	return &entity.CompleteCollection{
		Guarantee:    guarantee,
		Transactions: transactions,
	}
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
	pushRandAddr(size, e)
}

func (e *CustomAddressGenerator) initAtIndex(rand uint64) CustomAddressGenerator {
	e.curSeed = rand
	return *e
}

func pushRandAddr(size uint64, customAddrGen *CustomAddressGenerator) {
	if len(customAddrGen.seeds) == 0 {
		for i := 0; i < int(size); i++ {
			r := uint64(rand.Intn(maxIndex)) + 1
			custAddr := customAddrGen.initAtIndex(r)
			addr := custAddr.CurrentAddress()
			customAddrGen.addrArr[i] = addr
			//fmt.Printf("generated address: %+v\n", addr)
		}
	} else {
		for i := 0; i < int(size); i++ {
			custAddr := customAddrGen.initAtIndex(customAddrGen.seeds[i])
			addr := custAddr.CurrentAddress()
			customAddrGen.addrArr[i] = addr
			//fmt.Printf("generated address: %+v\n", addr)
		}
	}
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

func uint64ToAddress(v uint64) flow.Address {
	var b [AddressLength]byte
	binary.BigEndian.PutUint64(b[:], v)
	return flow.Address(b)
}

func uint64ToPrivKeyBytes(v uint64) []byte {
	var b [crypto.KeyGenSeedMinLenECDSAP256]byte
	binary.BigEndian.PutUint64(b[:], v)
	return b[:]
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

func convertJSONToCSV(sourceFile *os.File, outputFile *os.File) (int, error) {
	linesWritten := 0
	// struct with all possible fields for log messages
	type LogOutput struct {
		Level                string `json:"level"`
		CollectionID         string `json:"collection_id"`
		NumberOfTransactions int64  `json:"numberOfTransactions"`
		BlockID              string `json:"block_id"`
		TxID                 string `json:"tx_id"`
		TraceID              string `json:"traceID"`
		TxIndex              int64  `json:"tx_index"`
		Height               int64  `json:"height"`
		SystemChunk          bool   `json:"system_chunk"`
		ParallelExecution    bool   `json:"parallel_execution"`
		ComputationUsed      int64  `json:"computation_used"`
		TimeSpentInNS        int64  `json:"timeSpentInNS"`
		Time                 int64  `json:"time"`
		Message              string `json:"message"`
	}

	// split file into lines (each line is a JSON object)
	scanner := bufio.NewScanner(sourceFile)
	scanner.Split(bufio.ScanLines)

	// unmarshall each line into a LogOutput struct
	var logOutputs []LogOutput
	for scanner.Scan() {
		// default values of log output
		logOutput := LogOutput{
			Level:                "none",
			CollectionID:         "none",
			NumberOfTransactions: -1,
			BlockID:              "none",
			TxID:                 "none",
			TraceID:              "none",
			TxIndex:              -1,
			Height:               -1,
			SystemChunk:          false,
			ParallelExecution:    false,
			ComputationUsed:      -1,
			TimeSpentInNS:        -1,
			Time:                 -1,
			Message:              "none",
		}
		bytes := scanner.Bytes()
		err := json.Unmarshal(bytes, &logOutput)
		if err != nil {
			return linesWritten, err
		}
		logOutputs = append(logOutputs, logOutput)
	}

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// use reflection to produce csv header and get struct field values
	r := reflect.ValueOf(LogOutput{})
	typeOfTO := r.Type()
	numFields := r.NumField()
	numRecords := len(logOutputs)

	// generate the csv headers from struct fields
	var header []string
	csvData := make([][]string, numRecords)
	for col := range csvData {
		csvData[col] = make([]string, numFields)
	}

	// iterate over each field, aggregate log data per field
	for i := 0; i < numFields; i++ {
		// generate the csv header from field names as we iterate
		header = append(header, typeOfTO.Field(i).Name)
		for j := 0; j < numRecords; j++ {
			value := reflect.ValueOf(logOutputs[j]).Field(i).Interface()
			strValue := fmt.Sprintf("%v", value)
			csvData[j][i] = strValue
		}
	}
	// write header to csv file
	if err := writer.Write(header); err != nil {
		return linesWritten, err
	}
	// write all log data to csv
	for row := range csvData {
		data := csvData[row]
		// filter out all entries that do not have timing data
		if !stringSliceContainsSubstring("executing", data) {
			// write row to file
			if err := writer.Write(data); err != nil {
				return linesWritten, err
			}
			linesWritten++
		}
	}
	return linesWritten, nil
}

func stringSliceContainsSubstring(substring string, slice []string) bool {
	for _, elem := range slice {
		if strings.Contains(elem, substring) {
			return true
		}
	}
	return false
}

// GenerateAccountPrivateKeys generates a number of private keys.
func GenerateAccountPrivateKeys(numberOfPrivateKeys int, seedArr []uint64) ([]flow.AccountPrivateKey, error) {
	var privateKeys []flow.AccountPrivateKey
	for i := 0; i < numberOfPrivateKeys; i++ {
		pk, err := GenerateAccountPrivateKey(seedArr[i])
		if err != nil {
			return nil, err
		}
		privateKeys = append(privateKeys, pk)
	}

	return privateKeys, nil
}

// GenerateAccountPrivateKey generates a private key.
func GenerateAccountPrivateKey(seedInt uint64) (flow.AccountPrivateKey, error) {
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

func executeBlockAndNotVerify(t *testing.T,
	txs [][]*flow.TransactionBody,
	chain flow.Chain,
	opts []fvm.Option,
	bootstrapOpts []fvm.BootstrapProcedureOption) *execution.ComputationResult {
	rt := fvm.NewInterpreterRuntime()
	vm := fvm.NewVirtualMachine(rt)

	logger := zerolog.Nop()

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

func CreateAccountCreationTransactions(t *testing.T, chain flow.Chain, numberOfPrivateKeys int, seedArr []uint64) ([]flow.AccountPrivateKey, []flow.TransactionBody) {
	accountKeys, err := GenerateAccountPrivateKeys(numberOfPrivateKeys, seedArr)
	require.NoError(t, err)
	var txs []flow.TransactionBody

	for i := 0; i < len(accountKeys); i++ {
		keyBytes, err := flow.EncodeRuntimeAccountPublicKey(accountKeys[i].PublicKey(1000))
		require.NoError(t, err)

		// define the cadence script
		script := fmt.Sprintf(`
		transaction {
		  prepare(signer: AuthAccount) {
			let acct = AuthAccount(payer: signer)
			acct.addPublicKey("%s".decodeHex())
		  }
		}
	`, hex.EncodeToString(keyBytes))

		// create the transaction to create the account
		tx := flow.NewTransactionBody().
			SetScript([]byte(script)).
			AddAuthorizer(chain.ServiceAddress())

		txs = append(txs, *tx)
	}
	return accountKeys, txs
}

func SignTransactionAsServiceAccounts(txs []flow.TransactionBody, seqNum uint64, chain flow.Chain) {
	for i := 0; i < len(txs); i++ {
		txs[i].SetProposalKey(chain.ServiceAddress(), 0, seqNum)
		txs[i].SetPayer(chain.ServiceAddress())
		err := testutil.SignEnvelope(&txs[i], chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
		if err != nil {
			panic(fmt.Errorf("cannot sign envelope: %w", err))
		}
	}
}

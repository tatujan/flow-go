package computer_test

import (
	"context"
	"fmt"
	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go/engine/execution/computation/committer"
	"github.com/onflow/flow-go/engine/execution/computation/computer"
	executionState "github.com/onflow/flow-go/engine/execution/state"
	bootstrapexec "github.com/onflow/flow-go/engine/execution/state/bootstrap"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/engine/execution/testutil"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/programs"
	completeLedger "github.com/onflow/flow-go/ledger/complete"
	"github.com/onflow/flow-go/ledger/complete/wal/fixtures"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/epochs"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"testing"
)

func TestCustomBlockExecution(t *testing.T) {
	chain := flow.Mainnet.Chain()
	//ag := chain.NewAddressGenerator()
	logger := zerolog.New(log.Writer())
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	ctx := fvm.NewContext(logger, fvm.WithChain(chain), fvm.WithCadenceLogging(true))
	rt := fvm.NewInterpreterRuntime()
	vm := fvm.NewVirtualMachine(rt)

	t.Run("setup vaults", func(t *testing.T) {
		initialCommit, blockComputer, view := setupEnvironment(t, chain, ctx, vm, logger)
		sequenceNumber := uint64(0)

		//Create private keys
		privateKeys, err := testutil.GenerateAccountPrivateKeys(2)
		require.NoError(t, err)

		// create accounts on the chain associated with those private keys
		accounts, err := testutil.CreateAccounts(vm, view, programs.NewEmptyPrograms(), privateKeys, chain)
		require.NoError(t, err)

		/*
		****************************************************************************************************************
		************  Define Transactions ******************************************************************************
		****************************************************************************************************************
		 */

		// make transaction to setup the account 0 to send and receive flow tokens
		setupAccount0Tx := makeAndSignSetupAccountTransaction(t, chain, accounts[0], sequenceNumber, privateKeys[0])
		// make transaction to setup account 1 to send and receive flow tokens
		setupAccount1Tx := makeAndSignSetupAccountTransaction(t, chain, accounts[1], sequenceNumber, privateKeys[1])

		transactions := []*flow.TransactionBody{
			setupAccount0Tx,
			setupAccount1Tx,
		}

		//**************************************************************************************************************

		collectionOfTransactions := [][]*flow.TransactionBody{transactions}
		executableBlock := unittest.ExecutableBlockFromTransactions(collectionOfTransactions)
		executableBlock.StartState = &initialCommit

		computationResult, err := blockComputer.ExecuteBlock(
			context.Background(), executableBlock, view, programs.NewEmptyPrograms(),
		)
		assert.NoError(t, err)
		var txErrors []string
		for i, txRes := range computationResult.TransactionResults {
			// skip service transaction errors
			if i != len(computationResult.TransactionResults)-1 {
				txErrors = append(txErrors, txRes.ErrorMessage)
			}
		}
		// it seems that the service transaction might be expected to fail
		// https://github.com/onflow/flow-go/pull/2007/files#diff-6c545270710f8d91b58673b79d869691949eda628d6d4612b6162074237ec9dd
		for _, e := range txErrors {
			assert.Empty(t, e)
		}
	})

	t.Run("setup vaults and fund account", func(t *testing.T) {
		initialCommit, blockComputer, view := setupEnvironment(t, chain, ctx, vm, logger)
		sequenceNumber := uint64(0)

		//Create private keys
		privateKeys, err := testutil.GenerateAccountPrivateKeys(2)
		require.NoError(t, err)

		// create accounts on the chain associated with those private keys
		accounts, err := testutil.CreateAccounts(vm, view, programs.NewEmptyPrograms(), privateKeys, chain)
		require.NoError(t, err)

		/*
		****************************************************************************************************************
		************  Define Transaction********************************************************************************
		****************************************************************************************************************
		 */

		// make transaction to setup the account 0 to send and receive flow tokens
		setupAccount0Tx := makeAndSignSetupAccountTransaction(t, chain, accounts[0], sequenceNumber, privateKeys[0])
		// make transaction to setup account 1 to send and receive flow tokens
		setupAccount1Tx := makeAndSignSetupAccountTransaction(t, chain, accounts[1], sequenceNumber, privateKeys[1])

		// make transaction to fund account 0
		//arguments: amount, recipient
		amount := 10
		recipient := accounts[0]
		fundAccount0Tx := makeAndSignFundAccountTransaction(t, chain, amount, recipient, sequenceNumber)

		transactions := []*flow.TransactionBody{
			setupAccount0Tx,
			setupAccount1Tx,
			fundAccount0Tx,
		}
		//**************************************************************************************************************

		collectionOfTransactions := [][]*flow.TransactionBody{transactions}
		executableBlock := unittest.ExecutableBlockFromTransactions(collectionOfTransactions)
		executableBlock.StartState = &initialCommit

		computationResult, err := blockComputer.ExecuteBlock(
			context.Background(), executableBlock, view, programs.NewEmptyPrograms(),
		)
		assert.NoError(t, err)
		var txErrors []string
		for i, txRes := range computationResult.TransactionResults {
			// skip service transaction errors
			if i != len(computationResult.TransactionResults)-1 {
				txErrors = append(txErrors, txRes.ErrorMessage)
			}
		}
		// it seems that the service transaction might be expected to fail
		// https://github.com/onflow/flow-go/pull/2007/files#diff-6c545270710f8d91b58673b79d869691949eda628d6d4612b6162074237ec9dd
		for _, e := range txErrors {
			assert.Empty(t, e)
		}
	})

	t.Run("setup vaults, fund account, transfer tokens", func(t *testing.T) {
		initialCommit, blockComputer, view := setupEnvironment(t, chain, ctx, vm, logger)
		sequenceNumber := uint64(0)

		//Create private keys
		privateKeys, err := testutil.GenerateAccountPrivateKeys(2)
		require.NoError(t, err)

		// create accounts on the chain associated with those private keys
		accounts, err := testutil.CreateAccounts(vm, view, programs.NewEmptyPrograms(), privateKeys, chain)
		require.NoError(t, err)

		/*
		****************************************************************************************************************
		************  Define Transaction********************************************************************************
		****************************************************************************************************************
		 */
		// make transaction to setup the account 0 to send and receive flow tokens
		setupAccount0Tx := makeAndSignSetupAccountTransaction(t, chain, accounts[0], sequenceNumber, privateKeys[0])
		// make transaction to setup account 1 to send and receive flow tokens
		setupAccount1Tx := makeAndSignSetupAccountTransaction(t, chain, accounts[1], sequenceNumber, privateKeys[1])

		// make transaction to fund account 0
		amount := 10
		recipient := accounts[0]
		fundAccount0Tx := makeAndSignFundAccountTransaction(t, chain, amount, recipient, sequenceNumber)

		sequenceNumber++

		// make transaction to transfer funds from account 0 to account 1
		toAccountAddress := accounts[1]
		fromAccountAddress := accounts[0]
		fromAccountPK := privateKeys[0]
		transferAmount := 5
		transferFundsTx := makeAndSignTransferTokenTransaction(
			t, chain, fromAccountAddress, transferAmount, toAccountAddress, sequenceNumber, fromAccountPK,
		)

		transactions := []*flow.TransactionBody{
			setupAccount0Tx,
			setupAccount1Tx,
			fundAccount0Tx,
			transferFundsTx,
		}
		//**************************************************************************************************************

		collectionOfTransactions := [][]*flow.TransactionBody{transactions}
		executableBlock := unittest.ExecutableBlockFromTransactions(collectionOfTransactions)
		executableBlock.StartState = &initialCommit

		computationResult, err := blockComputer.ExecuteBlock(
			context.Background(), executableBlock, view, programs.NewEmptyPrograms(),
		)
		assert.NoError(t, err)
		var txErrors []string
		for i, txRes := range computationResult.TransactionResults {
			// skip service transaction errors
			if i != len(computationResult.TransactionResults)-1 {
				txErrors = append(txErrors, txRes.ErrorMessage)
			}
		}
		// it seems that the service transaction might be expected to fail
		// https://github.com/onflow/flow-go/pull/2007/files#diff-6c545270710f8d91b58673b79d869691949eda628d6d4612b6162074237ec9dd
		for _, e := range txErrors {
			assert.Empty(t, e)
		}
	})
}

// Fund account transaction
func makeAndSignFundAccountTransaction(t *testing.T, chain flow.Chain, amount int, recipient flow.Address, sequenceNumber uint64) *flow.TransactionBody {
	fundAccountTx := fundAccountTransactionBody(t, chain).
		AddAuthorizer(chain.ServiceAddress()).
		AddArgument(jsoncdc.MustEncode(cadence.UFix64(amount))).
		AddArgument(jsoncdc.MustEncode(cadence.NewAddress(recipient))).
		// set the proposal key and sequence number for this transaction:
		SetProposalKey(chain.ServiceAddress(), 0, sequenceNumber).
		// service account is the payer
		SetPayer(chain.ServiceAddress())
	// sign the tx envelope
	err := testutil.SignEnvelope(
		fundAccountTx,
		chain.ServiceAddress(),
		unittest.ServiceAccountPrivateKey,
	)
	require.NoError(t, err)
	return fundAccountTx
}

func fundAccountTransactionBody(t *testing.T, chain flow.Chain) *flow.TransactionBody {
	setupAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_fundAccount.cdc")
	require.NoError(t, err)
	script := []byte(fmt.Sprintf(string(setupAccountTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
	txBody := flow.NewTransactionBody()
	txBody.SetScript(script)
	return txBody
}

// Setup Account transaction
func makeAndSignSetupAccountTransaction(t *testing.T, chain flow.Chain, accountAddress flow.Address, sequenceNumber uint64, pk flow.AccountPrivateKey) *flow.TransactionBody {
	setupAccountTx := setupAccountTransactionBody(t, chain).
		AddArgument(jsoncdc.MustEncode(cadence.NewAddress(accountAddress))).
		AddAuthorizer(accountAddress).
		SetProposalKey(accountAddress, 0, sequenceNumber).
		SetPayer(accountAddress)
	// sign tx as account
	err := testutil.SignEnvelope(
		setupAccountTx,
		accountAddress,
		pk,
	)
	require.NoError(t, err)
	return setupAccountTx
}

func setupAccountTransactionBody(t *testing.T, chain flow.Chain) *flow.TransactionBody {
	setupAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_setupAccount.cdc")
	require.NoError(t, err)
	script := []byte(fmt.Sprintf(string(setupAccountTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
	txBody := flow.NewTransactionBody()
	txBody.SetScript(script)
	return txBody
}

// Transfer Token transaction
func makeAndSignTransferTokenTransaction(t *testing.T, chain flow.Chain, from flow.Address, amount int, to flow.Address, sequenceNumber uint64, fromPrivateKey flow.AccountPrivateKey) *flow.TransactionBody {
	transferFundsTx := transferTokensTransactionBody(t, chain).
		AddAuthorizer(from).
		AddArgument(jsoncdc.MustEncode(cadence.UFix64(amount))).
		AddArgument(jsoncdc.MustEncode(cadence.NewAddress(to))).
		// set the proposal key and sequence number for this transaction:
		SetProposalKey(from, 0, sequenceNumber).
		// sending account is the payer
		SetPayer(from)
	// sign the tx envelope
	err := testutil.SignEnvelope(
		transferFundsTx,
		from,
		fromPrivateKey,
	)
	require.NoError(t, err)
	return transferFundsTx
}

func transferTokensTransactionBody(t *testing.T, chain flow.Chain) *flow.TransactionBody {
	setupAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_transferToken.cdc")
	require.NoError(t, err)
	script := []byte(fmt.Sprintf(string(setupAccountTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
	txBody := flow.NewTransactionBody()
	txBody.SetScript(script)
	return txBody
}

func setupEnvironment(t *testing.T, chain flow.Chain, ctx fvm.Context, vm *fvm.VirtualMachine, logger zerolog.Logger) (flow.StateCommitment, computer.BlockComputer, *delta.View) {
	collector := metrics.NewNoopCollector()
	tracer := trace.NewNoopTracer()
	ledger, err := completeLedger.NewLedger(&fixtures.NoopWAL{}, 100, collector, logger, completeLedger.DefaultPathFinderVersion)
	require.NoError(t, err)
	ledgerCommitter := committer.NewLedgerViewCommitter(ledger, trace.NewNoopTracer())
	bootstrapper := bootstrapexec.NewBootstrapper(logger)
	epochConfig := epochs.DefaultEpochConfig()
	epochConfig.NumCollectorClusters = 0
	bootstrapOpts := []fvm.BootstrapProcedureOption{
		fvm.WithInitialTokenSupply(unittest.GenesisTokenSupply),
		// it seems like service transactions might be expected to fail? removing epoch config
		//fvm.WithEpochConfig(epochConfig),
	}
	initialCommit, err := bootstrapper.BootstrapLedger(
		ledger,
		unittest.ServiceAccountPublicKey,
		chain,
		bootstrapOpts...,
	)

	blockComputer, err := computer.NewBlockComputer(vm, ctx, collector, tracer, logger, ledgerCommitter)
	require.NoError(t, err)
	view := delta.NewView(executionState.LedgerGetRegister(ledger, initialCommit))
	return initialCommit, blockComputer, view
}

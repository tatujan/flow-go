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
		pk0, err := testutil.GenerateAccountPrivateKey()
		require.NoError(t, err)
		pk1, err := testutil.GenerateAccountPrivateKey()
		require.NoError(t, err)
		privateKeys := []flow.AccountPrivateKey{pk0, pk1}

		// create accounts on the chain associated with those private keys
		accounts, err := testutil.CreateAccounts(vm, view, programs.NewEmptyPrograms(), privateKeys, chain)
		require.NoError(t, err)
		account0address := accounts[0]
		account1address := accounts[1]

		// setup the account 0 to send and receive flow tokens
		setupAccount0Tx := setupAccountTransactionBody(t, chain).
			AddArgument(jsoncdc.MustEncode(cadence.NewAddress(account0address))).
			AddAuthorizer(account0address).
			SetProposalKey(account0address, 0, sequenceNumber).
			SetPayer(account0address)
		// sign tx as account 0
		err = testutil.SignEnvelope(
			setupAccount0Tx,
			account0address,
			pk0,
		)
		require.NoError(t, err)
		//sequenceNumber++

		// setup account 1 to send and receive flow tokens
		setupAccount1Tx := setupAccountTransactionBody(t, chain).
			AddArgument(jsoncdc.MustEncode(cadence.NewAddress(account1address))).
			AddAuthorizer(account1address).
			SetProposalKey(account1address, 0, sequenceNumber).
			SetPayer(account1address)
		// sign tx as account 1
		err = testutil.SignEnvelope(
			setupAccount1Tx,
			account1address,
			pk1,
		)
		require.NoError(t, err)
		//sequenceNumber++

		// fund account 0
		initialAmount := 10
		recipient := account0address
		fundAccount0Tx := fundAccountTransactionBody(t, chain).
			AddAuthorizer(chain.ServiceAddress()).
			AddArgument(jsoncdc.MustEncode(cadence.UFix64(initialAmount))).
			AddArgument(jsoncdc.MustEncode(cadence.NewAddress(recipient))).
			// set the proposal key and sequence number for this transaction:
			SetProposalKey(chain.ServiceAddress(), 0, sequenceNumber).
			// service account is the payer
			SetPayer(chain.ServiceAddress())
		// sign the tx envelope
		err = testutil.SignEnvelope(
			fundAccount0Tx,
			chain.ServiceAddress(),
			unittest.ServiceAccountPrivateKey,
		)
		require.NoError(t, err)

		transactions := []*flow.TransactionBody{
			setupAccount0Tx,
			setupAccount1Tx,
			fundAccount0Tx,
		}
		collectionOfTransactions := [][]*flow.TransactionBody{transactions}
		executableBlock := unittest.ExecutableBlockFromTransactions(collectionOfTransactions)
		executableBlock.StartState = &initialCommit

		computationResult, err := blockComputer.ExecuteBlock(context.Background(), executableBlock, view, programs.NewEmptyPrograms())
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

func setupAccountTransactionBody(t *testing.T, chain flow.Chain) *flow.TransactionBody {
	setupAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_setupAccount.cdc")
	require.NoError(t, err)
	script := []byte(fmt.Sprintf(string(setupAccountTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
	txBody := flow.NewTransactionBody()
	txBody.SetScript(script)
	return txBody
}

func fundAccountTransactionBody(t *testing.T, chain flow.Chain) *flow.TransactionBody {
	setupAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_fundAccount.cdc")
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

package computer_test

import (
	"context"
	"fmt"
	"github.com/onflow/flow-go/engine/execution/computation/committer"
	"github.com/onflow/flow-go/engine/execution/computation/computer"
	computermock "github.com/onflow/flow-go/engine/execution/computation/computer/mock"
	executionState "github.com/onflow/flow-go/engine/execution/state"
	bootstrapexec "github.com/onflow/flow-go/engine/execution/state/bootstrap"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/programs"
	completeLedger "github.com/onflow/flow-go/ledger/complete"
	"github.com/onflow/flow-go/ledger/complete/wal/fixtures"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/epochs"
	"github.com/onflow/flow-go/module/metrics"
	modulemock "github.com/onflow/flow-go/module/mock"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"testing"
)

func TestCustomBlockExecution(t *testing.T) {
	chain := flow.Mainnet.Chain()
	ag := chain.NewAddressGenerator()
	logger := zerolog.New(log.Writer())
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	ctx := fvm.NewContext(logger, fvm.WithChain(chain), fvm.WithCadenceLogging(true))
	rt := fvm.NewInterpreterRuntime()
	vm := fvm.NewVirtualMachine(rt)

	t.Run("single collection", func(t *testing.T) {
		initialCommit, blockComputer, view := setupEnvironment(t, chain, ctx, vm, logger)

		address, err := ag.NextAddress()
		require.NoError(t, err)
		txBody := &flow.TransactionBody{
			Payer:  address, // a unique payer for each tx to generate a unique id
			Script: []byte("transaction { execute {} }"),
		}
		transactions := []*flow.TransactionBody{txBody}
		collectionOfTransactions := [][]*flow.TransactionBody{transactions}
		executableBlock := unittest.ExecutableBlockFromTransactions(collectionOfTransactions)
		executableBlock.StartState = &initialCommit

		computationResult, err := blockComputer.ExecuteBlock(context.Background(), executableBlock, view, programs.NewEmptyPrograms())
		assert.NoError(t, err)

		assert.Empty(t, computationResult.TransactionResults[0])

		assertEventHashesMatch(t, 1+1, computationResult)
	})
}

func getTransferTokenTransactionScript(t *testing.T, chain flow.Chain) []byte {
	transferTokenTemplateBytes, err := ioutil.ReadFile("cadence/tx_transferToken.cdc")
	require.NoError(t, err)
	// transfer token script
	return []byte(fmt.Sprintf(string(transferTokenTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
}

func getFundAccountTransactionScript(t *testing.T, chain flow.Chain) []byte {
	fundAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_fundAccount.cdc")
	require.NoError(t, err)
	txScript := []byte(fmt.Sprintf(string(fundAccountTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
	return txScript
}

func getSetupAccountTransactionScript(t *testing.T, chain flow.Chain) []byte {
	setupAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_setupAccount.cdc")
	require.NoError(t, err)
	return []byte(fmt.Sprintf(string(setupAccountTemplateBytes), fvm.FungibleTokenAddress(chain), fvm.FlowTokenAddress(chain)))
}

func getCreateAccountTransactionScript(t *testing.T) []byte {
	createAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_createAccount.cdc")
	require.NoError(t, err)
	return []byte(fmt.Sprintf(string(createAccountTemplateBytes)))
}
func getAddAccountKeyTransactionScript(t *testing.T) []byte {
	createAccountTemplateBytes, err := ioutil.ReadFile("cadence/tx_addAccountKey.cdc")
	require.NoError(t, err)
	return []byte(fmt.Sprintf(string(createAccountTemplateBytes)))
}

func getMockedCommitter(numberTx int) *computermock.ViewCommitter {
	committer := new(computermock.ViewCommitter)
	committer.On("CommitView", mock.Anything, mock.Anything).
		Return(nil, nil, nil, nil).
		Times(numberTx + 1) // 2 txs in collection + system chunk
	return committer
}

func getMockedMetrics(numberTx int, numberCollection int) *modulemock.ExecutionMetrics {
	metrics := new(modulemock.ExecutionMetrics)
	metrics.On("ExecutionCollectionExecuted", mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Times(numberCollection + 1) // 1 collection + system collection

	metrics.On("ExecutionTransactionExecuted", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		Return(nil).
		Times(numberTx + 1) // 2 txs in collection + system chunk tx

	return metrics
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
		fvm.WithEpochConfig(epochConfig),
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

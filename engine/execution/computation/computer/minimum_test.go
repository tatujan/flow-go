package computer_test

import (
	"context"
	"fmt"
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
	logger := zerolog.New(log.Writer())
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	ctx := fvm.NewContext(logger, fvm.WithChain(chain), fvm.WithCadenceLogging(true))
	rt := fvm.NewInterpreterRuntime()
	vm := fvm.NewVirtualMachine(rt)

	t.Run("setup vaults", func(t *testing.T) {
		initialCommit, blockComputer, view := setupEnvironment(t, chain, ctx, vm, logger)
		sequenceNumber := uint64(0)
		createAccount0Tx := flow.NewTransactionBody().
			// use basic account creation script
			SetScript(getSetupAccountTransactionScript(t, chain)).
			AddAuthorizer(chain.ServiceAddress()).
			SetProposalKey(chain.ServiceAddress(), 0, sequenceNumber).
			SetPayer(chain.ServiceAddress())
		// sign tx as service account
		err := testutil.SignEnvelope(
			createAccount0Tx,
			chain.ServiceAddress(),
			unittest.ServiceAccountPrivateKey,
		)
		require.NoError(t, err)
		sequenceNumber++

		transactions := []*flow.TransactionBody{createAccount0Tx}
		collectionOfTransactions := [][]*flow.TransactionBody{transactions}
		executableBlock := unittest.ExecutableBlockFromTransactions(collectionOfTransactions)
		executableBlock.StartState = &initialCommit

		computationResult, err := blockComputer.ExecuteBlock(context.Background(), executableBlock, view, programs.NewEmptyPrograms())
		assert.NoError(t, err)
		var txErrors []string
		for _, txRes := range computationResult.TransactionResults {
			txErrors = append(txErrors, txRes.ErrorMessage)
		}
		for _, e := range txErrors {
			assert.Empty(t, e)
		}
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

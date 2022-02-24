package computer_test

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/onflow/flow-go/engine/execution/computation/committer"
	"github.com/onflow/flow-go/engine/execution/computation/computer"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	led "github.com/onflow/flow-go/ledger"
	ledgermock "github.com/onflow/flow-go/ledger/mock"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module/mempool/entity"
	modulemock "github.com/onflow/flow-go/module/mock"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	f func(t *testing.T, vm *fvm.VirtualMachine, chain flow.Chain, ctx fvm.Context, view state.View),
) func(t *testing.T) {
	return func(t *testing.T) {
		numTxPerCol := 3
		numCol := 3
		seeds := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		//rag := &RandomAddressGenerator{}

		// todo: DONE! 1. Initialize Computer.VirtualMachine
		//vm := new(computermock.VirtualMachine)
		runtime := fvm.NewInterpreterRuntime()
		chain := flow.Testnet.Chain()
		vm := fvm.NewVirtualMachine(runtime)
		baseOpts := []fvm.Option{
			fvm.WithChain(chain),
		}
		opts := append(baseOpts, vmt.contextOptions...)
		execCtx := fvm.NewContext(zerolog.Nop(), opts...)

		// todo: [WIP] 2. Init computer.ViewCommitter
		// code snippet to unmock ledger
		//wal := &fixtures.NoopWAL{}
		//ledger, err := complete.NewLedger(wal, 100, &metrics.NoopCollector{}, zerolog.Logger{}, complete.DefaultPathFinderVersion)
		//curS := ledger.InitialState()
		//q := utils.QueryFixture()
		//q.SetState(curS)
		//require.NoError(t, err)
		//retProof, err := ledger.Prove(q)
		//require.NoError(t, err)
		//trieProof, err := encoding.DecodeTrieBatchProof(retProof)
		//assert.True(t, proof.VerifyTrieBatchProof(trieProof, curS))

		// unmocking the ledger requires to provide proofs of state commitment.
		// However dummy transactions w/o signatures do not yield valid proofs.
		ledger := new(ledgermock.Ledger)
		var expectedStateCommitment led.State
		copy(expectedStateCommitment[:], []byte{1, 2, 3})
		ledger.On("Set", mock.Anything).
			Return(expectedStateCommitment, nil, nil).
			Times(numTxPerCol*numCol + 1)
		expectedProof := led.Proof([]byte{2, 3, 4})
		ledger.On("Prove", mock.Anything).
			Return(expectedProof, nil).
			Times(numTxPerCol*numCol + 1)

		blockcommitter := committer.NewLedgerViewCommitter(ledger, trace.NewNoopTracer())

		// todo: 3. Initialize module.ExectionMetrics
		metrics := new(modulemock.ExecutionMetrics)
		// ExecutionColllectionExecuted takes time.Duration, compUsed(uint64), txCounts, colCounts
		metrics.On("ExecutionCollectionExecuted", mock.Anything, mock.Anything, mock.Anything).
			Return(nil).
			Times(numCol) // num tx + system chunk

		//dur, compUsed, eventCounts, failed
		metrics.On("ExecutionTransactionExecuted", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil).
			Times(numCol*numTxPerCol + 1)

		// todo: 4. Init New Computer from vm, execCtx, metric, trace, logger, comitter
		exe, err := computer.NewBlockComputer(vm, execCtx, metrics, trace.NewNoopTracer(), zerolog.Nop(), blockcommitter)
		require.NoError(t, err)

		// Generate your own block with 2 collections and 4 txs in total
		//block := generateBlock(numCol, numTxPerCol, rag)
		customAddr := new(CustomAddressGenerator)
		customAddr.init(uint64(numCol*numTxPerCol), seeds)
		block := generateCustomBlock(numCol, numTxPerCol, customAddr)

		// returns nill register value
		view := delta.NewView(func(owner, controller, key string) (flow.RegisterValue, error) {
			return nil, nil
		})

		result, err := exe.ExecuteBlock(context.Background(), block, view, programs.NewEmptyPrograms())
		assert.NoError(t, err)
		assert.Len(t, result.StateSnapshots, numCol+1)
		assert.Len(t, result.TrieUpdates, numCol+1)

		assertEventHashesMatch(t, numCol+1, result)
		//vm.AssertExpectations(t)
		f(t, vm, chain, execCtx, view)

	}
}

func TestPrograms(t *testing.T) {
	t.Run("CustomBlock Testing w/ no mocks",
		newVMTest().run(
			func(t *testing.T, vm *fvm.VirtualMachine, chain flow.Chain, ctx fvm.Context, view state.View) {
			},
		),
	)
}

func generateCustomBlock(numberOfCol, numOfTxs int, customAddr *CustomAddressGenerator) *entity.ExecutableBlock {
	collections := make([]*entity.CompleteCollection, numberOfCol)
	guarantees := make([]*flow.CollectionGuarantee, numberOfCol)
	completeCollections := make(map[flow.Identifier]*entity.CompleteCollection)

	for i := 0; i < numberOfCol; i++ {
		collection := generateCustomCollection(numOfTxs, customAddr, nil)
		collections[i] = collection
		guarantees[i] = collection.Guarantee
		completeCollections[collection.Guarantee.ID()] = collection
	}

	block := flow.Block{
		Header: &flow.Header{
			Timestamp: flow.GenesisTime,
			Height:    42,
			View:      42,
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

func generateCustomCollection(numOfTxs int, customAddr *CustomAddressGenerator, visitor func(body *flow.TransactionBody)) *entity.CompleteCollection {
	transactions := make([]*flow.TransactionBody, numOfTxs)

	for i := 0; i < numOfTxs; i++ {
		nextAddress, err := customAddr.NextAddress()
		if err != nil {
			panic(fmt.Errorf("cannot generate next address in test: %w", err))
		}
		txBody := &flow.TransactionBody{
			Payer:  nextAddress, // a unique payer for each tx to generate a unique id
			Script: []byte("transaction: { execute {updateVar:{'x': 1, 'y':2}} }"),
		}
		if visitor != nil {
			visitor(txBody)
		}
		transactions[i] = txBody
		fmt.Printf("tx data: %s for address: %+v\n", string(txBody.Script), txBody.Payer)

	}

	collection := flow.Collection{Transactions: transactions}

	guarantee := &flow.CollectionGuarantee{CollectionID: collection.ID()}

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

func (e *CustomAddressGenerator) CurrentAddress() flow.Address {
	return uint64ToAddress(e.curSeed)
}

func (e *CustomAddressGenerator) Bytes() []byte {
	panic("not implemented")
}

func (e *CustomAddressGenerator) AddressCount() uint64 {
	panic("not implemented")
}

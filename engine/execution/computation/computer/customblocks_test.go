package computer_test

import (
	"bufio"
	"context"
	"encoding/binary"
	"encoding/csv"
	"encoding/json"
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
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/module/trace"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	f func(t *testing.T, vm *fvm.VirtualMachine, chain flow.Chain, ctx fvm.Context, view state.View),
) func(t *testing.T) {
	return func(t *testing.T) {
		numTxPerCol := 3
		numCol := 3
		seeds := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		//rag := &RandomAddressGenerator{}

		// todo: 1. Initialize Computer.VirtualMachine
		//vm := new(computermock.VirtualMachine)
		runtime := fvm.NewInterpreterRuntime()
		chain := flow.Testnet.Chain()
		vm := fvm.NewVirtualMachine(runtime)
		baseOpts := []fvm.Option{
			fvm.WithChain(chain),
		}

		opts := append(baseOpts, vmt.contextOptions...)
		execCtx := fvm.NewContext(zerolog.Nop(), opts...)
		//view := utils.NewSimpleView()
		//// VM Run takes, (fvm.Context, fvm.Procedure, state.View, programs.Programs)
		//vm.On("Run", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
		//	Return(nil).
		//	Run(func(args mock.Arguments) {
		//		tx := args[1].(*fvm.TransactionProcedure) // get transactions Procedures?
		//		tx.Events = generateEvents(1, tx.TxIndex) // generate events from txIndex
		//	}).Times(numTxPerCol*numCol + 1) // set how many times this mock is going to return ( numTX + systemchunk)

		// todo: 2. Init computer.ViewCommitter
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
		committer := committer.NewLedgerViewCommitter(ledger, trace.NewNoopTracer())

		//committer := new(computermock.ViewCommitter)
		//// View committer takes (state.View, flow.StateCmmitment)
		//// Returns flow.StateCommitment, ByteArray, *ledger.TrieUpdate, error
		//committer.On("CommitView", mock.Anything, mock.Anything).
		//	Return(nil, nil, nil, nil).
		//	Times(numTxPerCol*numCol + 1)

		// todo: 3. Initialize module.ExectionMetrics
		logFilename := "logFile"
		csvFilename := "customBlockLogOutput.csv"
		file, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		logger := zerolog.New(file)
		tracer, err := trace.NewTracer(logger, "CustomBlockTrace", "test", trace.SensitivityCaptureAll)
		metrics := metrics.NewExecutionCollector(tracer, prometheus.DefaultRegisterer)

		// todo: 4. Init New Computer from vm, execCtx, metric, trace, logger, comitter
		exe, err := computer.NewBlockComputer(vm, execCtx, metrics, tracer, logger, committer)
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

		// open file containing logs in JSON format
		sourceFile, err := os.Open(logFilename)
		if err != nil {
			// do nothing?
		}
		// create csv file
		outputFile, err := os.Create(csvFilename)
		if err != nil {
			// do nothing?
		}
		// convert the JSON logs to CSV file
		lineswritten, err := convertJSONToCSV(sourceFile, outputFile)
		if err != nil {
			// do nothing
		}
		if lineswritten == 0 {
			panic("Unexpected logs, most likely block execution contained errors. See " + logFilename)
		}

		outputFile.Close()
		sourceFile.Close()
		// remove original json log file
		os.Remove(logFilename)
		expectedLines := numCol*numTxPerCol + numCol + 2 // +1 for system tx, +1 for block execution log
		assert.Equal(t, lineswritten, expectedLines)

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

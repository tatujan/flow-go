package conflicts_test

import (
	"github.com/onflow/flow-go/engine/execution/computation/conflicts"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestTransactionConflictGraph(t *testing.T) {

	numTx := 5
	numRegisters := numTx * 2
	registers := createRegisters(numRegisters)
	nonConflictingTransactions := createNonConflictingTransactions(numTx, registers)
	seperateConflictingTransactions := createSeperateConflictingTransactions(numTx, registers)
	cascadedCOnflictingTransactions := createCascadedConflictingTransactions(numTx, registers)
	seperatedConflictingTransactions := createSeperatedAndCascadedConflictingTransactions(numTx, registers)

	t.Run("Transactions add successfully", func(t *testing.T) {
		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(nonConflictingTransactions[n])
		}
		c.Close()
		<-cDone
		require.Equal(t, numTx, c.TransactionCount())
	})

	t.Run("No conflicts for non-conflicting transactions", func(t *testing.T) {
		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(nonConflictingTransactions[n])
		}
		c.Close()
		<-cDone
		require.Equal(t, 0, c.ConflictCount())

	})

	t.Run("Transactions touching same register in different collections are conflicting", func(t *testing.T) {
		conflictingTransactions := nonConflictingTransactions
		// set first tx touchset to be the same as the second
		conflictingTransactions[1].RegisterTouchSet = conflictingTransactions[0].RegisterTouchSet
		// set the registers touched by the lst tx to be the same registers touched by the second to last
		conflictingTransactions[numTx-1].RegisterTouchSet = conflictingTransactions[numTx-2].RegisterTouchSet

		conflictingTx1 := conflictingTransactions[1]
		conflictingTx2 := conflictingTransactions[numTx-1]

		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(nonConflictingTransactions[n])
		}
		c.Close()
		<-cDone
		for _, printTx := range conflictingTransactions {
			idStr := printTx.TransactionID.String()
			print("Tx index: ", printTx.TxIndex, " Tx ID: ", idStr[len(idStr)-8:], "\n")
		}
		print("\n" + c.String() + "\n")
		require.True(t, c.ContainsNode(conflictingTx1.TransactionID))
		require.True(t, c.ContainsNode(conflictingTx2.TransactionID))
		require.Equal(t, 2, c.ConflictCount())

	})

	t.Run("Transactions touching same register in same collections are not conflicting", func(t *testing.T) {
		conflictingTransactions := nonConflictingTransactions
		// set first tx touchset and collectionID to be the same as the second
		conflictingTransactions[1].RegisterTouchSet = conflictingTransactions[0].RegisterTouchSet
		conflictingTransactions[1].CollectionID = conflictingTransactions[0].CollectionID

		// set last tx touchset and collectionID to be the same as the second-to-last
		conflictingTransactions[numTx-1].RegisterTouchSet = conflictingTransactions[numTx-2].RegisterTouchSet
		conflictingTransactions[numTx-1].CollectionID = conflictingTransactions[numTx-2].CollectionID

		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(nonConflictingTransactions[n])
		}
		c.Close()
		<-cDone
		require.Equal(t, 0, c.ConflictCount())

	})

	t.Run("neighboring Transactions touching same register in different collections are conflicting", func(t *testing.T) {

		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(seperateConflictingTransactions[n])
		}
		c.Close()
		<-cDone
		for _, printTx := range seperateConflictingTransactions {
			idStr := printTx.TransactionID.String()
			print("Tx index: ", printTx.TxIndex, " Tx ID: ", idStr[len(idStr)-8:], "\n")
		}
		print("\n" + c.String() + "\n")
		for i, _ := range seperateConflictingTransactions {
			if i%2 == 1 {
				require.True(t, c.ContainsNode(seperateConflictingTransactions[i].TransactionID))
			}
		}
		require.Equal(t, 2, c.ConflictCount())

	})
	t.Run("All neigborhing Transactions touching same register in different collections are conflicting", func(t *testing.T) {

		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(cascadedCOnflictingTransactions[n])
		}
		c.Close()
		<-cDone
		for _, printTx := range cascadedCOnflictingTransactions {
			idStr := printTx.TransactionID.String()
			print("Tx index: ", printTx.TxIndex, " Tx ID: ", idStr[len(idStr)-8:], "\n")
		}
		print("\n" + c.String() + "\n")
		for i, _ := range cascadedCOnflictingTransactions {
			require.True(t, c.ContainsNode(cascadedCOnflictingTransactions[i].TransactionID))
		}
		require.Equal(t, numTx-1, c.ConflictCount())
	})

	t.Run("Conflicts to create a two branched Conflict tree", func(t *testing.T) {
		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for n := 0; n < numTx; n++ {
			c.StoreTransaction(seperatedConflictingTransactions[n])
		}
		c.Close()
		<-cDone
		for _, printTx := range seperatedConflictingTransactions {
			idStr := printTx.TransactionID.String()
			print("Tx index: ", printTx.TxIndex, " Tx ID: ", idStr[len(idStr)-8:], "\n")
		}
		print("\n" + c.String() + "\n")
		for i, _ := range []int{0: 4} {
			require.True(t, c.ContainsNode(seperatedConflictingTransactions[i].TransactionID))
		}
		//require.Equal(t, numTx-1, c.ConflictCount())
	})

}

func createSeperatedAndCascadedConflictingTransactions(numNonConflictingTx int, registers []flow.RegisterID) []conflicts.Transaction {
	var seperatedAndCascadedConflictingTransactions []conflicts.Transaction
	var touchset []flow.RegisterID
	for i := 0; i < numNonConflictingTx; i++ {
		txID := unittest.IdentifierFixture()
		collID := unittest.IdentifierFixture()
		switch i {
		case 1:
			{
				touchset = []flow.RegisterID{registers[0]}
				touchset = append(touchset, registers[2])
			}
		case 2:
			{
				touchset = []flow.RegisterID{registers[1]}
				touchset = append(touchset, registers[3])
			}
		case 3:
			{
				touchset = []flow.RegisterID{registers[3]}
				touchset = append(touchset, registers[4])
			}
		default:
			touchset = registers[i*2 : i*2+2]
		}

		tx := conflicts.Transaction{
			TransactionID:    txID,
			CollectionID:     collID,
			RegisterTouchSet: touchset,
			TxIndex:          uint32(i),
		}

		seperatedAndCascadedConflictingTransactions = append(seperatedAndCascadedConflictingTransactions, tx)
	}
	return seperatedAndCascadedConflictingTransactions
}

func createCascadedConflictingTransactions(numConflictingTx int, registers []flow.RegisterID) []conflicts.Transaction {
	var conflictingTxs []conflicts.Transaction
	for i := 0; i < numConflictingTx; i++ {
		txID := unittest.IdentifierFixture()
		collID := unittest.IdentifierFixture()
		touchset := registers[i*2 : i*2+2]
		if i != 0 {
			touchset = registers[i : i+2]
		}

		tx := conflicts.Transaction{
			TransactionID:    txID,
			CollectionID:     collID,
			RegisterTouchSet: touchset,
			TxIndex:          uint32(i),
		}

		conflictingTxs = append(conflictingTxs, tx)
	}
	return conflictingTxs
}

func createSeperateConflictingTransactions(numConflictingTx int, registers []flow.RegisterID) []conflicts.Transaction {
	var conflictingTxs []conflicts.Transaction
	for i := 0; i < numConflictingTx; i++ {
		txID := unittest.IdentifierFixture()
		collID := unittest.IdentifierFixture()
		touchset := registers[i*2 : i*2+2]
		if i%2 == 1 {
			touchset = registers[i*2-1 : i*2-1+2]
		}

		tx := conflicts.Transaction{
			TransactionID:    txID,
			CollectionID:     collID,
			RegisterTouchSet: touchset,
			TxIndex:          uint32(i),
		}

		conflictingTxs = append(conflictingTxs, tx)
	}
	return conflictingTxs
}

func createNonConflictingTransactions(numNonConflictingTx int, registers []flow.RegisterID) []conflicts.Transaction {
	// set up a set of non conflicting transactions (none share registers)
	var nonConflictingTransactions []conflicts.Transaction
	for i := 0; i < numNonConflictingTx; i++ {
		txID := unittest.IdentifierFixture()
		collID := unittest.IdentifierFixture()
		touchset := registers[i*2 : i*2+2]

		tx := conflicts.Transaction{
			TransactionID:    txID,
			CollectionID:     collID,
			RegisterTouchSet: touchset,
			TxIndex:          uint32(i),
		}

		nonConflictingTransactions = append(nonConflictingTransactions, tx)
	}
	return nonConflictingTransactions
}

func createRegisters(numRegisters int) []flow.RegisterID {
	// set up a set of registers
	var registers []flow.RegisterID
	for j := 0; j < numRegisters; j++ {
		owner := "owner" + strconv.Itoa(j)
		controller := "controller" + strconv.Itoa(j)
		key := strconv.Itoa(j)
		r := flow.NewRegisterID(owner, controller, key)
		registers = append(registers, r)
	}
	return registers
}

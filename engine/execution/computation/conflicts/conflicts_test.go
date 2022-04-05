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
		require.True(t, c.containsNode(conflictingTx1.TransactionID))
		require.True(t, c.containsNode(conflictingTx2.TransactionID))
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

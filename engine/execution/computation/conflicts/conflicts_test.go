package conflicts_test

import (
	"github.com/onflow/flow-go/engine/execution/computation/conflicts"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransactionConflictGraph(t *testing.T) {

	t.Run("Transactions added successfully", func(t *testing.T) {
		numTx := 5
		c := conflicts.NewConflicts(numTx)
		cDone := make(chan bool, 1)
		go func() {
			c.Run()
			cDone <- true
		}()

		for i := 0; i < numTx; i++ {
			id := unittest.IdentifierFixture()
			v := delta.NewView(func(owner, controller, key string) (flow.RegisterValue, error) {
				return nil, nil
			}).NewChild()
			c.AddTransaction(conflicts.Message{TransactionID: id, TransactionView: &v})
		}
		c.Close()
		<-cDone
		require.Equal(t, numTx, c.TransactionCount())
	})

}

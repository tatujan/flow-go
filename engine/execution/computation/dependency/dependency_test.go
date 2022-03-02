package dependency_test

import (
	"github.com/onflow/flow-go/engine/execution/computation/dependency"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/utils/unittest"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTransactionDependencyGraph(t *testing.T) {

	t.Run("Transactions added successfully", func(t *testing.T) {

		d := dependency.NewDependency()
		numTx := 5
		for i := 0; i < numTx; i++ {
			id := unittest.IdentifierFixture()
			v := delta.NewView(func(owner, controller, key string) (flow.RegisterValue, error) {
				return nil, nil
			})
			d.AddTransaction(id, v)
		}

		require.Equal(t, numTx, d.TransactionCount())
	})

}

package reporters_test

import (
	"math"
	"testing"

	"github.com/onflow/atree"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/flow-go/cmd/util/ledger/migrations"
	"github.com/onflow/flow-go/cmd/util/ledger/reporters"
	"github.com/onflow/flow-go/engine/execution/state"
	state2 "github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/utils"
	"github.com/onflow/flow-go/model/flow"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func createAccountPayloadKey(a flow.Address, key string) ledger.Key {
	return ledger.Key{
		KeyParts: []ledger.KeyPart{
			ledger.NewKeyPart(state.KeyPartOwner, a.Bytes()),
			ledger.NewKeyPart(state.KeyPartController, []byte("")),
			ledger.NewKeyPart(state.KeyPartKey, []byte(key)),
		},
	}
}

func TestLookupValues(t *testing.T) {
	dir := t.TempDir()
	mig := migrations.OrderedMapMigration{
		Log:       zerolog.Logger{},
		OutputDir: dir,
	}

	address1 := flow.HexToAddress("0x1")

	encodeValue := func(v interpreter.Value) ledger.Value {
		storable, _ := v.Storable(mig.NewStorage, atree.Address(address1), math.MaxUint64)
		encodedInt, _ := atree.Encode(storable, interpreter.CBOREncMode)
		return encodedInt
	}

	one := interpreter.NewIntValueFromInt64(1)

	payload := []ledger.Payload{
		{Key: createAccountPayloadKey(address1, state2.KeyExists), Value: []byte{1}},
		{Key: createAccountPayloadKey(address1, state2.KeyStorageUsed), Value: utils.Uint64ToBinary(1)},
		{Key: createAccountPayloadKey(address1, "storage\x1fFoo"), Value: encodeValue(one)},
	}
	migratedPayload, _ := mig.Migrate(payload)

	l := migrations.NewView(migratedPayload)
	bp := reporters.NewBalanceReporter(flow.Testnet.Chain(), l)

	stored, err := bp.ReadStored(address1, common.PathDomainStorage, "Foo")
	require.NoError(t, err)
	require.Equal(t, stored, one)
}

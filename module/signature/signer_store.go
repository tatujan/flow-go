package signature

import (
	"errors"
	"fmt"

	"github.com/onflow/flow-go/model/encoding"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/storage"
)

// EpochAwareSignerStore implements the SignerStore interface. It is epoch
// aware, and provides the appropriate threshold signers on a per-view basis,
// using the database to retrieve the relevant DKG keys.
type EpochAwareSignerStore struct {
	epochLookup module.EpochLookup                // used to fetch epoch counter by view
	keys        storage.DKGKeys                   // used to fetch DKG private key by epoch
	signers     map[uint64]module.ThresholdSigner // cache of signers by epoch
}

// NewEpochAwareSignerStore instantiates a new EpochAwareSignerStore
func NewEpochAwareSignerStore(epochLookup module.EpochLookup, keys storage.DKGKeys) *EpochAwareSignerStore {
	return &EpochAwareSignerStore{
		epochLookup: epochLookup,
		keys:        keys,
		signers:     make(map[uint64]module.ThresholdSigner),
	}
}

// GetThresholdSigner returns the threshold-signer for signing objects at a
// given view. The view determines the epoch, which determines the DKG private
// key underlying the signer.
// It returns:
//  - (signer, true, nil) if DKG was completed in the epoch of the view
//  - (nil, false, nil) if DKG was not completed in the epoch of the view
//  - (nil, false, err) if there is any exception
func (s *EpochAwareSignerStore) GetThresholdSigner(view uint64) (module.ThresholdSigner, bool, error) {
	epoch, err := s.epochLookup.EpochForViewWithFallback(view)
	if err != nil {
		return nil, false, fmt.Errorf("could not get epoch by view %v: %w", view, err)
	}
	signer, ok := s.signers[epoch]
	if ok {
		// A nil signer means that we don't have a Random Beacon key for this epoch.
		return signer, signer != nil, nil
	}

	privDKGData, err := s.keys.RetrieveMyDKGPrivateInfo(epoch)
	// DKG was not completed, there is no DKG private key
	if errors.Is(err, storage.ErrNotFound) {
		s.signers[epoch] = nil
		return nil, false, nil
	}

	if err != nil {
		return nil, false, fmt.Errorf("could not retrieve DKG private key for epoch counter %v, at view %v, err: %w", epoch, view, err)
	}

	// DKG was completed and a random beacon key is available
	signer = NewThresholdProvider(encoding.RandomBeaconTag, privDKGData.RandomBeaconPrivKey)
	s.signers[epoch] = signer

	return signer, true, nil
}

/*~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~*/

// SingleSignerStore implements the ThresholdSignerStore interface. It only
// keeps one signer and is not epoch-aware. It is used only for the
// bootstrapping process.
type SingleSignerStore struct {
	signer module.ThresholdSigner
}

// NewSingleSignerStore instantiates a new SingleSignerStore.
func NewSingleSignerStore(signer module.ThresholdSigner) *SingleSignerStore {
	return &SingleSignerStore{
		signer: signer,
	}
}

// GetThresholdSigner returns the signer.
func (s *SingleSignerStore) GetThresholdSigner(view uint64) (module.ThresholdSigner, error) {
	return s.signer, nil
}

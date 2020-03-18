// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package procedure

import (
	"fmt"
	"math"

	"github.com/dgraph-io/badger/v2"

	"github.com/dapperlabs/flow-go/model/flow"
	"github.com/dapperlabs/flow-go/storage/badger/operation"
)

func InsertBlock(block *flow.Block) func(*badger.Txn) error {
	return func(tx *badger.Txn) error {

		// store the block header
		err := operation.InsertHeader(&block.Header)(tx)
		if err != nil {
			return fmt.Errorf("could not insert block header: %w", err)
		}

		// insert the block payload
		err = InsertPayload(&block.Payload)(tx)
		if err != nil {
			return fmt.Errorf("could not insert block payload: %w", err)
		}

		// index the block payload
		err = IndexPayload(&block.Header, &block.Payload)(tx)
		if err != nil {
			return fmt.Errorf("could not index block payload: %w", err)
		}

		return nil

	}
}

func RetrieveBlock(blockID flow.Identifier, block *flow.Block) func(*badger.Txn) error {
	return func(tx *badger.Txn) error {

		// get the block header
		err := operation.RetrieveHeader(blockID, &block.Header)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve header: %w", err)
		}

		// get the block payload
		err = RetrievePayload(block.Header.ID(), &block.Payload)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve payload: %w", err)
		}

		return nil
	}
}

func FinalizeBlock(blockID flow.Identifier) func(*badger.Txn) error {
	return func(tx *badger.Txn) error {

		// retrieve the header to check the parent
		var header flow.Header
		err := operation.RetrieveHeader(blockID, &header)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve block: %w", err)
		}

		// retrieve the current finalized state boundary
		var boundary uint64
		err = operation.RetrieveBoundary(&boundary)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve boundary: %w", err)
		}

		// retrieve the ID of the boundary head
		var headID flow.Identifier
		err = operation.RetrieveNumber(boundary, &headID)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve head: %w", err)
		}

		// check that the head ID is the parent of the block we finalize
		if header.ParentID != headID {
			return fmt.Errorf("can't finalize non-child of chain head")
		}

		// insert the number to block mapping
		err = operation.InsertNumber(header.View, header.ID())(tx)
		if err != nil {
			return fmt.Errorf("could not insert number mapping: %w", err)
		}

		// update the finalized boundary
		err = operation.UpdateBoundary(header.View)(tx)
		if err != nil {
			return fmt.Errorf("could not update finalized boundary: %w", err)
		}

		// NOTE: we don't want to prune forks that have become invalid here, so
		// that we can keep validating entities and generating slashing
		// challenges for some time - the pruning should happen some place else
		// after a certain delay of blocks

		return nil
	}
}

func Bootstrap(genesis *flow.Block) func(*badger.Txn) error {
	return func(tx *badger.Txn) error {

		// insert the block header & payload
		err := InsertBlock(genesis)(tx)
		if err != nil {
			return fmt.Errorf("could not insert genesis block: %w", err)
		}

		// apply the stake deltas
		err = ApplyDeltas(genesis.Height, genesis.Identities)(tx)
		if err != nil {
			return fmt.Errorf("could not apply stake deltas: %w", err)
		}

		// get first seal
		seal := genesis.Seals[0]

		// index the block seal
		err = operation.IndexSealIDByBlock(genesis.ID(), seal.ID())(tx)
		if err != nil {
			return fmt.Errorf("could not index seal by block: %w", err)
		}

		result := flow.ExecutionResult{ExecutionResultBody: flow.ExecutionResultBody{
			PreviousResultID: flow.ZeroID,
			BlockID:          genesis.ID(),
			FinalStateCommit: seal.FinalState,
		}}

		// index the commit for the execution node
		err = operation.IndexStateCommitment(genesis.ID(), seal.FinalState)(tx)
		if err != nil {
			return fmt.Errorf("could not index commit: %w", err)
		}

		// insert first execution result
		err = operation.InsertExecutionResult(&result)(tx)
		if err != nil {
			return fmt.Errorf("could not insert genesis result: %w", err)
		}

		// index first execution block for genesis block
		err = operation.IndexExecutionResult(genesis.ID(), result.ID())(tx)
		if err != nil {
			return fmt.Errorf("could not index genesis result: %w", err)
		}

		// insert the block number mapping
		err = operation.InsertNumber(0, genesis.ID())(tx)
		if err != nil {
			return fmt.Errorf("could not initialize boundary: %w", err)
		}

		// insert the finalized boundary
		err = operation.InsertBoundary(genesis.Height)(tx)
		if err != nil {
			return fmt.Errorf("could not update boundary: %w", err)
		}

		return nil
	}
}

func RetrieveLatestFinalizedHeader(header *flow.Header) func(tx *badger.Txn) error {
	var number uint64 = math.MaxUint64
	blockID := flow.ZeroID
	return RetrieveHeader(&number, &blockID, header)
}

func RetrieveHeader(number *uint64, blockID *flow.Identifier, header *flow.Header) func(tx *badger.Txn) error {
	return func(tx *badger.Txn) error {

		// set the number to boundary if it's at max uint64
		if *number == math.MaxUint64 {
			err := operation.RetrieveBoundary(number)(tx)
			if err != nil {
				return fmt.Errorf("could not retrieve boundary: %w", err)
			}
		}

		// check if hash is nil and try to get it from height
		if *blockID == flow.ZeroID {
			err := operation.RetrieveNumber(*number, blockID)(tx)
			if err != nil {
				return fmt.Errorf("could not retrieve hash (%d): %w", number, err)
			}
		}

		// get the height for our desired target hash
		err := operation.RetrieveHeader(*blockID, header)(tx)
		if err != nil {
			return fmt.Errorf("could not retrieve header (%x): %w", blockID, err)
		}

		return nil
	}
}

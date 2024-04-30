package postgres

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/gaze-network/indexer-network/common/errs"
	"github.com/gaze-network/indexer-network/core/types"
	"github.com/gaze-network/indexer-network/modules/bitcoin/repository/postgres/gen"
	"github.com/jackc/pgx/v5"
	"github.com/samber/lo"
)

func (r *Repository) GetLatestBlockHeader(ctx context.Context) (types.BlockHeader, error) {
	model, err := r.queries.GetLatestBlockHeader(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.BlockHeader{}, errors.Join(errs.NotFound, err)
		}
		return types.BlockHeader{}, errors.Wrap(err, "failed to get latest block header")
	}

	data, err := mapBlockHeaderModelToType(model)
	if err != nil {
		return types.BlockHeader{}, errors.Wrap(err, "failed to map block header model to type")
	}

	return data, nil
}

func (r *Repository) InsertBlocks(ctx context.Context, blocks []*types.Block) error {
	if len(blocks) == 0 {
		return nil
	}

	blockParams, txParams, txoutParams, txinParams := mapBlocksTypeToParams(blocks)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	defer tx.Rollback(ctx)

	queries := r.queries.WithTx(tx)

	if err := queries.BatchInsertBlocks(ctx, blockParams); err != nil {
		return errors.Wrap(err, "failed to batch insert block headers")
	}

	if err := queries.BatchInsertTransactions(ctx, txParams); err != nil {
		return errors.Wrap(err, "failed to batch insert transactions")
	}

	// Should insert txout first, then txin
	// Because txin references txout
	if err := queries.BatchInsertTransactionTxOuts(ctx, txoutParams); err != nil {
		return errors.Wrap(err, "failed to batch insert transaction txins")
	}

	if err := queries.BatchInsertTransactionTxIns(ctx, txinParams); err != nil {
		return errors.Wrap(err, "failed to batch insert transaction txins")
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (r *Repository) RevertBlocks(ctx context.Context, from int64) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}
	defer tx.Rollback(ctx)

	queries := r.queries.WithTx(tx)
	if err := queries.RevertData(ctx, int32(from)); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return errors.Wrap(err, "failed to revert data")
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (r *Repository) GetBlockHeaderByHeight(ctx context.Context, blockHeight int64) (types.BlockHeader, error) {
	blockModel, err := r.queries.GetBlockByHeight(ctx, int32(blockHeight))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return types.BlockHeader{}, errors.Join(errs.NotFound, err)
		}
		return types.BlockHeader{}, errors.Wrap(err, "failed to get block by height")
	}

	data, err := mapBlockHeaderModelToType(blockModel)
	if err != nil {
		return types.BlockHeader{}, errors.Wrap(err, "failed to map block header model to type")
	}
	return data, nil
}

func (r *Repository) GetBlocksByHeightRange(ctx context.Context, from int64, to int64) ([]*types.Block, error) {
	blocks, err := r.queries.GetBlocksByHeightRange(ctx, gen.GetBlocksByHeightRangeParams{
		FromHeight: int32(from),
		ToHeight:   int32(to),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get blocks by height range")
	}

	if len(blocks) == 0 {
		return []*types.Block{}, nil
	}

	txs, err := r.queries.GetTransactionsByHeightRange(ctx, gen.GetTransactionsByHeightRangeParams{
		FromHeight: int32(from),
		ToHeight:   int32(to),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transactions by height range")
	}

	txHashes := lo.Map(txs, func(tx gen.BitcoinTransaction, _ int) string { return tx.TxHash })

	txOuts, err := r.queries.GetTransactionTxOutsByTxHashes(ctx, txHashes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction txouts by tx hashes")
	}

	txIns, err := r.queries.GetTransactionTxInsByTxHashes(ctx, txHashes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transaction txins by tx hashes")
	}

	// Grouping result by block height and tx hash
	groupedTxs := lo.GroupBy(txs, func(tx gen.BitcoinTransaction) int32 { return tx.BlockHeight })
	groupedTxOuts := lo.GroupBy(txOuts, func(txOut gen.BitcoinTransactionTxout) string { return txOut.TxHash })
	groupedTxIns := lo.GroupBy(txIns, func(txIn gen.BitcoinTransactionTxin) string { return txIn.TxHash })

	var errs []error
	result := lo.Map(blocks, func(blockModel gen.BitcoinBlock, _ int) *types.Block {
		header, err := mapBlockHeaderModelToType(blockModel)
		if err != nil {
			errs = append(errs, errors.Wrap(err, "failed to map block header model to type"))
			return nil
		}

		txsModel := groupedTxs[blockModel.BlockHeight]
		return &types.Block{
			Header: header,
			Transactions: lo.Map(txsModel, func(txModel gen.BitcoinTransaction, _ int) *types.Transaction {
				tx, err := mapTransactionModelToType(txModel, groupedTxIns[txModel.TxHash], groupedTxOuts[txModel.TxHash])
				if err != nil {
					errs = append(errs, errors.Wrap(err, "failed to map transaction model to type"))
					return nil
				}
				return &tx
			}),
		}
	})
	if len(errs) > 0 {
		return nil, errors.Wrap(errors.Join(errs...), "failed while mapping result")
	}
	return result, nil
}
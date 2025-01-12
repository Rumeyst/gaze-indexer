// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: nodes.sql

package gen

import (
	"context"
)

const clearDelegate = `-- name: ClearDelegate :execrows
UPDATE nodes
SET "delegated_to" = ''
WHERE "delegate_tx_hash" = ''
`

func (q *Queries) ClearDelegate(ctx context.Context) (int64, error) {
	result, err := q.db.Exec(ctx, clearDelegate)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createNode = `-- name: CreateNode :exec
INSERT INTO nodes (sale_block, sale_tx_index, node_id, tier_index, delegated_to, owner_public_key, purchase_tx_hash, delegate_tx_hash)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type CreateNodeParams struct {
	SaleBlock      int64
	SaleTxIndex    int32
	NodeID         int32
	TierIndex      int32
	DelegatedTo    string
	OwnerPublicKey string
	PurchaseTxHash string
	DelegateTxHash string
}

func (q *Queries) CreateNode(ctx context.Context, arg CreateNodeParams) error {
	_, err := q.db.Exec(ctx, createNode,
		arg.SaleBlock,
		arg.SaleTxIndex,
		arg.NodeID,
		arg.TierIndex,
		arg.DelegatedTo,
		arg.OwnerPublicKey,
		arg.PurchaseTxHash,
		arg.DelegateTxHash,
	)
	return err
}

const getNodeCountByTierIndex = `-- name: GetNodeCountByTierIndex :many
SELECT (tiers.tier_index)::int AS tier_index, count(nodes.tier_index)
FROM generate_series($3::int,$4::int) AS tiers(tier_index)
LEFT JOIN 
	(SELECT sale_block, sale_tx_index, node_id, tier_index, delegated_to, owner_public_key, purchase_tx_hash, delegate_tx_hash 
	FROM nodes 
	WHERE sale_block = $1 AND 
		sale_tx_index= $2) 
	AS nodes ON tiers.tier_index = nodes.tier_index 
GROUP BY tiers.tier_index
ORDER BY tiers.tier_index
`

type GetNodeCountByTierIndexParams struct {
	SaleBlock   int64
	SaleTxIndex int32
	FromTier    int32
	ToTier      int32
}

type GetNodeCountByTierIndexRow struct {
	TierIndex int32
	Count     int64
}

func (q *Queries) GetNodeCountByTierIndex(ctx context.Context, arg GetNodeCountByTierIndexParams) ([]GetNodeCountByTierIndexRow, error) {
	rows, err := q.db.Query(ctx, getNodeCountByTierIndex,
		arg.SaleBlock,
		arg.SaleTxIndex,
		arg.FromTier,
		arg.ToTier,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetNodeCountByTierIndexRow
	for rows.Next() {
		var i GetNodeCountByTierIndexRow
		if err := rows.Scan(&i.TierIndex, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNodesByDeployment = `-- name: GetNodesByDeployment :many
SELECT sale_block, sale_tx_index, node_id, tier_index, delegated_to, owner_public_key, purchase_tx_hash, delegate_tx_hash 
FROM nodes
WHERE sale_block = $1 AND
    sale_tx_index = $2
`

type GetNodesByDeploymentParams struct {
	SaleBlock   int64
	SaleTxIndex int32
}

func (q *Queries) GetNodesByDeployment(ctx context.Context, arg GetNodesByDeploymentParams) ([]Node, error) {
	rows, err := q.db.Query(ctx, getNodesByDeployment, arg.SaleBlock, arg.SaleTxIndex)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.SaleBlock,
			&i.SaleTxIndex,
			&i.NodeID,
			&i.TierIndex,
			&i.DelegatedTo,
			&i.OwnerPublicKey,
			&i.PurchaseTxHash,
			&i.DelegateTxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNodesByIds = `-- name: GetNodesByIds :many
SELECT sale_block, sale_tx_index, node_id, tier_index, delegated_to, owner_public_key, purchase_tx_hash, delegate_tx_hash
FROM nodes
WHERE sale_block = $1 AND
    sale_tx_index = $2 AND
    node_id = ANY ($3::int[])
`

type GetNodesByIdsParams struct {
	SaleBlock   int64
	SaleTxIndex int32
	NodeIds     []int32
}

func (q *Queries) GetNodesByIds(ctx context.Context, arg GetNodesByIdsParams) ([]Node, error) {
	rows, err := q.db.Query(ctx, getNodesByIds, arg.SaleBlock, arg.SaleTxIndex, arg.NodeIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.SaleBlock,
			&i.SaleTxIndex,
			&i.NodeID,
			&i.TierIndex,
			&i.DelegatedTo,
			&i.OwnerPublicKey,
			&i.PurchaseTxHash,
			&i.DelegateTxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNodesByOwner = `-- name: GetNodesByOwner :many
SELECT sale_block, sale_tx_index, node_id, tier_index, delegated_to, owner_public_key, purchase_tx_hash, delegate_tx_hash 
FROM nodes
WHERE sale_block = $1 AND
    sale_tx_index = $2 AND
    owner_public_key = $3
ORDER BY tier_index
`

type GetNodesByOwnerParams struct {
	SaleBlock      int64
	SaleTxIndex    int32
	OwnerPublicKey string
}

func (q *Queries) GetNodesByOwner(ctx context.Context, arg GetNodesByOwnerParams) ([]Node, error) {
	rows, err := q.db.Query(ctx, getNodesByOwner, arg.SaleBlock, arg.SaleTxIndex, arg.OwnerPublicKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.SaleBlock,
			&i.SaleTxIndex,
			&i.NodeID,
			&i.TierIndex,
			&i.DelegatedTo,
			&i.OwnerPublicKey,
			&i.PurchaseTxHash,
			&i.DelegateTxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNodesByPubkey = `-- name: GetNodesByPubkey :many
SELECT nodes.sale_block, nodes.sale_tx_index, nodes.node_id, nodes.tier_index, nodes.delegated_to, nodes.owner_public_key, nodes.purchase_tx_hash, nodes.delegate_tx_hash
FROM nodes JOIN events ON nodes.purchase_tx_hash = events.tx_hash
WHERE sale_block = $1 AND
    sale_tx_index = $2 AND
    owner_public_key = $3 AND
    delegated_to = $4
`

type GetNodesByPubkeyParams struct {
	SaleBlock      int64
	SaleTxIndex    int32
	OwnerPublicKey string
	DelegatedTo    string
}

func (q *Queries) GetNodesByPubkey(ctx context.Context, arg GetNodesByPubkeyParams) ([]Node, error) {
	rows, err := q.db.Query(ctx, getNodesByPubkey,
		arg.SaleBlock,
		arg.SaleTxIndex,
		arg.OwnerPublicKey,
		arg.DelegatedTo,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Node
	for rows.Next() {
		var i Node
		if err := rows.Scan(
			&i.SaleBlock,
			&i.SaleTxIndex,
			&i.NodeID,
			&i.TierIndex,
			&i.DelegatedTo,
			&i.OwnerPublicKey,
			&i.PurchaseTxHash,
			&i.DelegateTxHash,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setDelegates = `-- name: SetDelegates :execrows
UPDATE nodes
SET delegated_to = $4, delegate_tx_hash = $3
WHERE sale_block = $1 AND
    sale_tx_index = $2 AND
    node_id = ANY ($5::int[])
`

type SetDelegatesParams struct {
	SaleBlock      int64
	SaleTxIndex    int32
	DelegateTxHash string
	Delegatee      string
	NodeIds        []int32
}

func (q *Queries) SetDelegates(ctx context.Context, arg SetDelegatesParams) (int64, error) {
	result, err := q.db.Exec(ctx, setDelegates,
		arg.SaleBlock,
		arg.SaleTxIndex,
		arg.DelegateTxHash,
		arg.Delegatee,
		arg.NodeIds,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

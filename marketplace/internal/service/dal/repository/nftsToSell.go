package repository

import (
	"context"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/service/dal/repository/model"
)

func (db PgConnection) Exec(ctx context.Context, query string, args []interface{}) (model.ExecResult, error) {
	var nftRecord model.ExecResult
	result, err := db.conn.ExecContext(ctx, query, args...)
	if err != nil {
		return nftRecord, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return nftRecord, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nftRecord, err
	}

	nftRecord.LastInsertId = lastInsertId
	nftRecord.RowsAffected = rowAffected
	return nftRecord, err
}

func (db PgConnection) ExecNftQuery(ctx context.Context, query string, args []interface{}) (model.NftToSell, error) {
	row := db.conn.QueryRowxContext(ctx, query, args...)
	var nftRecord model.NftToSell
	err := row.StructScan(&nftRecord)
	if err != nil {
		return nftRecord, err
	}

	return nftRecord, nil
}

func (db PgConnection) SelectNftQuery(ctx context.Context, query string, args []interface{}) ([]model.NftToSell, error) {
	rows, err := db.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	var accounts []model.NftToSell
	defer rows.Close()
	for rows.Next() {
		var account model.NftToSell
		err = rows.StructScan(&account)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

package dal

import (
	"context"
	"errors"
	"fmt"
	"github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/marketplace/internal/service/dal/repository/model"
	"strconv"
	"strings"
	"time"
)

type NftsToSellIntrf interface {
	Exec(ctx context.Context, query string, args ...interface{}) (model.ExecResult, error)
	ExecNftQuery(ctx context.Context, query string, args ...interface{}) (model.NftToSell, error)
	SelectNftQuery(ctx context.Context, query string, args ...interface{}) ([]model.NftToSell, error)
}

func (con DataAccessLayer) CreateNftRecordToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error) {
	args := []interface{}{
		input.Owner,
		input.ContractAddress,
		input.Creator,
		input.TokenId,
		input.TokenStandard,
		input.Status,
		input.CreatedAt,
		input.UpdatedAt,
	}
	query := `
		INSERT INTO public.nfts_to_sell (
		  tokenId,
		  owner,
		  contractAddress,
		  creator,
		  tokenStandard,
		  status,
		  createdAt,
		  updatedAt
		) VALUES (
		  $1, $2, $3, $4, $5, $6, $7, $8
		) RETURNING
		  tokenId,
		  owner,
		  contractAddress,
		  creator,
		  tokenStandard,
		  status,
		  createdAt,
		  updatedAt
	`

	resp, err := con.dbConnection.ExecNftQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// TODO IMPROVE QUERY TO AVOID CODE DUPLICATED
func (con DataAccessLayer) GetNftsToSell(ctx context.Context, input map[string]interface{}) ([]model.NftToSell, error) {
	var where []string

	query := `
			SELECT 
			  tokenId,
			  owner,
			  contractAddress,
			  creator,
			  tokenStandard,
			  status,
			  createdAt,
			  updatedAt
			FROM public.nfts_to_sell 
			WHERE 
	`

	if ids, ok := input["tokenId"]; ok {
		var strIds []string
		for _, value := range ids.([]int32) {
			strIds = append(strIds, strconv.Itoa(int(value)))
		}
		query += `tokenId IN (` + strings.Join(strIds, ", ") + `)`
		where = append(where, query)
	}

	if ids, ok := input["creator"]; ok {
		var strIds []string
		for _, value := range ids.([]int32) {
			strIds = append(strIds, strconv.Itoa(int(value)))
		}
		query += `creator IN (` + strings.Join(strIds, ", ") + `)`
		where = append(where, query)

	}

	whereStr := strings.Join(where, " OR ")
	query += whereStr

	limit := `30`
	if value, ok := input["limit"].(string); ok {
		limit = value
	}

	query += `LIMIT ` + limit

	if offset, ok := input["offset"].(string); ok {
		query += `
			OFFSET ` + offset
	}

	records, err := con.dbConnection.SelectNftQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (con DataAccessLayer) UpdateNftToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error) {
	var nftToSell model.NftToSell
	var args []interface{}
	var set []string

	if input.TokenId == "" {
		return nftToSell, errors.New("tokenId is a value needed")
	}

	if input.Status != "" {
		set = append(set, `status = $1`)
		args = append(args, input.Status)
	}

	if len(set) == 0 {
		return nftToSell, errors.New("there is not input fields to update")
	}

	updatedAt := time.Now().String()
	set = append(set, `updatedAt = $2`)
	args = append(args, updatedAt)

	args = append(args, input.TokenId)
	query := fmt.Sprintf(`
		UPDATE	
			public.nfts_to_sell
		SET
			%s
		WHERE
			tokenId = $3
		RETURNING
			  tokenId,
			  owner,
			  contractAddress,
			  creator,
			  tokenStandard,
			  status,
			  createdAt,
			  updatedAt
	`, strings.Join(set, ", "))

	resp, err := con.dbConnection.ExecNftQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (con DataAccessLayer) DeleteNftToSell(ctx context.Context, input model.NftToSell) (model.ExecResult, error) {
	var resp model.ExecResult
	if input.TokenId == "" {
		return resp, errors.New("tokenId is a value needed")
	}

	var args []interface{}
	args = append(args, input.TokenId)
	query := `
		DELETE FROM	
			public.nfts_to_sell
		WHERE
			 tokenId = $1
		`

	resp, err := con.dbConnection.Exec(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

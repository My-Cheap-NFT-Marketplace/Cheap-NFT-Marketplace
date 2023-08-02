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
	}

	query := `
		INSERT INTO public.nfts_to_sell (
		  token_id,
		  owner,
		  contract_address,
		  creator,
		  token_standard,
		  status
		) VALUES (
		  $1, $2, $3, $4, $5, $6
		) RETURNING
		  token_id,
		  owner,
		  contract_address,
		  creator,
		  token_standard,
		  status,
		  created_at,
		  updated_at
	`

	resp, err := con.dbConnection.ExecNftQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// TODO IMPROVE QUERY TO AVOID CODE DUPLICATED
func (con DataAccessLayer) GetNftsToSell(ctx context.Context, input model.QueryNft) ([]model.NftToSell, error) {
	var where []string
	var args []interface{}
	query := `
			SELECT 
			  token_id,
			  owner,
			  contract_address,
			  creator,
			  token_standard,
			  status,
			  created_at,
			  updated_at
			FROM public.nfts_to_sell
	`

	if input.TokenId != nil {
		args = append(args, *input.TokenId)
		tokenId := `token_id = $` + strconv.Itoa(len(args))
		where = append(where, tokenId)
	}

	if input.Owner != nil {
		args = append(args, *input.Owner)
		owner := `owner = $` + strconv.Itoa(len(args))
		where = append(where, owner)
	}

	whereStr := strings.Join(where, " AND ")
	query += `WHERE ` + whereStr

	if input.Limit != nil {
		query += `LIMIT ` + *input.Limit
	}

	if input.Offset != nil {
		query += `OFFSET ` + *input.Offset
	}

	records, err := con.dbConnection.SelectNftQuery(ctx, query, args)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (con DataAccessLayer) UpdateNftToSell(ctx context.Context, input model.NftToSell) (model.NftToSell, error) {
	var nftToSell model.NftToSell
	var args []interface{}
	var set []string
	var where []string

	if input.TokenId == nil {
		return nftToSell, errors.New("tokenId is a value needed")
	}

	if input.Status == nil {
		args = append(args, *input.Status)
		set = append(set, `status = $`+strconv.Itoa(len(args)))
	}

	if len(set) == 0 {
		return nftToSell, errors.New("there is not input fields to update")
	}

	updatedAt := time.Now().String()
	args = append(args, updatedAt)
	set = append(set, `updated_at = $`+strconv.Itoa(len(args)))

	args = append(args, *input.TokenId)
	where = append(where, `token_id = $`+strconv.Itoa(len(args)))
	query := fmt.Sprintf(`
		UPDATE	
			public.nfts_to_sell
		SET
			%s
		WHERE
			%s
		RETURNING
		  token_id,
		  owner,
		  contract_address,
		  creator,
		  token_standard,
		  status,
		  created_at,
		  updated_at
	`, strings.Join(set, ", "), strings.Join(where, " AND "))

	resp, err := con.dbConnection.ExecNftQuery(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (con DataAccessLayer) DeleteNftToSell(ctx context.Context, input model.NftToSell) (model.ExecResult, error) {
	var resp model.ExecResult
	if input.TokenId == nil {
		return resp, errors.New("tokenId is a value needed")
	}

	var args []interface{}
	args = append(args, *input.TokenId)
	query := `
		DELETE FROM	
			public.nfts_to_sell
		WHERE
			 token_id = $1
		`

	resp, err := con.dbConnection.Exec(ctx, query, args)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

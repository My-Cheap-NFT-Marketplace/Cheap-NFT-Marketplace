package dal

import (
	"context"
	"errors"
	"fmt"
	dbModel "github.com/My-Cheap-NFT-Marketplace/Cheap-NFT-Marketplace/transaction-eventing/internal/messaging/topic/dal/repository/model"
	"strconv"
	"strings"
	"time"
)

func (con DalOrder) CreateOrderToSell(ctx context.Context, input dbModel.OrderItem) error {
	args := []interface{}{
		input.Trx,
		input.TokenId,
		input.Owner,
		input.ContractAddress,
		input.Creator,
		input.TokenStandard,
		input.Status,
	}

	query := `
		INSERT INTO public.ntf_orders (
		  trx,
		  token_id,
		  owner,
		  contract_address,
		  creator,
		  token_standard,
		  status
		) VALUES (
		  $1, $2, $3, $4, $5, $6, $7
		) RETURNING
		  trx,
		  token_id,
		  owner,
		  contract_address,
		  creator,
		  token_standard,
		  status,
		  created_at,
		  updated_at
	`

	_, err := con.dbConnection.ExecOrderQuery(ctx, query, args)
	return err
}

func (con DalOrder) UpdateStatusOrder(ctx context.Context, input dbModel.OrderItem) error {
	var args []interface{}
	var set []string
	var where []string

	if input.Trx == nil {
		return errors.New("trx is a value needed")
	}

	if input.Status == nil {
		return errors.New("status is a value needed")
	}

	args = append(args, *input.Status)
	set = append(set, `status = $`+strconv.Itoa(len(args)))

	updatedAt := time.Now().String()
	args = append(args, updatedAt)
	set = append(set, `updated_at = $`+strconv.Itoa(len(args)))

	args = append(args, *input.Trx)
	where = append(where, `trx = $`+strconv.Itoa(len(args)))
	query := fmt.Sprintf(`
		UPDATE	
			public.ntf_orders
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

	_, err := con.dbConnection.ExecOrderQuery(ctx, query, args)
	return err
}

package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/neel229/forum/pkg/util"
)

// Store provides all the functions to execute
// db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function to perform
// db transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); err != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// StockCreationTxParams is a struct
// which holds input data to create a creator
// account and it's corresponding stock
type StockCreationTxParams struct {
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	UserName            string `json:"user_name"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	PreferredCurrencyID int64  `json:"preferred_currency_id"`
	Ticker              string `json:"ticker"`
	Details             string `json:"details"`
}

// StockCreationTxResults contains
// a Creator account, corresponding Stock
// and CreatorStock
type StockCreationTxResults struct {
	Creator      Creator          `json:"creator"`
	Stock        Stock            `json:"stock"`
	CreatorStock CreatorStock     `json:"creator_stock"`
	CPortfolio   CreatorPortfolio `json:"creator_portfolioj"`
}

// StockCreationTx is a db transaction
// in which we create a creator's id
// followed by a corresponding stock
func (s *Store) StockCreationTx(ctx context.Context, arg StockCreationTxParams) (StockCreationTxResults, error) {
	var result StockCreationTxResults

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		// Create a creator account
		arg1 := CreateCreatorParams{
			FirstName:           arg.FirstName,
			LastName:            arg.LastName,
			UserName:            arg.UserName,
			Email:               arg.Email,
			Password:            arg.Password,
			PreferredCurrencyID: arg.PreferredCurrencyID,
			VirginTokensLeft:    10000000,
		}
		result.Creator, err = q.CreateCreator(ctx, arg1)
		if err != nil {
			return err
		}

		mintPrice := util.RandomInt32(10, 50)

		// Create a creator's stock
		arg2 := CreateStockParams{
			Ticker:       arg.Ticker,
			Details:      arg.Details,
			MintPrice:    mintPrice,
			CurrentPrice: mintPrice,
		}
		result.Stock, err = q.CreateStock(ctx, arg2)
		if err != nil {
			return err
		}

		// Create an entry in  creator stock table
		// with creatorID and stockID from above two
		// queries
		arg3 := CreateCreatorStockParams{
			CreatorID: result.Creator.ID,
			StockID:   result.Stock.ID,
		}
		result.CreatorStock, err = q.CreateCreatorStock(ctx, arg3)
		if err != nil {
			return err
		}

		arg4 := CreateCreatorPortfolioParams{
			CreatorID: result.Creator.ID,
			StockID:   result.Stock.ID,
			Quantity:  "10000000.000000",
		}
		result.CPortfolio, err = q.CreateCreatorPortfolio(ctx, arg4)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

// VTradeTxParams holds the input
// parameters for the VTradeTx
type VTradeTxParams struct {
	StockID       int64  `json:"stock_id"`
	CreatorID     int64  `json:"creator_id"`
	FanID         int64  `json:"fan_id"`
	Quantity      string `json:"quantity"`
	UnitPrice     string `json:"unit_price"`
	Details       string `json:"details"`
	VirginOfferID int64  `json:"virgin_offer_id"`
}

// VTradeTxResult holds the
// table records which are created/
// updated after successful VTradeTx
type VTradeTxResult struct {
	VTrade VirginTrade `json:"virgin_trade"`
}

// VTradeTx is a db transaction
// in which a fan buys some tokens
// from the creator.
func (s *Store) VTradeTx(ctx context.Context, arg VTradeTxParams) (VTradeTxResult, error) {
	var result VTradeTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		// Create a virgin trade entry corresponding to the transaction
		arg1 := CreateVirginTradeParams{
			StockID:       arg.StockID,
			CreatorID:     arg.CreatorID,
			FanID:         arg.FanID,
			Quantity:      arg.Quantity,
			UnitPrice:     arg.UnitPrice,
			Details:       arg.Details,
			VirginOfferID: arg.VirginOfferID,
		}
		result.VTrade, err = q.CreateVirginTrade(ctx, arg1)
		if err != nil {
			return err
		}

		// Deduct the stock quantity from creator's account & portfolio
		tokensLeft, _ := q.GetVirginTokensLeft(ctx, arg.CreatorID)
		quantity, _ := strconv.ParseFloat(arg.Quantity, 64)
		arg2 := UpdateVirginTokensLeftParams{
			ID:               arg.CreatorID,
			VirginTokensLeft: tokensLeft - int32(quantity),
		}
		if err = s.UpdateVirginTokensLeft(ctx, arg2); err != nil {
			return err
		}

		creator, _ := q.GetCreator(ctx, arg.CreatorID)
		arg3 := UpdateCreatorStockQuantityParams{
			CreatorID: arg.CreatorID,
			StockID:   arg.StockID,
			Quantity:  strconv.Itoa(int(creator.VirginTokensLeft)),
		}
		if err = s.UpdateCreatorStockQuantity(ctx, arg3); err != nil {
			return err
		}

		// Add the stock quantity in fan's portfolio
		portfolio, _ := q.GetPortfolioByFanID(ctx, arg.FanID)
		if portfolio.StockID == arg.StockID {
			quantity1, _ := strconv.ParseFloat(portfolio.Quantity, 64)
			quantity2, _ := strconv.ParseFloat(arg.Quantity, 64)
			quantity3 := quantity1 + quantity2
			quantity := strconv.FormatFloat(quantity3, 'e', 6, 64)
			arg4 := UpdateFanStockQuantityParams{
				FanID:    arg.FanID,
				StockID:  arg.StockID,
				Quantity: quantity,
			}
			if err = s.UpdateFanStockQuantity(ctx, arg4); err != nil {
				return err
			}
		} else {
			arg5 := CreateFanPortfolioParams{
				FanID:    arg.FanID,
				StockID:  arg.StockID,
				Quantity: arg.Quantity,
			}
			if _, err = s.CreateFanPortfolio(ctx, arg5); err != nil {
				return err
			}
		}
		return nil
	})

	return result, err
}

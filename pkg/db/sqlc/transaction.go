package db

import (
	"context"
	"database/sql"
	"fmt"
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
	PreferredCurrencyID int32  `json:"preferred_currency_id"`
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

		// Create a creator's stock
		arg2 := CreateStockParams{
			Ticker:  arg.Ticker,
			Details: arg.Details,
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

// // VTradeTransactionParams contains the
// // necessery input params
// type VTradeTransactionParams struct {
// 	CreatorID int64  `json:"creator_id"`
// 	FanID     int64  `json:"fan_id"`
// 	StockID   int64  `json:"stock_id"`
// 	Quantity  string `json:"quantity"`
// 	UnitPrice string `json:"unit_price"`
// 	// VirginOfferID int64  `json:"virgin_offer_id"`
// }

// // VTradeTransactionResult contains the
// // fields which will be changed after the
// // transaction.
// type VTradeTransactionResult struct {
// 	VTrade     VirginTrade `json:"virgin_trade"`
// 	CAccount   Creator     `json:"creator_account"`
// 	CPortfolio Portfolio   `json:"creator_portfolio"`
// 	FPortfolio Portfolio   `json:"fan_portfolio"`
// }

// // VTradeTransaction performs a stock trade
// // from a creator's account to a fan account.
// func (s *Store) VTradeTransaction(ctx context.Context, arg VTradeTransactionParams) (VTradeTransactionResult, error) {
// 	var result VTradeTransactionResult

// 	err := s.execTx(ctx, func(q *Queries) error {
// 		var err error
// 		vOffer, _ := q.GetVirginOfferByCreator(ctx, arg.CreatorID) // Get the virgin offer using creator id.
// 		arg1 := CreateVirginTradeParams{
// 			StockID:       arg.StockID,
// 			CreatorID:     arg.CreatorID,
// 			BuyerID:       arg.FanID,
// 			Quantity:      arg.Quantity,
// 			UnitPrice:     arg.UnitPrice,
// 			Details:       "",
// 			VirginOfferID: vOffer.ID,
// 		}
// 		result.VTrade, err = q.CreateVirginTrade(ctx, arg1)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})

// 	return result, err
// }

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

// VTradeTransactionParams contains the
// necessery input params
type VTradeTransactionParams struct {
	CreatorID int64  `json:"creator_id"`
	FanID     int64  `json:"fan_id"`
	StockID   int64  `json:"stock_id"`
	Quantity  string `json:"quantity"`
	UnitPrice string `json:"unit_price"`
	// VirginOfferID int64  `json:"virgin_offer_id"`
}

// VTradeTransactionResult contains the
// fields which will be changed after the
// transaction.
type VTradeTransactionResult struct {
	VTrade     VirginTrade `json:"virgin_trade"`
	CAccount   Creator     `json:"creator_account"`
	CPortfolio Portfolio   `json:"creator_portfolio"`
	FPortfolio Portfolio   `json:"fan_portfolio"`
}

// VTradeTransaction performs a stock trade
// from a creator's account to a fan account.
func (s *Store) VTradeTransaction(ctx context.Context, arg VTradeTransactionParams) (VTradeTransactionResult, error) {
	var result VTradeTransactionResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		vOffer, _ := q.GetVirginOfferByCreator(ctx, arg.CreatorID) // Get the virgin offer using creator id.
		arg1 := CreateVirginTradeParams{
			StockID:       arg.StockID,
			CreatorID:     arg.CreatorID,
			BuyerID:       arg.FanID,
			Quantity:      arg.Quantity,
			UnitPrice:     arg.UnitPrice,
			Details:       "",
			VirginOfferID: vOffer.ID,
		}
		result.VTrade, err = q.CreateVirginTrade(ctx, arg1)
		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}

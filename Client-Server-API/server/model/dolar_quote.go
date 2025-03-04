package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

const (
	queryInsert = `INSERT INTO cotacao (bid, created_date) VALUES (?, ?)`
)

type DolarQuote struct {
	bid         float64
	createdDate string
}

func NewDolarQuote(bid float64) *DolarQuote {
	cd := time.Now().Format("2006-01-02 15:04:05")
	return &DolarQuote{
		bid:         bid,
		createdDate: cd,
	}
}

func (d *DolarQuote) CreateDolarQuote(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, queryInsert, d.bid, d.createdDate)
	if err != nil {
		return fmt.Errorf("Error to insert new dolar quote at database: %v", err)
	}
	return nil
}

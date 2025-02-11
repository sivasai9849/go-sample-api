package postgres

import (
	"context"

	"gorm.io/gorm"
)

type Transaction interface {
    Begin() error
    Commit() error
    Rollback() error
}

type txKey struct{}

func WithTransaction(ctx context.Context, db *gorm.DB) (context.Context, *gorm.DB) {
    tx := db.Begin()
    return context.WithValue(ctx, txKey{}, tx), tx
}
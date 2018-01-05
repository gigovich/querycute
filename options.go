package querycute

import (
	"context"
	"database/sql"
)

// WithTx set transaction to query object
func WithTx(tx *sql.Tx) OptionFunc {
	return func(q *Query) error {
		q.tx = tx
		return nil
	}
}

// WithCtx set context to query object
func WithCtx(ctx context.Context) OptionFunc {
	return func(q *Query) error {
		q.ctx = ctx
		return nil
	}
}

// OptionFunc configure query
type OptionFunc func(q *Query) error

package dollarquoterepository

import (
	"context"
	"time"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/db/dollarquotedb"
)

type Repository interface {
	SaveDollarQuote(ctx context.Context, entity dollarquote.DollarQuote) (*dollarquote.DollarQuote, error)
	FindDollarQuoteByID(ctx context.Context, id int64) (*dollarquote.DollarQuote, error)
}

type repository struct {
	db dollarquotedb.Client
}

func New(db dollarquotedb.Client) Repository {
	return &repository{db: db}
}

func (r *repository) SaveDollarQuote(ctx context.Context, dollarQuote dollarquote.DollarQuote) (*dollarquote.DollarQuote, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	entity, err := r.db.SaveDollarQuote(ctxTimeout, dollarquotedb.NewDollarQuoteEntity(dollarQuote))
	if err != nil {
		return nil, err
	}

	return entity.ToDollarQuote(), nil
}

func (r *repository) FindDollarQuoteByID(ctx context.Context, id int64) (*dollarquote.DollarQuote, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	entity, err := r.db.FindDollarQuoteByID(ctxTimeout, id)
	if err != nil {
		return nil, err
	}

	return entity.ToDollarQuote(), nil
}

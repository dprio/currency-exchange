package dollarquoterepository

import (
	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/db/dollarquotedb"
)

type Repository interface {
	SaveDollarQuote(entity dollarquote.DollarQuote) (*dollarquote.DollarQuote, error)
	FindDollarQuoteByID(id int64) (*dollarquote.DollarQuote, error)
}

type repository struct {
	db dollarquotedb.Client
}

func New(db dollarquotedb.Client) Repository {
	return &repository{db: db}
}

func (r *repository) SaveDollarQuote(dollarQuote dollarquote.DollarQuote) (*dollarquote.DollarQuote, error) {
	entity, err := r.db.SaveDollarQuote(dollarquotedb.NewDollarQuoteEntity(dollarQuote))
	if err != nil {
		return nil, err
	}

	return entity.ToDollarQuote(), nil
}

func (r *repository) FindDollarQuoteByID(id int64) (*dollarquote.DollarQuote, error) {
	entity, err := r.db.FindDollarQuoteByID(id)
	if err != nil {
		return nil, err
	}

	return entity.ToDollarQuote(), nil
}

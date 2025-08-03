package dollarquoterepository

import (
	"context"
	"fmt"
	"time"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarexchangerate"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/db/dollarexchangeratedb"
)

type Repository interface {
	SaveDollarExchangeRate(ctx context.Context, entity dollarexchangerate.DollarExchangeRate) (*dollarexchangerate.DollarExchangeRate, error)
	FindDollarExchangeRateByID(ctx context.Context, id int64) (*dollarexchangerate.DollarExchangeRate, error)
}

type repository struct {
	db dollarexchangeratedb.Client
}

func New(db dollarexchangeratedb.Client) Repository {
	return &repository{db: db}
}

func (r *repository) SaveDollarExchangeRate(ctx context.Context, exchangeRate dollarexchangerate.DollarExchangeRate) (*dollarexchangerate.DollarExchangeRate, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	entity, err := r.db.SaveDollarExchangeRate(ctxTimeout, dollarexchangeratedb.NewDollarQuoteEntity(exchangeRate))
	if err != nil {
		fmt.Printf("error saving dollar quote. [Error: %s]\n", err.Error())
		return nil, err
	}

	return entity.ToDollarExchnageRate(), nil
}

func (r *repository) FindDollarExchangeRateByID(ctx context.Context, id int64) (*dollarexchangerate.DollarExchangeRate, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	entity, err := r.db.FindDollarExchangeRate(ctxTimeout, id)
	if err != nil {
		fmt.Printf("error finding dollar exchange rate. [Error: %s]\n", err.Error())
		return nil, err
	}

	return entity.ToDollarExchnageRate(), nil
}

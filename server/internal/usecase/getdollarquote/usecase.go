package getdollarquote

import (
	"context"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarquote"
)

type UseCase interface {
	Execute(ctx context.Context) (*dollarquote.DollarQuote, error)
}

type (
	dollarQuoteHTTPClientGateway interface {
		GetUSDQuote(ctx context.Context) (*dollarquote.DollarQuote, error)
	}

	repositoy interface {
		SaveDollarQuote(ctx context.Context, entity dollarquote.DollarQuote) (*dollarquote.DollarQuote, error)
	}

	useCase struct {
		dollarQuoteHTTPClientGateway dollarQuoteHTTPClientGateway
		repositoy                    repositoy
	}
)

func New(dollarQuoteHTTPClientGateway dollarQuoteHTTPClientGateway) UseCase {
	return &useCase{
		dollarQuoteHTTPClientGateway: dollarQuoteHTTPClientGateway,
	}
}

func (u *useCase) Execute(ctx context.Context) (*dollarquote.DollarQuote, error) {
	dollarQuote, err := u.dollarQuoteHTTPClientGateway.GetUSDQuote(ctx)
	if err != nil {
		return nil, err
	}

	return u.repositoy.SaveDollarQuote(ctx, *dollarQuote)
}

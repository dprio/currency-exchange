package getdollarexchangerate

import (
	"context"
	"fmt"

	"github.com/dprio/currency-exchange/server/internal/domain/dollarexchangerate"
)

type UseCase interface {
	Execute(ctx context.Context) (*dollarexchangerate.DollarExchangeRate, error)
}

type (
	economiaAPIClientGateway interface {
		GetUSDExchangeRate(ctx context.Context) (*dollarexchangerate.DollarExchangeRate, error)
	}

	repositoy interface {
		SaveDollarExchangeRate(ctx context.Context, entity dollarexchangerate.DollarExchangeRate) (*dollarexchangerate.DollarExchangeRate, error)
	}

	useCase struct {
		dollarQuoteHTTPClientGateway economiaAPIClientGateway
		repositoy                    repositoy
	}
)

func New(dollarQuoteHTTPClientGateway economiaAPIClientGateway, repo repositoy) UseCase {
	return &useCase{
		dollarQuoteHTTPClientGateway: dollarQuoteHTTPClientGateway,
		repositoy:                    repo,
	}
}

func (u *useCase) Execute(ctx context.Context) (*dollarexchangerate.DollarExchangeRate, error) {

	fmt.Println("executing getDollarQuoteUseCase")
	dollarQuote, err := u.dollarQuoteHTTPClientGateway.GetUSDExchangeRate(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("dollarQuote recovered. DollarQuote: [%+v]\n", dollarQuote)
	return u.repositoy.SaveDollarExchangeRate(ctx, *dollarQuote)
}

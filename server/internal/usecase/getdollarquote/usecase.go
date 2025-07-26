package getdollarquote

import (
	"context"
	"fmt"

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

func New(dollarQuoteHTTPClientGateway dollarQuoteHTTPClientGateway, repo repositoy) UseCase {
	return &useCase{
		dollarQuoteHTTPClientGateway: dollarQuoteHTTPClientGateway,
		repositoy:                    repo,
	}
}

func (u *useCase) Execute(ctx context.Context) (*dollarquote.DollarQuote, error) {

	fmt.Println("executing getDollarQuoteUseCase")
	dollarQuote, err := u.dollarQuoteHTTPClientGateway.GetUSDQuote(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Printf("dollarQuote recovered. DollarQuote: [%+v]\n", dollarQuote)
	return u.repositoy.SaveDollarQuote(ctx, *dollarQuote)
}

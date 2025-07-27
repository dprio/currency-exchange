package getdollarquote

import (
	"context"
	"fmt"
)

type UseCase interface {
	Execute(ctx context.Context) (float32, error)
}

type (
	dollarQuoteGateway interface {
		GetDollarQuote(ctx context.Context) (float32, error)
	}

	usecase struct {
		dollarQuoteGateway dollarQuoteGateway
	}
)

func New(dollarQuoteGateway dollarQuoteGateway) UseCase {
	return &usecase{
		dollarQuoteGateway: dollarQuoteGateway,
	}
}

func (u *usecase) Execute(ctx context.Context) (float32, error) {
	quote, err := u.dollarQuoteGateway.GetDollarQuote(ctx)
	if err != nil {
		fmt.Printf("error recovering dolla quote. [Error: %s]\n", err.Error())
		return 0, err
	}

	return quote, nil
}

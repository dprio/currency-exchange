package getdollarexchangerate

import (
	"context"
	"fmt"
	"os"
)

type UseCase interface {
	Execute(ctx context.Context) (float32, error)
}

type (
	economiaAPIGateway interface {
		GetDollarExchangeRate(ctx context.Context) (float32, error)
	}

	usecase struct {
		economiaAPIGateway economiaAPIGateway
	}
)

func New(economiaAPIGateway economiaAPIGateway) UseCase {
	return &usecase{
		economiaAPIGateway: economiaAPIGateway,
	}
}

func (u *usecase) Execute(ctx context.Context) (float32, error) {
	quote, err := u.economiaAPIGateway.GetDollarExchangeRate(ctx)
	if err != nil {
		fmt.Printf("error recovering dolla exchange rate. [Error: %s]\n", err.Error())
		return 0, err
	}

	f, err := os.OpenFile("arquivo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error creating file. [Error: %s]\n", err.Error())
		return 0, err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "Dólar: %f\n", quote)
	if err != nil {
		fmt.Printf("error writing to file. [Error: %s]\n", err.Error())
		return 0, err
	}

	return quote, nil
}

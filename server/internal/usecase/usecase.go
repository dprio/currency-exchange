package usecase

import (
	infrastructure "github.com/dprio/currency-exchange/server/internal/infrastructure/httpclient"
	"github.com/dprio/currency-exchange/server/internal/usecase/getdollarquote"
)

type UseCase struct {
	GetDollarQuoteUseCase getdollarquote.UseCase
}

func New(httpClient infrastructure.HTTPClient) UseCase {
	return UseCase{
		GetDollarQuoteUseCase: getdollarquote.New(httpClient.DollarQuoteHTTPClient),
	}
}

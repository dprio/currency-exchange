package repository

import (
	"github.com/dprio/currency-exchange/server/internal/gateway/repository/dollarquoterepository"
	"github.com/dprio/currency-exchange/server/internal/infrastructure/db"
)

type Repositories struct {
	DollarQuoteRepository dollarquoterepository.Repository
}

func New(db db.DB) Repositories {
	return Repositories{
		DollarQuoteRepository: dollarquoterepository.New(db.DollarQuoteClient),
	}
}

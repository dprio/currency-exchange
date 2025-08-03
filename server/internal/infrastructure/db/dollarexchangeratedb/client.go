package dollarexchangeratedb

import (
	"context"
	"database/sql"
)

type Client interface {
	SaveDollarExchangeRate(ctx context.Context, entity DollarExchangeRateEntity) (*DollarExchangeRateEntity, error)
	FindDollarExchangeRate(ctx context.Context, id int64) (*DollarExchangeRateEntity, error)
}

type client struct {
	db *sql.DB
}

func NewClient(db *sql.DB) Client {
	return &client{db: db}
}

func (c *client) SaveDollarExchangeRate(ctx context.Context, entity DollarExchangeRateEntity) (*DollarExchangeRateEntity, error) {
	query := "INSERT INTO dollar_exchange_rate (value, created_at) VALUES (?, ?)"
	result, err := c.db.ExecContext(ctx, query, entity.Value, entity.CreatedAt)
	if err != nil {
		return nil, err
	}

	resp := entity
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	resp.ID = &id

	return &resp, nil
}

func (c *client) FindDollarExchangeRate(ctx context.Context, id int64) (*DollarExchangeRateEntity, error) {
	query := "SELECT id, value, created_at FROM dollar_exchange_rate WHERE id = ?"
	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var resp DollarExchangeRateEntity
	err = row.Scan(&resp.ID, &resp.Value, &resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

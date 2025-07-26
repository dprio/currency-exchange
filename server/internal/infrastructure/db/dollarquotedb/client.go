package dollarquotedb

import (
	"context"
	"database/sql"
)

type Client interface {
	SaveDollarQuote(ctx context.Context, entity DollarQuoteEntity) (*DollarQuoteEntity, error)
	FindDollarQuoteByID(ctx context.Context, id int64) (*DollarQuoteEntity, error)
}

type client struct {
	db *sql.DB
}

func NewClient(db *sql.DB) Client {
	return &client{db: db}
}

func (c *client) SaveDollarQuote(ctx context.Context, entity DollarQuoteEntity) (*DollarQuoteEntity, error) {
	query := "INSERT INTO dollar_quotes (value, created_at) VALUES (?, ?)"
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

func (c *client) FindDollarQuoteByID(ctx context.Context, id int64) (*DollarQuoteEntity, error) {
	query := "SELECT id, value, created_at FROM dollar_quotes WHERE id = ?"
	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	var resp DollarQuoteEntity
	err = row.Scan(&resp.ID, &resp.Value, &resp.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

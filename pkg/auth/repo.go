package auth

import (
	"context"
	"database/sql"
)

func getUserByEmail(ctx context.Context, db *sql.DB, userName string) (userItems, error) {
	query := `SELECT id, username, password,created_at FROM users WHERE username=?`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return userItems{}, err
	}
	defer stmt.Close()

	var data userItems
	if err := stmt.QueryRowContext(ctx, userName).Scan(
		&data.UserID,
		&data.Username,
		&data.Password,
		&data.CreatedAt,
	); err != nil {
		return userItems{}, err
	}
	return data, nil
}

func getProduct(ctx context.Context, db *sql.DB) ([]productItems, error) {
	getCount := `SELECT COUNT(*) FROM products`
	count := 0
	if err := db.QueryRowContext(ctx, getCount).Scan(&count); err != nil {
		return nil, err
	}
	buffer := make([]productItems, 0, count)

	query := `SELECT id, name,stock,created_at FROM products`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data productItems
	for rows.Next() {
		if err := rows.Scan(
			&data.ProductId,
			&data.Name,
			&data.Stock,
			&data.CreatedAt,
		); err != nil {
			return nil, err
		}
		buffer = append(buffer, data)
	}
	return buffer, nil
}

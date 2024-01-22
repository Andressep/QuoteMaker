package repository

import (
	"context"

	domain "github.com/Andressep/QuoteMaker/internal/core/domain/entity"
	"github.com/Andressep/QuoteMaker/internal/core/ports"
	"github.com/jmoiron/sqlx"
)

type sqlSellerRepository struct {
	db *sqlx.DB
}

const saveSellerQuery = `
INSERT INTO seller (name)
VALUES ($1)
RETURNING id, name;
`

// SaveSeller implements ports.SellerRepository.
func (r *sqlSellerRepository) SaveSeller(ctx context.Context, args domain.Seller) (domain.Seller, error) {
	row := r.db.QueryRowContext(ctx, saveSellerQuery, args.Name)
	var i domain.Seller

	err := row.Scan(
		&i.ID,
		&i.Name,
	)
	return i, err
}

const getSellerByID = `
SELECT id, name
FROM seller
WHERE id = $1;
`

// GetSellerByID implements ports.SellerRepository.
func (r *sqlSellerRepository) GetSellerByID(ctx context.Context, id int) (*domain.Seller, error) {
	row := r.db.QueryRowContext(ctx, getSellerByID, id)
	var i domain.Seller

	err := row.Scan(
		&i.ID,
		&i.Name,
	)
	return &i, err
}

const deleteSellerQuery = `
DELETE FROM seller
WHERE id = $1;
`

// DeleteSeller implements ports.SellerRepository.
func (r *sqlSellerRepository) DeleteSeller(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteSellerQuery, id)
	return err
}

const listSellersQuery = `
SELECT id, name
FROM seller
ORDER BY id
LIMIT $1 OFFSET $2;
`

// ListSellers implements ports.SellerRepository.
func (r *sqlSellerRepository) ListSellers(ctx context.Context, limit int, offset int) ([]domain.Seller, error) {
	var sellers []domain.Seller
	err := r.db.SelectContext(ctx, &sellers, listSellersQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	return sellers, nil
}

func NewSellerRepository(db *sqlx.DB) ports.SellerRepository {
	return &sqlSellerRepository{
		db: db,
	}
}

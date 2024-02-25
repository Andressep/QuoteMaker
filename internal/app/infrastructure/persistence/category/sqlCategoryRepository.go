package category

import (
	"context"
	"database/sql"

	domain "github.com/Andressep/QuoteMaker/internal/app/domain/category"
)

type sqlCategoryRepository struct {
	db *sql.DB
}

const saveCategoryQuery = `
INSERT INTO category (category_name)
VALUES ($1)
RETURNING id, category_name;
`

// SaveCategory implements ports.CategoryRepository.
func (r *sqlCategoryRepository) SaveCategory(ctx context.Context, args domain.Category) (domain.Category, error) {
	row := r.db.QueryRowContext(ctx, saveCategoryQuery, args.CategoryName)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)
	return i, err
}

const updateCategoryQuery = `
UPDATE category
SET category_name = $1
WHERE id = $2;
`

// UpdateCategory implements ports.CategoryRepository.
func (r *sqlCategoryRepository) UpdateCategory(ctx context.Context, category domain.Category) error {
	_, err := r.db.ExecContext(ctx, updateCategoryQuery, category.CategoryName, category.ID)
	if err != nil {
		return err
	}
	return nil
}

const getCategoryByIDQuery = `
SELECT id, category_name
FROM category
WHERE id = $1;
`

// GetCategoryByID implements ports.CategoryRepository.
func (r *sqlCategoryRepository) GetCategoryByID(ctx context.Context, id int) (*domain.Category, error) {
	row := r.db.QueryRowContext(ctx, getCategoryByIDQuery, id)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)
	return &i, err
}

const getCategoryByNameQuery = `
SELECT id, category_name
FROM category
WHERE category_name = $1;
`

// GetCategoryByName implements ports.CategoryRepository.
func (r *sqlCategoryRepository) GetCategoryByName(ctx context.Context, name string) (domain.Category, error) {
	row := r.db.QueryRowContext(ctx, getCategoryByNameQuery, name)
	var i domain.Category

	err := row.Scan(
		&i.ID,
		&i.CategoryName,
	)

	return i, err
}

const listCategoryQuery = `
SELECT id, category_name
FROM category
ORDER BY id
LIMIT $1 OFFSET $2;
`

// ListCategorys implements ports.CategoryRepository.
func (r *sqlCategoryRepository) ListCategorys(ctx context.Context, limit int, offset int) ([]domain.Category, error) {
	rows, err := r.db.QueryContext(ctx, listCategoryQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorys []domain.Category
	for rows.Next() {
		var i domain.Category
		if err := rows.Scan(
			&i.ID,
			&i.CategoryName); err != nil {
			return nil, err
		}
		categorys = append(categorys, i)
	}

	// Verificar por errores al finalizar la iteración
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categorys, nil
}

const deleteCategoryQuery = `
DELETE FROM category
WHERE id = $1;
`

// DeleteCategory implements ports.CategoryRepository.
func (r *sqlCategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, deleteCategoryQuery, id)
	return err
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &sqlCategoryRepository{
		db: db,
	}
}

package domain

import (
	"time"
)

type Quotation struct {
	ID          int        `db:"id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	TotalPrice  float64    `db:"total_price"`
	IsPurchased bool       `db:"is_purchased"`
	PurchasedAt *time.Time `db:"purchased_at"`
	IsDelivered bool       `db:"is_delivered"`
	DeliveredAt *time.Time `db:"delivered_at"`
	Products    []QuoteProduct
}

type QuoteProduct struct {
	ID          int `db:"id"`
	QuotationID int `db:"quotation_id"`
	ProductID   int `db:"product_id"`
	Quantity    int `db:"quantity"`
}

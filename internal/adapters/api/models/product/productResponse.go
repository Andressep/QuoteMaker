package models

type ProductResponse struct {
	ID         int32   `json:"id"`
	Name       string  `json:"name"`
	CategoryID int32   `json:"category_id"`
	Length     float32 `json:"length"`
	Price      float64 `json:"price"`
	Weight     float32 `json:"weight"`
	Code       string  `json:"code"`
}

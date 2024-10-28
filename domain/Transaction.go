package domain


type Transaction struct {
	TransactionID uint `json:"transaction_id"`
	UserID uint `json:"user_id"`
	Amount uint `json:"amount"`
	TotalPrice float64 `json:"total_price"`
	Discount float64 `json:"discount"`
	ShippingCost float64 `json:"shipping_cost"`
	TotalCost float64 `json:"total_cost"`
	Items map[uint]int `json:"items"`
}
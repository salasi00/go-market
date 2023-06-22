package transactiondto

type TransactionRequest struct {
	Total     int `json:"total" form:"total"`
	UserID    int `json:"user_id" form:"user_id"`
	ProductID int `json:"product_id" form:"product_id"`
}

package models

type Transaction struct {
	ID        int     `json:"id"`
	Total     int     `json:"total"`
	UserID    int     `json:"user_id"`
	User      User    `json:"-"`
	ProductID int     `json:"product_id"`
	Product   Product `json:"product"`
}

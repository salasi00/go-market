package models

type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name" gorm:"type: varchar(255)" form:"name"`
	Desc   string `json:"desc" gorm:"type: varchar(255)" form:"desc"`
	Price  int    `json:"price" gorm:"type: int" form:"price"`
	UserID int    `json:"user_id" form:"user_id"`
	User   User   `json:"user"`
}

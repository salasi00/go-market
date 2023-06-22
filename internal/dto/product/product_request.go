package productdto

type ProductRequest struct {
	Name  string `json:"name" form:"name"`
	Desc  string `josn:"desc" form:"desc"`
	Price int    `json:"price" form:"price"`
	UserID int `json:"user_id" form:"user_id"`
}

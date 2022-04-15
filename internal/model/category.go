package model

type Category struct {
	BaseModel
	CategoryName string `json:"category_name"`
	ArticlesID   uint   `json:"articles_id"`
}

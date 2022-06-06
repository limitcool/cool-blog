package model

type Category struct {
	BaseModel
	CategoryName string `json:"category_name,omitempty"`
	ArticlesID   uint   `json:"articles_id,omitempty"`
}

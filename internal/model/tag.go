package model

type Tag struct {
	BaseModel
	TagName string `json:"tag_name"`
	Tid     uint   `json:"tid"`
}

func (t Tag) TableName() string {
	return "tags"
}

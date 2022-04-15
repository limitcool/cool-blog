package model

type Tag struct {
	BaseModel
	TagName string `json:"tag_name"`
}

//func (t Tag) TableName() string {
//	return "tags"
//}

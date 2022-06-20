package model

type Tag struct {
	BaseModel
	TagName string `json:"tag_name,omitempty"`
}

//func (t Tag) TableName() string {
//	return "tags"
//}

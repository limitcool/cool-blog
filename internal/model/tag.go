package model

type Tag struct {
	BaseModel
	TagName string `gorm:"unique" json:"tag_name"`
}

//func (t Tag) TableName() string {
//	return "tags"
//}

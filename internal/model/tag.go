package model

type Tag struct {
	TagName   string `json:"tag_name"`
	ArticleId uint
}

func (t Tag) TableName() string {
	return "tags"
}

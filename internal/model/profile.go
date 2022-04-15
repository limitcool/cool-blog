package model

type Profile struct {
	BaseModel
	Desc   string `gorm:"default:;" json:"desc"`
	Img    string `gorm:"default:;" json:"img"`
	UserId uint   `json:"user_id"`
}

func (p Profile) TableName() string {
	return "profile"
}

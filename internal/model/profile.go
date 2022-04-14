package model

type Profile struct {
	BaseModel
	Desc string `json:"desc"`
	Img  string `json:"img"`
}

func (p Profile) TableName() string {
	return "profile"
}

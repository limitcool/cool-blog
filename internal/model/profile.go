package model

type Profile struct {
	ID   int
	Desc string `json:"desc"`
	Img  string `json:"img"`
}

func (p Profile) TableName() string {
	return "profile"
}

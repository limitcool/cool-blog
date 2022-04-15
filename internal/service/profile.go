package service

type ProfileRequest struct {
	Desc string `json:"desc"`
	Img  string `json:"img"`
}

func (svc *Service) ProfileCreate(param *ProfileRequest) (uint, error) {
	id, err := svc.dao.ProfileCreate(param.Desc, param.Img)
	if err != nil {
		return 0, err
	}
	return id, err
}

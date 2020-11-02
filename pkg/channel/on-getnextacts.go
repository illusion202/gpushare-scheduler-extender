package channel

import "fmt"

type OnGetNextActs struct {
	Name string
	Func func(body *PostBody) (*OnGetNextActsResp, error)
}

func (s OnGetNextActs) Handler(body *PostBody) (*OnGetNextActsResp, error) {
	resp, err := s.Func(body)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewOnGetNextActs() *OnGetNextActs {
	return &OnGetNextActs{
		Name: "ongetnextacts",
		Func: func(body *PostBody) (*OnGetNextActsResp, error) {
			// TODO do something
			fmt.Println("====================== ongetnextacts")
			return &OnGetNextActsResp{}, nil
		},
	}
}

package channel

import "fmt"

type OnSubmit struct {
	Name string
	Func func(body *PostBody) error
}

func (s OnSubmit) Handler(body *PostBody) *OnSubmitResp {
	err := s.Func(body)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	return &OnSubmitResp{
		State:    STATE_DONE,
		Msg:      "成功",
		ErrorMsg: errMsg,
	}
}

func NewOnSubmit() *OnSubmit {
	return &OnSubmit{
		Name: "onsubmit",
		Func: func(body *PostBody) error {
			// TODO do something
			fmt.Println("====================== onsubmit")
			return nil
		},
	}
}

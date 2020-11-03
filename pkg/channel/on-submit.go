package channel

import (
	"encoding/json"
	"fmt"
	"log"
)

type OnSubmit struct {
	Name string
	Func func(body *PostBody) (*OnSubmitResp, error)
}

func (s OnSubmit) Handler(body *PostBody) *OnSubmitResp {
	resp, err := s.Func(body)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
		return &OnSubmitResp{
			ErrorMsg: errMsg,
		}
	}
	return resp
}

func NewOnSubmit() *OnSubmit {
	return &OnSubmit{
		Name: "onsubmit",
		Func: func(body *PostBody) (*OnSubmitResp, error) {
			submitJson, err := json.Marshal(body)
			if err != nil {
				log.Printf("error: onSubmit post body Marshal error: %s", err.Error())
			} else {
				log.Printf("info: onSubmit post body: %s", string(submitJson))
			}

			curActName := body.CurAct.ActName
			var ret = OnSubmitResp{}
			switch curActName {
			case ACT_BIZSREREVIEW:
				{
					ret, err = onSubmitBizSreReview(body)
				}
			case ACT_ADDQUOTA:
				{
					ret, err = onSubmitAddQuota(body)
				}
			case ACT_TRANSFERQUOTA:
				{
					ret, err = onSubmitTransferQuota(body)
				}
			default:
				{
					log.Printf("error: unexpected curActName [%s]", curActName)
					return nil, fmt.Errorf("error: unexpected curActName [%s]", curActName)
				}
			}
			if err != nil {
				log.Printf("error: onSubmit current_act_name[%s], error: %s", curActName, err.Error())
			}
			return &ret, err
		},
	}
}

// 业务线SRE审批
func onSubmitBizSreReview(body *PostBody) (resp OnSubmitResp, err error) {
	log.Println("debug: onSubmitBizSreReview")

	return resp, err
}

// 配额新增
func onSubmitAddQuota(body *PostBody) (resp OnSubmitResp, err error) {
	log.Println("debug: onSubmitAddQuota")

	return resp, err
}

// 配额转移
func onSubmitTransferQuota(body *PostBody) (resp OnSubmitResp, err error) {
	log.Println("debug: onSubmitTransferQuota")

	return resp, err
}

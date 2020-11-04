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
					// do nothing
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
	// 下一步
	nextActName := body.NextAct.ActName
	switch nextActName {
	case ACT_ADDQUOTA:
		{
			// 配额新增，调用配额新增接口（容器云配额申请供KRP调用接口：杨挺 https://docs.corp.kuaishou.com/d/home/fcADS92Swm2-KBX3JrnQsOEJD#）
			// 设置state为DOING，失败设置为FAIL
		}
	case ACT_TRANSFERQUOTA:
		{
			// 设置state为DONE
			resp = OnSubmitResp{
				State: STATE_DONE,
			}
		}
	default:
		log.Printf("error: onSubmitBizSreReview unexpected nextActName [%s]", nextActName)
		return resp, fmt.Errorf("error: onSubmitBizSreReview unexpected nextActName [%s]", nextActName)
	}

	return resp, err
}

// 配额转移
func onSubmitTransferQuota(body *PostBody) (resp OnSubmitResp, err error) {
	log.Println("debug: onSubmitTransferQuota")

	// 调用容器云配额迁移接口（https://wiki.corp.kuaishou.com/pages/viewpage.action?pageId=450791699 /api/v2/resource-quota/transfer）
	// 设置state为DONE，失败设置为FAIL
	return resp, err
}

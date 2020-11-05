package channel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
			client := &http.Client{}

			failRep := OnSubmitResp{
				State: STATE_FAIL,
			}

			successRep := OnSubmitResp{
				State: STATE_DOING,
			}

			// TODO 确认配额新增请求body参数怎么填充
			reqBody := QuotaAddReq{
				OrderId:  40622,
				PageNum:  1,
				PageSize: 10,
			}

			reqJson, err := json.Marshal(reqBody)
			if err != nil {
				log.Printf("error: onSubmitBizSreReview request body Marshal error: %s", err.Error())
				return failRep, err
			}

			log.Printf("debug: onSubmitBizSreReview apply quota request body: %s", string(reqJson))

			req, err := http.NewRequest("POST", ApplyQuotaURL+QUOTA_APPLY_URI,
				strings.NewReader(string(reqJson)))
			if err != nil {
				return failRep, err
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Token "+ApplyQuotaToken)
			httpResp, err := client.Do(req)

			defer httpResp.Body.Close()

			resBody, err := ioutil.ReadAll(httpResp.Body)
			if err != nil {
				log.Printf("error: onSubmitBizSreReview response body read error: %s", err.Error())
				return resp, err
			}
			result := &QuotaAddResp{}
			json.Unmarshal(resBody, result) // 解析json字符串
			log.Printf("debug: onSubmitBizSreReview apply quota response body: %s", string(resBody))

			if !result.Success {
				return failRep, fmt.Errorf(result.ErrorMsg)
			}
			return successRep, nil
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

	// 调用容器云配额迁移接口（https://wiki.corp.kuaishou.com/pages/viewpage.action?pageId=450791699
	// /api/v2/resource-quota/transfer）
	// 设置state为DONE，失败设置为FAIL
	client := &http.Client{}

	failRep := OnSubmitResp{
		State: STATE_FAIL,
	}

	successRep := OnSubmitResp{
		State: STATE_DONE,
	}

	reqBody := QuotaTransferReq{
		Approved:     true,
		UserId:       "sunxiaofei",
		NodeId:       1,
		TNodeId:      1,
		TUserId:      "haoshiying",
		Region:       "HB1",
		AZ:           "YZ",
		Group:        "all",
		ResourcePool: "默认资源池",
		Description:  "xxx",
		Quota: &Quota{
			Cpu:    10,
			Memory: 10,
			Gpu:    1,
		},
	}

	reqJson, err := json.Marshal(reqBody)
	if err != nil {
		log.Printf("error: onSubmitTransferQuota request body Marshal error: %s", err.Error())
		return failRep, err
	}

	log.Printf("debug: onSubmitTransferQuota transfer quota request body: %s", string(reqJson))

	req, err := http.NewRequest("POST", TransferQuotaURL+QUOTA_TRANSFER_URI,
		strings.NewReader(string(reqJson)))
	if err != nil {
		return failRep, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", TransferQuotaToken)
	httpResp, err := client.Do(req)

	defer httpResp.Body.Close()

	resBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		log.Printf("error: onSubmitTransferQuota response body read error: %s", err.Error())
		return resp, err
	}
	result := &QuotaTransferResp{}
	json.Unmarshal(resBody, result) // 解析json字符串
	log.Printf("debug: onSubmitTransferQuota transfer quota response body: %s", string(resBody))

	if result.StatusCode != 0 {
		return failRep, fmt.Errorf(result.Message)
	}
	return successRep, nil
}

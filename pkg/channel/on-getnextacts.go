package channel

import (
	"encoding/json"
	"fmt"
	"log"
)

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
			getNextActsJson, err := json.Marshal(body)
			if err != nil {
				log.Printf("error: onGetNextAct post body Marshal error: %s", err.Error())
			} else {
				log.Printf("info: onGetNextAct post body: %s", string(getNextActsJson))
			}
			curActName := body.CurAct.ActName
			var ret = OnGetNextActsResp{}
			switch curActName {
			case ACT_BIZSREREVIEW:
				{
					ret, err = onGetNextActsBizSreReview(body)
				}
			case ACT_ADDQUOTA:
				{
					ret, err = onGetNextActsAddQuota(body)
				}
			case ACT_TRANSFERQUOTA:
				{
					ret, err = onGetNextActsTransferQuota(body)
				}
			default:
				{
					log.Printf("error: unexpected curActName [%s]", curActName)
					return nil, fmt.Errorf("error: unexpected curActName [%s]", curActName)
				}
			}
			if err != nil {
				log.Printf("error: onGetNextAct current_act_name[%s], error: %s", curActName, err.Error())
			}
			return &ret, err
		},
	}
}

// 业务线SRE审批，需要做分支判断
// 1、配额满足：TransferQuota（配额转移）
// 2、配额不满足：AddQuota（配额新增）
// 根据entityData里的属性"operation"判断，取属性值：
// 1、quotaMigration：配额转移
// 2、quotaNew，quotaUrgence：配额新增
func onGetNextActsBizSreReview(body *PostBody) (resp OnGetNextActsResp, err error) {
	log.Println("debug: onGetNextActsBizSreReview")

	if body.EntityData == nil || len(body.EntityData) <= 0 {
		return resp, fmt.Errorf("error: onGetNextActsBizSreReview entityData is nil or empty")
	}

	operation := body.EntityData["operation"]
	var nextAct string
	switch operation {
	case QUOTANEW:
		{
			nextAct = ACT_ADDQUOTA
		}
	case QUOTAURGENCE:
		{
			nextAct = ACT_ADDQUOTA
		}
	case QUOTAMIGRATION:
		{
			nextAct = ACT_TRANSFERQUOTA
		}
	default:
		{
			return resp, fmt.Errorf("error: onGetNextActsBizSreReview unexpected operation type [%s]", operation)
		}
	}

	// validate if NextCandidateActs contains required nextAct or not.
	var containsAct bool
	for _, act := range body.NextCandidateActs {
		if act.ActName == nextAct {
			containsAct = true
			resp.NextCandidateActs = append(resp.NextCandidateActs, nextAct)
			break
		}
	}

	if !containsAct {
		return resp, fmt.Errorf("error: onGetNextActsBizSreReview, NextCandidateActs doesn`t contains act [%s]", nextAct)
	}

	return resp, err
}

// 配额新增，不做分支判断，直接返回nextCandidateActs名称
// 一般默认为配额转移：TransferQuota
func onGetNextActsAddQuota(body *PostBody) (resp OnGetNextActsResp, err error) {
	log.Println("debug: onGetNextActsAddQuota")
	for _, act := range body.NextCandidateActs {
		resp.NextCandidateActs = append(resp.NextCandidateActs, act.ActName)
	}

	return resp, nil
}

// 配额转移，不做分支判断，直接返回nextCandidateActs名称
// 一般默认为完结：End
func onGetNextActsTransferQuota(body *PostBody) (resp OnGetNextActsResp, err error) {
	log.Println("debug: onGetNextActsTransferQuota")
	for _, act := range body.NextCandidateActs {
		resp.NextCandidateActs = append(resp.NextCandidateActs, act.ActName)
	}

	return resp, nil
}

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
			actName := body.CurAct.ActName
			var ret = OnGetNextActsResp{}
			switch actName {
			case ACT_BIZSREREVIEW:
				{
					ret, err = onGetNextActsBizSreReview(body)
				}
			case ACT_ADDQUOTA: // do nothing
			case ACT_TRANSFERQUOTA: // do nothing
			default:
				{
					log.Printf("error: unexpected actName [%s]", actName)
					return nil, fmt.Errorf("error: unexpected actName [%s]", actName)
				}
			}
			if err != nil {
				log.Printf("error: onGetNextAct current_act_name[%s], error: %s", actName, err.Error())
			}
			return &ret, err
		},
	}
}

// 业务线SRE审批
func onGetNextActsBizSreReview(body *PostBody) (resp OnGetNextActsResp, err error) {
	log.Println("debug: onGetNextActsBizSreReview")

	return resp, err
}

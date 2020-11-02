package channel

type PostBody struct {
	Handler           string                 `json:"handler"`
	Buttons           []interface{}          `json:"buttons"`
	ActMode           string                 `json:"actMode"`
	CurAct            *Act                   `json:"curAct"`
	NextAct           *Act                   `json:"nextAct"`
	NextCandidateActs []*Act                 `json:"nextCandidateActs"`
	Entity            *Entity                `json:"entity"`
	EntityData        map[string]interface{} `json:"entityData"`
}

type Act struct {
	EditFlag      string   `json:"editFlag"`
	ActDesc       string   `json:"actDesc"`
	DelFlag       int64    `json:"delFlag"`
	ActName       string   `json:"actName"`
	FlowVer       int64    `json:"flowVer"`
	DenyEndBtn    int64    `json:"denyEndBtn"`
	OpTime        int64    `json:"opTime"`
	ActTitle      string   `json:"actTitle"`
	ActType       string   `json:"actType"`
	ActRule       *ActRule `json:"actRule"`
	Id            int64    `json:"id"`
	SubEntityType int64    `json:"subEntityType"`
	FlowId        int64    `json:"flowId"`
}

type ActRule struct {
	HandleType string `json:"handleType"`
	BaseOn     string `json:"baseOn"`
	ActId      int64  `json:"actId"`
	Id         int64  `json:"id"`
	PreActId   int64  `json:"preActId"`
	ActName    string `json:"actName"`
	FormField  string `json:"formField"`
}

type Entity struct {
	SerialNum         string `json:"serialNum"`
	EntityId          string `json:"entityId"`
	PlatformId        int64  `json:"platformId"`
	DelFlag           int64  `json:"delFlag"`
	SubEntityTotalNum int64  `json:"subEntityTotalNum"`
	FlowName          string `json:"flowName"`
	FlowVer           int64  `json:"flowVer"`
	EntityTitle       string `json:"entityTitle"`
	FlowTitle         string `json:"flowTitle"`
	CreateTime        int64  `json:"createTime"`
	SubEntityOverNum  int64  `json:"subEntityOverNum"`
	DraftUid          string `json:"draftUid"`
	Id                int64  `json:"id"`
	State             string `json:"state"`
	FlowId            int64  `json:"flowId"`
	DraftName         string `json:"draftName"`
}

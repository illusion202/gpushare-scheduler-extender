package channel

// 容器云配额申请Request
// 示例：
// curl -H 'Authorization: Token 4d75a6f1d7ee711039bff4b842c7d0f5' \
// -H 'Content-Type: application/json' \
// -d '{"orderId":40622,"pageNum":1,"pageSize":10}' \
// -X POST https://bp-beta.corp.kuaishou.com/api/scm/purchaseforecast/purchase-forecast-list/v1/forecast/list
type QuotaAddReq struct {
	OrderId  int64 `json:"orderId"`
	PageNum  int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
}

// 容器云配额申请Response
type QuotaAddResp struct {
	Status      int64           `json:"status"`
	ErrorMsg    string          `json:"error_msg"`
	Total       int64           `json:"total"`
	Data        []*QuotaAddData `json:"data"`
	Success     bool            `json:"success"`
	NotSuccess  bool            `json:"notSuccess"`
	SystemError bool            `json:"systemError"`
}

type QuotaAddData struct {
	Id                     int64             `json:"id"`
	PurchaseForecastListId int64             `json:"purchaseForecastListId"`
	ProcessId              string            `json:"processId"`
	Type                   int64             `json:"type"`
	Creator                string            `json:"creator"`
	ForecastType           int64             `json:"forecastType"`
	LStatus                int64             `json:"lstatus"`
	Status                 int64             `json:"status"`
	StatusName             string            `json:"statusName"`
	DeliveryStatus         string            `json:"deliveryStatus"`
	DemandType             string            `json:"demandType"`
	NodeL1                 int64             `json:"nodeL1"`
	NodeL1Name             string            `json:"nodeL1Name"`
	NodeL2                 int64             `json:"nodeL2"`
	NodeL2Name             string            `json:"nodeL2Name"`
	SuitId                 int64             `json:"suitId"`
	SuitName               string            `json:"suitName"`
	KWaiTypeId             int64             `json:"kwaiTypeId"`
	KWaiType               string            `json:"kwaiType"`
	Amount                 int64             `json:"amount"`
	IdcId                  int64             `json:"idcId"`
	Idc                    string            `json:"idc"`
	AzId                   int64             `json:"azId"`
	AZ                     string            `json:"az"`
	KeyDriver              string            `json:"keydriver"`
	Formula                string            `json:"formula"`
	ForecastCycleId        int64             `json:"forecastCycleId"`
	ForeCastCycleName      string            `json:"foreCastCycleName"`
	CycleStartTime         string            `json:"cycleStartTime"`
	CycleEndTime           string            `json:"cycleEndTime"`
	DemandTime             string            `json:"demandTime"`
	CreateTime             string            `json:"createTime"`
	UpdateTime             string            `json:"updateTime"`
	CurUser                string            `json:"curUser"`
	ActTitle               string            `json:"actTitle"`
	Servers                []*QuotaAddServer `json:"servers"`
}

type QuotaAddServer struct {
	Id                       int64  `json:"id"`
	Status                   int64  `json:"status"`
	PurchaseForecastDetailId int64  `json:"purchaseForecastDetailId"`
	ServerId                 int64  `json:"serverId"`
	NodeId                   int64  `json:"nodeId"`
	NodePath                 string `json:"nodePath"`
	PcsId                    string `json:"pcsId"`
	Hostname                 string `json:"hostname"`
	CustomName               string `json:"customName"`
	SN                       string `json:"sn"`
	Creator                  string `json:"creator"`
	Remark                   string `json:"remark"`
	CreateTime               string `json:"createTime"`
	RenameKey                string `json:"renameKey"`
}

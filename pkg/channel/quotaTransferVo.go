package channel

// 容器云配额转移Request, URI: /api/v2/resource-quota/transfer
type QuotaTransferReq struct {
	Approved     bool   `json:"approved"`
	UserId       string `json:"user_id"`
	NodeId       int64  `json:"node_id"`
	TNodeId      int64  `json:"tnode_id"`
	TUserId      string `json:"tuser_id"`
	Region       string `json:"region"`
	AZ           string `json:"az"`
	Group        string `json:"group"`
	ResourcePool string `json:"resource_pool"`
	Description  string `json:"description"`
	Quota        *Quota `json:"quota"`
}

type Quota struct {
	Cpu    int64 `json:"cpu"`
	Memory int64 `json:"memory"`
	Gpu    int64 `json:"gpu"`
}

// 容器云配额转移Response
type QuotaTransferResp struct {
	StatusCode int64  `json:"status_code"`
	Message    string `json:"message"`
}

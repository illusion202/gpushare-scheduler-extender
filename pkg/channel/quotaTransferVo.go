package channel

// 容器云配额转移Request
// 示例：
// curl -H 'Authorization: 7d5c3f66-5986-422d-9d06-fbd951b90c6e' \
// -H 'Content-Type: application/json' \
// -d '{"approved": true,"user_id": "sunxiaofei","node_id":1,"tnode_id":1,"tuser_id":"haoshiying","region":"HB1","az":"YZ","group":"all","resource_pool":"默认资源池","description":"xxx","quota":{"cpu":10,"memory":10,"gpu":1}}' \
// -X POST https://kcs-test.corp.kuaishou.com/api/v2/resource-quota/transfer
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

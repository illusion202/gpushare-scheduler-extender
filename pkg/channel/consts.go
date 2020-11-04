package channel

const (
	// OnSubmit 动作状态
	STATE_DONE  = "DONE"  // 动作完成
	STATE_FAIL  = "FAIL"  // 接口调用失败
	STATE_DOING = "DOING" // 接口调用成功
	STATE_TODO  = "TODO"  // 暂不涉及

	// act_name
	ACT_APPLY         = "Apply"         // 拟稿
	ACT_BIZSREREVIEW  = "BizSreReview"  // 业务SRE审批
	ACT_ADDQUOTA      = "AddQuota"      // 配额新增
	ACT_TRANSFERQUOTA = "TransferQuota" // 配额转移
	ACT_END           = "End"           // 流程结束

	// 星环透传给流程平台的操作参数
	QUOTAMIGRATION = "quotaMigration" // 配额转移
	QUOTANEW       = "quotaNew"       // 配额新增
	QUOTAURGENCE   = "quotaUrgence"   // 配额紧急新增

	QUOTA_APPLY_URI    = "/api/scm/purchaseforecast/purchase-forecast-list/v1/forecast/list" // 配额申请URI
	QUOTA_TRANSFER_URI = "/api/v2/resource-quota/transfer"                                   // 配额转移URI
)

var URL string
var Token string

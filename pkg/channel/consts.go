package channel

const (
	// OnSubmit 动作状态
	STATE_DONE  = "DONE"
	STATE_FAIL  = "FAIL"
	STATE_DOING = "DOING"
	STATE_TODO  = "TODO"

	// act_name
	ACT_APPLY         = "Apply"         // 拟稿
	ACT_BIZSREREVIEW  = "BizSreReview"  // 业务SRE审批
	ACT_ADDQUOTA      = "AddQuota"      // 配额新增
	ACT_TRANSFERQUOTA = "TransferQuota" // 配额转移
	ACT_END           = "End"           // 流程结束
)

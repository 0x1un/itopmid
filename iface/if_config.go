package iface

type Configer interface {
	ItopUrl() string       // itop rest api url
	ItopUsername() string  // itop username
	DingAppkey() string    // dingtalk appkey
	DingAppsecret() string // dingtalk appsecret
	DingApprovID() string  // 钉钉审批单模板列表
	DingAgentID() string   // 钉钉应用的agentd
	DingUserID() string    // 钉钉用户的uid
	PostgresUrl() string   // postgres数据库链接
}

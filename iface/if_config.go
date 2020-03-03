package iface

type Configer interface {
	GetItopUrl() string       // itop rest api url
	GetItopUsername() string  // itop username
	GetDingAppkey() string    // dingtalk appkey
	GetDingAppsecret() string // dingtalk appsecret
	GetDingApprovID() string  // 钉钉审批单模板列表
	GetDingAgentID() string   // 钉钉应用的agentd
	GetDingUserID() string    // 钉钉用户的uid
	GetPostgresUrl() string   // postgres数据库链接
}

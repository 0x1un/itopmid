package iface

type Configer interface {
	GetDingDeptID() int64
	GetDingAgentID() int64
	GetItopUrl() string
	GetItopUsername() string
	GetItopPassword() string
	GetDingAppkey() string
	GetDingAppsecret() string
	GetDingApprovID() string
	GetDingUserID() string
	GetDatabaseURL() string
	ReadConfigFile(filename string)
}

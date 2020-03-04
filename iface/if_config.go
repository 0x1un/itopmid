package iface

type Configer interface {
	ReadConfigFile(filename string)
	GetItopUrl() string
	GetItopUsername() string
	GetItopPassword() string
	GetDingAppkey() string
	GetDingAppsecret() string
	GetDingApprovID() string
	GetDingAgentID() string
	GetDingUserID() string
	GetDatabaseURL() string
}

package support

type ItopMidConfig struct {
	ItopUrl       string `json:"itopurl"`
	ItopUsername  string `json:"itopusername"`
	DingAppkey    string `json:"dingappkey"`
	DingAppsecret string `json:"dingappsecret"`
	DingApprovID  string `json:"dingapprovid"`
	DingAgentID   string `json:"dingagentid"`
	DingUserID    string `json:"dinguserid"`
	PostgresUrl   string `json:"postgresurl"`
}

func (self *ItopMidConfig) Init() {}

func (self *ItopMidConfig) ReadConfigFile(filename string) {

}

func (self *ItopMidConfig) GetItopUrl() string {
	return self.ItopUrl
}

// func (self *ItopMidConfig)
// func (self *ItopMidConfig)
// func (self *ItopMidConfig)
// func (self *ItopMidConfig)

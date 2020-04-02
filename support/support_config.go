package support

import (
	"io/ioutil"

	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/util"
	json "github.com/json-iterator/go"
)

type ItopMidConfig struct {
	DingAgentID    int64  `json:"dingagentid"`
	DingDeptID     int64  `json:"dingdeptid"`
	ItopUrl        string `json:"itopurl"`
	ItopUsername   string `json:"itopusername"`
	ItopPasswd     string `json:"itoppasswd"`
	DingAppkey     string `json:"dingappkey"`
	DingAppsecret  string `json:"dingappsecret"`
	DingApprovID   string `json:"dingapprovid"`
	DingUserID     string `json:"dinguserid"`
	PostgresAddr   string `json:"postgresaddr"`
	PostgresPort   string `json:"postgresport"`
	PostgresUname  string `json:"postgresuname"`
	PostgresPasswd string `json:"postgrespasswd"`
	PostgresDb     string `json:"postgresdb"`
	Debug          bool   `json:"debug"`
}

func (self *ItopMidConfig) GetItopUrl() string {
	return self.ItopUrl
}
func (self *ItopMidConfig) GetItopUsername() string {
	return self.ItopUsername
}
func (self *ItopMidConfig) GetItopPassword() string {
	return self.ItopPasswd
}
func (self *ItopMidConfig) GetDingAppkey() string {
	return self.DingAppkey
}
func (self *ItopMidConfig) GetDingAppsecret() string {
	return self.DingAppsecret
}
func (self *ItopMidConfig) GetDingApprovID() string {
	return self.DingApprovID
}
func (self *ItopMidConfig) GetDingAgentID() int64 {
	return self.DingAgentID
}
func (self *ItopMidConfig) GetDingUserID() string {
	return self.DingUserID
}
func (self *ItopMidConfig) GetDatabaseURL() string {
	return util.GenDBUrl(self.PostgresAddr, self.PostgresPort, self.PostgresDb, self.PostgresUname, self.PostgresPasswd)
}

func (self *ItopMidConfig) GetDingDeptID() int64 {
	return self.DingDeptID
}

func (self *ItopMidConfig) ReadConfigFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		iface.LOGGER.Error("Failed open file:", err.Error())
	}
	if err := json.Unmarshal(content, self); err != nil {
		iface.LOGGER.Error("Failed to unmarshal json file:", err.Error())
	}
}

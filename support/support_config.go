package support

import (
	"fmt"
	"io/ioutil"

	"github.com/0x1un/itopmid/iface"
	json "github.com/json-iterator/go"
)

type ItopMidConfig struct {
	ItopUrl       string `json:"itopurl"`
	ItopUsername  string `json:"itopusername"`
	ItopPasswd    string `json:"itoppasswd"`
	DingAppkey    string `json:"dingappkey"`
	DingAppsecret string `json:"dingappsecret"`
	DingApprovID  string `json:"dingapprovid"`
	DingAgentID   string `json:"dingagentid"`
	DingUserID    string `json:"dinguserid"`
	PostgresUrl   string `json:"postgresurl"`
}

func Test() {
	i := &ItopMidConfig{}
	i.ReadConfigFile("./itopmid.json")
	fmt.Println(*i)
}

func (self *ItopMidConfig) Init() {}

func (self *ItopMidConfig) ReadConfigFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		iface.LOGGER.Error("Failed open file:", err.Error())
	}
	if err := json.Unmarshal(content, self); err != nil {
		iface.LOGGER.Error("Failed to unmarshal json file:", err.Error())
	}
}

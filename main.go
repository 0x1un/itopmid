package main

import (
	"github.com/0x1un/itopmid/core"
	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

func init() {
	// load config
	config := &support.ItopMidConfig{}
	config.ReadConfigFile("itopmid.json")
	iface.CONFIG = config
	iface.ITOP_USERNAME = iface.CONFIG.GetItopUsername()
	iface.ITOP_PASSWORD = iface.CONFIG.GetItopPassword()

	// load request
	request := &support.RequestData{}
	iface.REQUEST = request

	// load logger
	logger := &support.ItopMidLogger{}
	iface.LOGGER = logger

	// init database
	context := &support.ItopMidContext{}
	context.OpenDB("postgres")
	iface.CONTEXT = context

	// init retry queuer
	queue := &support.RetryQueue{}
	iface.RETRY_QUEUE = queue

	// init dingtalk client
	// client := api.NewClient(iface.CONFIG.GetDingAppkey(), iface.CONFIG.GetDingAppsecret())
	// client.ProcessReq.DeptId = iface.CONFIG.GetDingDeptID()
	// client.ProcessReq.AgentId = iface.CONFIG.GetDingAgentID()
	// client.ProcessReq.OriginatorUserId = iface.CONFIG.GetDingUserID()
	// client.ProcessReq.ProcessCode = iface.CONFIG.GetDingApprovID()
	// iface.CLIENT = client
}

func main() {
	core.UpdateItopTicket("R-000006")

	// defer iface.CONTEXT.CloseDB()

	// core.FetcheFromITOP(iface.CONFIG.GetItopUrl(), iface.REQUEST.GenUserRequest())
	// if iface.RETRY_QUEUE.Len() > 0 {
	// 	fmt.Println(iface.RETRY_QUEUE)
	// }
	// code := core.GetProcessStatusByID("629e713f-5973-469a-a1f5-359b0df22a7c")
	// fmt.Println(code)

}

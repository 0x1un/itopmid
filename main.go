package main

import (
	"fmt"

	"github.com/0x1un/boxes/dingtalk/api"
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
	client := api.NewClient(iface.CONFIG.GetDingAppkey(), iface.CONFIG.GetDingAppsecret())
	client.ProcessReq.DeptId = iface.CONFIG.GetDingDeptID()
	client.ProcessReq.AgentId = iface.CONFIG.GetDingAgentID()
	client.ProcessReq.OriginatorUserId = iface.CONFIG.GetDingUserID()
	client.ProcessReq.ProcessCode = iface.CONFIG.GetDingApprovID()
	iface.CLIENT = client
}

func main() {
	defer iface.CONTEXT.CloseDB()

	core.FetchItopTicketAndSendToDingtalk(iface.CONFIG.GetItopUrl(), iface.REQUEST.GenUserRequest())
	if iface.RETRY_QUEUE.Len() > 0 {
		fmt.Println(iface.RETRY_QUEUE)
	}

	// ticker := time.NewTicker(time.Duration(time.Second * 5))
	// defer ticker.Stop()
	// code := make(chan int)
	// go func() {
	// 	for range ticker.C {
	// 		code <- core.GetProcessStatusByID("89131548-abed-47ef-925f-f97f35d6a9a5")
	// 	}
	// }()
	// for c := range code {
	// 	switch c {
	// 	case core.PROCESS_IS_NEW:
	// 		fmt.Println("Process is new")
	// 	case core.PROCESS_IS_RUNNING:
	// 		fmt.Println("Process is running...")
	// 	case core.PROCESS_IS_COMPLETED:
	// 		fmt.Println("Process is completed")
	// 		return
	// 	case core.PROCESS_IS_TERMINATED:
	// 		fmt.Println("Process is terminated")
	// 		return
	// 	default:
	// 		fmt.Println("未知错误..")
	// 	}
	// }
}

package main

import (
	_ "net/http/pprof"

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

	// load logger
	logger := &support.ItopMidLogger{}
	iface.LOGGER = logger

	// init database
	context := &support.ItopMidContext{}
	context.OpenDB("postgres")
	iface.CONTEXT = context
}

func main() {

	request_data, err := core.NewRestAPIAuthData(iface.CONFIG.GetItopUsername(), iface.CONFIG.GetItopPassword())
	if err != nil {
		panic(err)
	}

	// 从itop中获取所有状态为开启的工单
	resp := core.FetcheFromITOP(iface.CONFIG.GetItopUrl(), request_data)
	for _, v := range resp.Object {
		if err := core.StoreTicketFromITOP(iface.CONTEXT.GetDB(), v.Filed); err != nil {
			iface.LOGGER.Error(err.Error())
		}
	}

	client := api.NewClient(iface.CONFIG.GetDingAppkey(), iface.CONFIG.GetDingAppsecret())
	client.ProcessReq.DeptId = iface.CONFIG.GetDingDeptID()
	client.ProcessReq.AgentId = iface.CONFIG.GetDingAgentID()
	client.ProcessReq.OriginatorUserId = iface.CONFIG.GetDingUserID()
	client.ProcessReq.ProcessCode = iface.CONFIG.GetDingApprovID()
	// 发送来自itop的工单至钉钉工单中
	core.SendToProv(client, resp)
}

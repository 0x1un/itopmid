package app

import (
	"time"

	"github.com/0x1un/boxes/dingtalk/api"
	"github.com/0x1un/itopmid/core"
	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

var APPLICATION Application

const (
	DURATION = time.Duration(5 * time.Second)
)

type Application interface {
	Start()
}

type ItopMidApp struct {
	Duration time.Duration
}

func Init() {
	APPLICATION = &ItopMidApp{
		Duration: time.Duration(time.Second * 5),
	}
}

func (self *ItopMidApp) init() {

	// load config
	config := &support.ItopMidConfig{}
	config.ReadConfigFile("conf/conf.json")
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

	// init retry queue & ticket queue
	rqueue := &support.Queue{}
	tqueue := &support.TicketQueue{}
	squeue := &support.TicketQueue{}
	iface.RETRY_QUEUE = rqueue
	iface.TICKET_QUEUE = tqueue
	iface.STATUS_QUEUE = squeue

	// init dingtalk client
	client := api.NewClient(iface.CONFIG.GetDingAppkey(), iface.CONFIG.GetDingAppsecret())
	client.ProcessReq.DeptId = iface.CONFIG.GetDingDeptID()
	client.ProcessReq.AgentId = iface.CONFIG.GetDingAgentID()
	client.ProcessReq.OriginatorUserId = iface.CONFIG.GetDingUserID()
	client.ProcessReq.ProcessCode = iface.CONFIG.GetDingApprovID()
	iface.CLIENT = client

}

func (self *ItopMidApp) Start() {
	self.init()
	defer iface.CONTEXT.CloseDB()
	ticker := time.NewTicker(self.Duration)
	defer ticker.Stop()

	// done := make(chan time.Time)
	for range ticker.C {
		core.FetchItopTicketAndSendToDingtalk(iface.CONFIG.GetItopUrl(), iface.REQUEST.GenUserRequest())
		for k, v := range iface.TICKET_QUEUE.Self() {
			go checkDingTicket(k, v)
		}
	}
}

func checkDingTicket(ref, id string) {
	c := core.GetProcessStatusByID(id)
	status := iface.STATUS_QUEUE.Get(id)
	switch c {
	case core.PROCESS_IS_NEW:
		iface.LOGGER.Info("%s: => new", ref)
	case core.PROCESS_IS_RUNNING:
		iface.LOGGER.Info("%s: => running", ref)
	case core.PROCESS_IS_COMPLETED:
		iface.LOGGER.Info("%s: => completed", ref)
		core.UpdateItopTicket(ref, status, "resolved")
		iface.TICKET_QUEUE.Del(ref)
	case core.PROCESS_IS_TERMINATED:
		iface.LOGGER.Info("%s: => terminated", ref)
		core.UpdateItopTicket(ref, status, "rejected")
		iface.TICKET_QUEUE.Del(ref)
	default:
		iface.LOGGER.Info("%s: => unkown", ref)
	}
}

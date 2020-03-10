package iface

import "github.com/0x1un/boxes/dingtalk/api"

var LOGGER Logger // 日志

var CONFIG Configer // 配置

var CONTEXT ItopMidContext // 上下文，包含数据库，缓存之类的东西

var ITOP_USERNAME string // ITOP username

var ITOP_PASSWORD string // ITOP password

var REQUEST Requester // ITOP 请求的数据，包含UserRequest, Contact, Person 之类的请求数据与响应数据

var CLIENT *api.DingTalkClient

var RETRY_QUEUE TicketRetryQueuer

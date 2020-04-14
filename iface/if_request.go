package iface

import (
	"strings"
)

type Requester interface {
	GenUserRequest() *strings.Reader
	GenPersonRequest(friendlyname string) *strings.Reader
}

type Queuer interface {
	Len() int
	Push(content interface{})
	Tail() interface{}
	Self() interface{}
	Pop() interface{}
}

type TicketQueuer interface {
	Set(k, v string)
	Del(k string)
	Self() map[string]string
	Get(key string) string
}

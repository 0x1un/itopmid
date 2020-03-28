package iface

import (
	"strings"
)

type Requester interface {
	GenUserRequest() *strings.Reader
	GenPersonRequest(friendlyname string) *strings.Reader
}

type Queuer interface {
	Pop() interface{}
	Push(content interface{})
	Tail() interface{}
	Self() interface{}
	Len() int
}

type TicketQueuer interface {
	Set(k, v string)
	Del(k string)
	Self() map[string]string
}

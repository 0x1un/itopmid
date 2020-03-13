package iface

import (
	"strings"
)

type Requester interface {
	GenUserRequest() *strings.Reader
	GenPersonRequest(friendlyname string) *strings.Reader
}

type TicketRetryQueuer interface {
	Pop() interface{}
	Push(content interface{})
	Len() int
}

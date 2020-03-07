package iface

import "strings"

type Requester interface {
	GenUserRequest() *strings.Reader
	GenContactRequest() *strings.Reader
}

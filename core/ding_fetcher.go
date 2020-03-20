package core

import (
	"github.com/0x1un/itopmid/iface"
)

const (
	PROCESS_IS_NEW        = 0
	PROCESS_IS_RUNNING    = 1
	PROCESS_IS_TERMINATED = 2
	PROCESS_IS_COMPLETED  = 3
	PROCESS_UNKOWN_ERROR  = 4
)

func GetProcessStatusByID(id string) int {
	res, err := iface.CLIENT.GetProcessInstanceDetail(id)
	if err != nil {
		iface.LOGGER.Error(err.Error())
		return PROCESS_UNKOWN_ERROR
	}
	if res.Errcode != 0 {
		return PROCESS_UNKOWN_ERROR
	}
	switch res.ProcessInstanc.Status {
	case "NEW":
		return PROCESS_IS_NEW
	case "RUNNING":
		return PROCESS_IS_RUNNING
	case "TERMINATED":
		return PROCESS_IS_TERMINATED
	case "COMPLETED":
		return PROCESS_IS_COMPLETED
	}
	return PROCESS_UNKOWN_ERROR
}

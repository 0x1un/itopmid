package core

import (
	"fmt"

	"github.com/0x1un/itopmid/iface"
)

const (
	PROCESS_IS_NEW        = 0
	PROCESS_IS_RUNNING    = 1
	PROCESS_IS_TERMINATED = 2
	PROCESS_IS_COMPLETED  = 3
)

func GetProcessStatusByID(id string) int {
	res, err := iface.CLIENT.GetProcessInstanceDetail(id)
	if err != nil {
		iface.LOGGER.Error(err.Error())
		return PROCESS_IS_TERMINATED
	}
	fmt.Println(res)
	return 0
}

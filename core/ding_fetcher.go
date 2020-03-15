package core

const (
	PROCESS_IS_NEW        = 0
	PROCESS_IS_RUNNING    = 1
	PROCESS_IS_TERMINATED = 2
	PROCESS_IS_COMPLETED  = 3
)

func GetProcessStatusByID(id string) int {
}

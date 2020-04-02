package core

import (
	"encoding/json"
	"net/http"

	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
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

// ref is itop ticket number;
// k is query conditaion (status: resolved/new);
// v is change value (status=new/pending/rejected/resolved...)
func UpdateItopTicket(ref, k, v string) {
	key := &support.UpdateKey{
		Status: k,
		Ref:    ref,
	}
	fields := &support.UpdateFields{
		Status: v,
	}
	// ud => update request data
	ud := support.NewRequestUpdateData("UserRequest", key, fields)
	rsp, err := Request(http.MethodPost, iface.CONFIG.GetItopUrl(), ud)
	if err != nil {
		iface.LOGGER.Error(err.Error())
		return
	}
	t := &support.UserReqResponse{}
	if err := json.Unmarshal(rsp, t); err != nil {
		iface.LOGGER.Error(err.Error())
		return
	}
	if t.Code != 0 {
		iface.LOGGER.Error(t.Message)
		return
	}
	if err := updateField(ref, "resolved", true); err != nil {
		iface.LOGGER.Error(err.Error())
		return
	}
	iface.LOGGER.Info("%s is updated", ref)
}

func updateField(ref, target string, value interface{}) error {
	h := iface.CONTEXT.GetDB().Begin()
	err := h.Table("itop_ticket").Where("ref=?", ref).Update(target, value).Error
	if err != nil {
		h.Rollback()
		return err
	}
	h.Commit()
	return nil
}

package core

import (
	"fmt"

	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

func SendSingleTicketToDingtalkProcess(content *support.ResponseContent) error {
	ref := content.Filed.Ref
	form := ConvertSingleUserRequest(content)
	resp, err := iface.CLIENT.SendProcess(*form)
	if err != nil {
		return err
	}
	if resp.ErrCode != 0 {
		return fmt.Errorf("%s", resp.ErrMsg)
	}
	if err := setItopTicketFlag(ref); err != nil {
		return err
	}
	iface.LOGGER.Info("Sent ticket: *%s* to dingtalk process", ref)
	return nil
}

func insertDingProcessID(processID string) error {
	db := iface.CONTEXT.GetDB().Begin()
	err := db.Table("ding_approve")
	return nil
}

// 标记itop工单为已发送, 字段 *send*
func setItopTicketFlag(ref string) error {
	h := iface.CONTEXT.GetDB().Begin()
	err := h.Table("itop_ticket").Where("ref=?", ref).Update("send", true).Error
	if err != nil {
		h.Rollback()
		return err
	}
	h.Commit()
	return nil
}

func existButNotSend(ref string) bool {
	result := &support.Fileds{}
	h := iface.CONTEXT.GetDB().Table("itop_ticket")
	if isNotFound := h.Where("ref = ? and send = ?", ref, false).Scan(result).RecordNotFound(); isNotFound {
		return false
	}
	return true
}

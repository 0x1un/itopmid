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
	if resp.ErrCode == 40000 {
		// TODO: refresh dingtalk access_token
	}
	if resp.ErrCode != 0 {
		return fmt.Errorf("%s", resp.ErrMsg)
	}
	content.Filed.DingProcessInstanceId = resp.ProcessInstanceId
	iface.LOGGER.Info("Sent ticket: *%s* to dingtalk process", ref)
	return nil
}

func insertDingProcessID(processid string) error {
	db := iface.CONTEXT.GetDB().Begin()
	if err := db.Table("ding_approve").Create(
		struct {
			ProcessId string `gorm:"column:process_id"`
		}{
			ProcessId: processid,
		}).Error; err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
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

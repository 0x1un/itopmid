package core

import (
	"fmt"

	"github.com/0x1un/boxes/dingtalk/api"
	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

// const (
// 	SEND_SUCC_CODE    = 0
// 	SEND_ERR_CODE     = 1
// 	DATABASE_ERR_CODE = 2
// )

func SendSingleTicketToDingtalkProcess(content *support.ResponseContent) error {
	ref := content.Filed.Ref
	form := ConvertSingleUserRequest(content)
	// if isSend(ref) {
	resp, err := iface.CLIENT.SendProcess(*form)
	if resp.ErrCode != 0 || err != nil {
		return fmt.Errorf("error msg: %s\n", resp.ErrMsg)
	}
	if err := setItopTicketFlag(ref); err != nil {
		return err
	}
	iface.LOGGER.Info("Sent ticket: *%s* to dingtalk process", ref)
	// }
	return nil
}

// 这里调用SendProcess批量发送工单
func sendToDingtalkProcess(c *api.DingTalkClient, resp support.UserReqResponse) {
	formComponent := ConvertBatchUserRequest(resp)
	for k, v := range formComponent {
		// 如果查到这个值并没有被发送过的记录，则将其发送
		if isSend(k) {
			response, err := c.SendProcess(v)
			// 如果发送失败，输出失败原因到日志
			if response.ErrCode != 0 || err != nil {
				iface.LOGGER.Error(response.ErrMsg)
				continue
			}
			if err := setItopTicketFlag(k); err != nil {
				iface.LOGGER.Error(err.Error())
			}
			iface.LOGGER.Info("Sent a ticket: %s to dingtalk", k)
		}
	}
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

func isSend(ref string) bool {
	result := &support.Fileds{}
	h := iface.CONTEXT.GetDB().Table("itop_ticket")
	if isNotFound := h.Where("ref = ? and send = ?", ref, false).Scan(result).RecordNotFound(); !isNotFound {
		// !isNotFound if found *send=false* then return true
		return true
	}
	return false
}

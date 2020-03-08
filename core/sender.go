package core

import (
	"github.com/0x1un/boxes/dingtalk/api"
	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

// const (
// 	SEND_SUCC_CODE    = 0
// 	SEND_ERR_CODE     = 1
// 	DATABASE_ERR_CODE = 2
// )

// 发送到钉钉遇到失败，放入此队列等待重试
var retryQueue = make(map[string]api.FormValues)

// 这里调用SendProcess批量发送工单
func SendToDingtalkProcess(c *api.DingTalkClient, resp support.UserReqResponse) {
	formComponent := ConvertUserRequest(resp)
	for k, v := range formComponent {
		response, err := c.SendProcess(v)
		if response.ErrCode != 0 || err != nil {
			iface.LOGGER.Error(response.ErrMsg)
			retryQueue[k] = v
			continue
		}
		if err := setItopTicketFlag(k); err != nil {
			iface.LOGGER.Error(err.Error())
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

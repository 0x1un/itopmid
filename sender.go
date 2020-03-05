package main

import (
	"github.com/0x1un/boxes/dingtalk/api"
	"github.com/0x1un/itopmid/iface"
)

// 这里调用SendProcess批量发送工单
func SendToProv(c *api.DingTalkClient, resp UserReqResponse) {
	formValueArray := ConvertUserRequest(resp)
	for _, v := range formValueArray {
		response, err := c.SendProcess(v)
		if response.ErrCode != 0 || err != nil {
			iface.LOGGER.Error(response.ErrMsg)
		}
		iface.LOGGER.Info("Sent a message succeeded")
	}
}

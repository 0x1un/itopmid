package main

import (
	"github.com/0x1un/boxes/dingtalk/api"
	"strings"
)

type location struct {
	City      string
	Seat      string
	Wb        string
	FaultType string
}

func ConvertUserRequest(resp UserReqResponse) (formValues []api.FormValues) {
	for _, v := range resp.Object {
		location := titleParse(v.Filed.Title, "|")
		fv := api.FillForm(location.City, location.Seat, location.Wb, "13800138000", location.FaultType, "单个台席", v.Filed.Description)
		formValues = append(formValues, fv)
	}
	return
}

// title: city|seat|wb account
// 关于itop中的工单格式如何转换为钉钉自定义的审批表单
// 这里约定一个规则，以钉钉工单为主，兼容钉钉工单表单格式
// 后期可能会以itop的标准门户工单格式为兼容对象
func titleParse(title, sep string) (location location) {
	if len(title) == 0 {
		return
	}
	res := strings.Split(title, sep)
	if len(res) < 3 {
		return
	}
	location.City = res[0]
	location.Seat = res[1]
	location.Wb = res[2]
	location.FaultType = res[3]
	return
}

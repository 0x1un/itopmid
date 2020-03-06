package main

import (
	"strings"

	"github.com/0x1un/boxes/dingtalk/api"
)

type location struct {
	city          string
	seat          string
	domainAccount string
	faultType     string
}

func ConvertUserRequest(resp UserReqResponse) (formValues []api.FormValues) {
	for _, v := range resp.Object {
		location := titleParse(v.Filed.Title, "|")
		fv := api.FillFormTemplate(
			location.city,          // 表单中的城市
			location.seat,          // 台席座号
			location.domainAccount, // 台席的域帐号
			"13800138000",          // 联系方式（手机号）
			location.faultType,     // 故障的类型
			"单个台席",                 // 范围（单个台席/多个台席）
			v.Filed.Description)    // 故障的详细描述
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
	location.city = res[0]
	location.seat = res[1]
	location.domainAccount = res[2]
	location.faultType = res[3]
	return
}

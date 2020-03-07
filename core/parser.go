package core

import (
	"strings"

	"github.com/0x1un/boxes/dingtalk/api"
	"github.com/0x1un/itopmid/support"
)

type location struct {
	city          string
	seat          string
	domainAccount string
	faultType     string
}

func ConvertUserRequest(resp support.UserReqResponse) map[string]api.FormValues {
	var formValues = make(map[string]api.FormValues)
	for _, v := range resp.Object {
		location := titleParse(v.Filed.Title, "|")
		fv := api.FillFormTemplate(
			location.city,          // 表单中的城市
			location.seat,          // 台席座号
			location.domainAccount, // 台席的域帐号
			"13800138000",          // 联系方式（手机号）
			location.faultType,     // 故障的类型
			func() string {
				switch v.Filed.Impact {
				case "1", "2", "部门", "服务":
					return "多个台席"
				case "3", "个体":
					return "单个台席"
				}
				return "单个台席"
			}(), // 范围（单个台席/多个台席）
			strings.Trim(v.Filed.Description, "<p></p>")) // 故障的详细描述
		formValues[v.Filed.Ref] = fv
	}
	return formValues
}

// title: city|seat|wb account
// 关于itop中的工单格式如何转换为钉钉自定义的审批表单
// 这里约定一个规则，以钉钉工单为主，兼容钉钉工单表单格式
// 后期可能会以itop的标准门户工单格式为兼容对象
func titleParse(title, sep string) *location {
	if len(title) == 0 {
		return nil
	}
	res := strings.Split(title, sep)
	if len(res) < 3 {
		return nil
	}
	l := new(location)
	l.city = support.GetStringFromArray(res, 0)
	l.seat = support.GetStringFromArray(res, 1)
	l.domainAccount = support.GetStringFromArray(res, 2)
	l.faultType = support.GetStringFromArray(res, 3)
	return l
}

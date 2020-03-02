package main

import (
	"encoding/json"
	"net/url"
	"strings"
)

// UserRequest structure
type Base struct {
	Code    int    `json:"code"`    // 返回的状态码
	Message string `json:"message"` // 返回的状态消息
}

type Fileds struct {
	Ref                    string `json:"ref" gorm:"column:ref"`                                         // itop工单中的序列号，唯一
	RequestType            string `json:"request_type" gorm:"column:request_type"`                       // 服务请求类型
	ServiceSubcategoryName string `json:"servicesubcategory_name" gorm:"column:servicesubcategory_name"` // 子服务名称 （最终的服务）
	Urgency                string `json:"urgency" gorm:"column:urgency"`                                 // 紧急度
	Origin                 string `json:"origin" gorm:"column:origin"`                                   // 工单来源
	CallerIdFriendlyName   string `json:"caller_id_friendlyname" gorm:"column:caller_id_friendlyname"`   // 工单发起者名称
	Impact                 string `json:"impact" gorm:"column:impact"`                                   // 影响范围
	Title                  string `json:"title" gorm:"column:title"`                                     // 标题
	Description            string `json:"description" gorm:"column:description"`                         // 描述
}

type ResponseContent struct {
	Base
	Class string `json:"class"`  // 所属组件类 (UserRequest)
	Key   string `json:"key"`    // 返回key号码
	Filed Fileds `json:"fields"` // 返回的数据
}

// UserRequest返回的响应内容
type UserReqResponse struct {
	Base                              // 返回的基本消息(错误码，错误信息)
	Object map[string]ResponseContent `json:"objects"` // 返回数据的集合对象
}

// 请求的数据结构体
type RequestData struct {
	Operation    string `json:"operation"`     // 请求操作
	Class        string `json:"class"`         // 请求的类(UserRequest)
	Key          string `json:"key"`           // OQL查询语句
	OutPutFields string `json:"output_fields"` // 需要输出哪些数据（此对应返回数据的Field
}

// 生成请求数据, 需要传入itop具有rest API权限的账户
func NewRestAPIAuthData(auth_user, auth_pwd string) (*strings.Reader, error) {
	req_data := make(url.Values)
	json_data := new(RequestData)
	json_data.Operation = "core/get"
	json_data.Class = "UserRequest"
	json_data.Key = "SELECT UserRequest WHERE operational_status = \"ongoing\""
	json_data.OutPutFields = "ref,request_type,servicesubcategory_name,urgency,origin,caller_id_friendlyname,impact,title,description,contacts_list"
	data, err := json.Marshal(json_data)
	if err != nil {
		return nil, err
	}
	req_data.Add("auth_user", auth_user)
	req_data.Add("auth_pwd", auth_pwd)
	req_data.Add("json_data", string(data))
	return strings.NewReader(req_data.Encode()), nil
}

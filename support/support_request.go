package support

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/0x1un/itopmid/iface"
)

const (
	// user request message
	USER_REQUEST_OPERATION    = "core/get"
	USER_REQUEST_CLASS        = "UserRequest"
	USER_REQUEST_KEY          = "SELECT UserRequest WHERE operational_status = \"ongoing\""
	USER_REQUEST_OUTPUTFIELDS = "ref,request_type,servicesubcategory_name,urgency,origin,caller_id_friendlyname,impact,title,description,contacts_list"

	// contact request message
	PERSON_REQUEST_OPERATION    = "core/get"
	PERSON_REQUEST_CLASS        = "Person"
	PERSON_REQUEST_KEY          = "SELECT Person WHERE friendlyname='%s'"
	PERSON_REQUEST_OUTPUTFIELDS = "phone,name,first_name,mobile_phone"
)

// UserRequest structure
type Base struct {
	Code    int    `json:"code"`    // 返回的状态码
	Message string `json:"message"` // 返回的状态消息
}

type Fileds struct {
	Ref                    string                   `json:"ref" gorm:"column:ref"`                                         // itop工单中的序列号，唯一
	RequestType            string                   `json:"request_type" gorm:"column:request_type"`                       // 服务请求类型
	ServiceSubcategoryName string                   `json:"servicesubcategory_name" gorm:"column:servicesubcategory_name"` // 子服务名称 （最终的服务）
	Urgency                string                   `json:"urgency" gorm:"column:urgency"`                                 // 紧急度
	Origin                 string                   `json:"origin" gorm:"column:origin"`                                   // 工单来源
	CallerIdFriendlyName   string                   `json:"caller_id_friendlyname" gorm:"column:caller_id_friendlyname"`   // 工单发起者名称
	Impact                 string                   `json:"impact" gorm:"column:impact"`                                   // 影响范围
	Title                  string                   `json:"title" gorm:"column:title"`                                     // 标题
	Description            string                   `json:"description" gorm:"column:description"`                         // 描述
	Contacts               []map[string]interface{} `json:"contacts_list" gorm:"-"`
	MobilePhone            string                   `json:"mobile_phone" gorm:"-"`
	Phone                  string                   `json:"phone" gorm:"-"`
	Name                   string                   `json:"name" gorm:"-"`
	FirstName              string                   `json:"first_name" gorm:"-"`
	DingProcessInstanceId  string                   `json:"-" gorm:"column:processid"`
	Resolved               bool                     `json:"-" gorm:"column:resolved"`
	IsSend                 bool                     `json:"-" gorm:"column:send"`
}

type ResponseContent struct {
	Base
	Class string `json:"class"`           // 所属组件类 (UserRequest)
	Key   string `json:"key"`             // 返回key号码
	Filed Fileds `json:"fields" gorm:"-"` // 返回的数据
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
type UpdateKey struct {
	Status string `json:"status"`
	Ref    string `json:"ref"`
}

type UpdateFields struct {
	Status string `json:"status"`
}

type RequestUpdate struct {
	Operation    string       `json:"operation"`
	Class        string       `json:"class"`
	Comment      string       `json:"comment"`
	OutputFields string       `json:"output_fields"`
	Key          UpdateKey    `json:"key"`
	Fields       UpdateFields `json:"fields"`
}

func NewRequestUpdateData(cls string, key *UpdateKey, fields *UpdateFields) *strings.Reader {
	rd := &RequestUpdate{
		Operation:    "core/update",
		Class:        cls,
		Comment:      "itopmid updates",
		Key:          *key,
		OutputFields: "status",
		Fields:       *fields,
	}
	jd, err := json.Marshal(rd)
	if err != nil {
		iface.LOGGER.Error(err.Error())
		return nil
	}
	param := make(url.Values)
	param.Set("auth_user", iface.ITOP_USERNAME)
	param.Set("auth_pwd", iface.ITOP_PASSWORD)
	param.Set("json_data", string(jd))
	return strings.NewReader(param.Encode())
}

func (self *RequestData) GenUserRequest() *strings.Reader {
	reqd := make(url.Values)
	br := buildRequest(
		USER_REQUEST_OPERATION,
		USER_REQUEST_CLASS,
		USER_REQUEST_KEY,
		USER_REQUEST_OUTPUTFIELDS,
	)
	reqd.Set("auth_user", iface.ITOP_USERNAME)
	reqd.Set("auth_pwd", iface.ITOP_PASSWORD)
	reqd.Set("json_data", string(br))
	return strings.NewReader(reqd.Encode())
}

// 生成获取Person(个人联系人)的请求信息
func (self *RequestData) GenPersonRequest(friendlyname string) *strings.Reader {
	reqd := make(url.Values)
	br := buildRequest(
		PERSON_REQUEST_OPERATION,
		PERSON_REQUEST_CLASS,
		fmt.Sprintf(PERSON_REQUEST_KEY, friendlyname),
		PERSON_REQUEST_OUTPUTFIELDS,
	)
	reqd.Set("auth_user", iface.ITOP_USERNAME)
	reqd.Set("auth_pwd", iface.ITOP_PASSWORD)
	reqd.Set("json_data", string(br))
	return strings.NewReader(reqd.Encode())
}

func buildRequest(
	operation, class, key, output string) []byte {
	r := &RequestData{
		Operation:    operation,
		Class:        class,
		Key:          key,
		OutPutFields: output,
	}
	data, err := json.Marshal(r)
	if err != nil {
		iface.LOGGER.Error(err.Error())
	}
	return data
}

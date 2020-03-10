package core

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

// 返回来自itop的标准门户工单数据
func FetcheFromITOP(url string, data io.Reader) {
	resp, err := request(http.MethodPost, url, data)
	if err != nil {
		iface.LOGGER.Panic(err.Error())
	}
	t := new(support.UserReqResponse)
	if err := json.Unmarshal(resp, t); err != nil {
		iface.LOGGER.Panic(err.Error())
	}
	// 将获取的工单数据插入到数据库
	for _, v := range t.Object {
		// 如果获取的工单在数据库中不存在，则创建一条新记录
		ref := v.Filed.Ref
		if !checkEntry(ref) || isSend(ref) {
			if err := SendSingleTicketToDingtalkProcess(&v); err != nil {
				iface.RETRY_QUEUE.Push(v)
				iface.LOGGER.Error("Failed to send to dingtalk process, push to retry queuer and retry send")
				continue
			}
			v.Filed.IsSend = true
			if err := insertTicketITOP(v.Filed); err != nil {
				iface.LOGGER.Error("Got error: %s", err.Error())
				continue
			}
			iface.LOGGER.Info("ref: %s is inserted", ref)
		} else {
			iface.LOGGER.Error("%s entry is already exists!", ref)
			continue
		}
	}
	// SendToDingtalkProcess(iface.CLIENT, *t)
}

// 对数据库插入itop工单数据，插入的数据为Fileds中的工单详情
func insertTicketITOP(ticket support.Fileds) error {
	var e error
	h := iface.CONTEXT.GetDB().Begin()
	h = h.Table("itop_ticket")
	e = h.Create(ticket).Error
	if e != nil {
		h.Rollback()
		return e
	}
	h.Commit()
	return nil
}

// 判断itop的工单是否已经被获取, 判断的依据是itop中工单唯一编码ref
// 如果没有 返回false, 反之亦然
func checkEntry(ref string) bool {
	h := iface.CONTEXT.GetDB().Begin()
	h = h.Table("itop_ticket")
	nf := h.Select("ref").Where("ref=?", ref).Scan(&struct{ Rf string }{}).RecordNotFound()
	if nf {
		return false
	} else {
		return true
	}
}

// 简单封装的http请求
func request(method, url string, data io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	switch data.(type) {
	case *strings.Reader:
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	case *bytes.Reader:
		req.Header.Add("Content-Type", "application/json")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

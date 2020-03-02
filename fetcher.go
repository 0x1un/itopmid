package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
)

func ListenITOP(url string, data io.Reader) <-chan UserReqResponse {

	return nil
}

// 返回来自itop的标准门户工单数据
func FetcheFromITOP(url string, data io.Reader) UserReqResponse {
	resp, err := request(http.MethodPost, url, data)
	if err != nil {
		panic(err)
	}
	t := new(UserReqResponse)
	if err := json.Unmarshal(resp, t); err != nil {
		panic(err)
	}
	return *t
}

// 对数据库插入itop工单数据，插入的数据为Fileds中的工单详情
func StoreTicketFromITOP(conn *gorm.DB, ticket Fileds) (e error) {
	h := conn.Begin()
	h = h.Table("itop_ticket")
	isNotFound := h.Select("ref").Where("ref=?", ticket.Ref).Scan(&struct{ Ref string }{}).RecordNotFound()
	if isNotFound {
		e = h.Create(ticket).Error
		if e != nil {
			h.Rollback()
			return e
		}
	} else {
		return errors.New(fmt.Sprintf("错误: 重复插入条目: %s", ticket.Ref))
	}
	h.Commit()
	return nil
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

package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

const (
	DEFAULT_MOBILE_PHONE = "13800138000"
)

func FetcheFromITOP(url string, data io.Reader) {
	resp, err := request(http.MethodPost, url, data)
	if err != nil {
		iface.LOGGER.Panic(err.Error())
	}
	t := new(support.UserReqResponse)
	if err := json.Unmarshal(resp, t); err != nil {
		iface.LOGGER.Panic(err.Error())
	}
	for _, v := range t.Object {
		ref := v.Filed.Ref
		phone := DEFAULT_MOBILE_PHONE
		if friendlyname := extractFriendlyNameByContact(v.Filed.Contacts); friendlyname != "" {
			reqData := iface.REQUEST.GenPersonRequest(friendlyname)
			presp, err := request(http.MethodPost, url, reqData)
			if err != nil {
				iface.LOGGER.Panic(err.Error())
			}
			// rct = responseContent
			rct := new(support.UserReqResponse)
			if err := json.Unmarshal(presp, rct); err != nil {
				iface.LOGGER.Error(err.Error())
			}
			fmt.Println(rct)
			for _, x := range rct.Object {
				if p := x.Filed.MobilePhone; len(p) != 0 {
					phone = p
				}
			}
		}
		if entryNotFound(ref) {
			v.Filed.MobilePhone = phone
			if err := SendSingleTicketToDingtalkProcess(&v); err != nil {
				iface.LOGGER.Debug(fmt.Sprintf("Failed to send ticket: %s", err.Error()))
				iface.RETRY_QUEUE.Push(v)
				continue
			}
			v.Filed.IsSend = true
			if err := insertTicketITOP(v.Filed); err != nil {
				iface.LOGGER.Debug("Got error: %s", err.Error())
				continue
			}
			iface.LOGGER.Info("ref: %s is inserted", ref)
		} else {
			iface.LOGGER.Debug("Entry may already exist or sended")
		}
	}
}

func insertTicketITOP(ticket support.Fileds) error {
	var err error
	dbCtx := iface.CONTEXT.GetDB().Begin()
	dbCtx = dbCtx.Table("itop_ticket")
	err = dbCtx.Create(ticket).Error
	if err != nil {
		dbCtx.Rollback()
		return err
	}
	dbCtx.Commit()
	return nil
}

func entryNotFound(ref string) bool {
	h := iface.CONTEXT.GetDB()
	h = h.Table("itop_ticket")
	nf := h.Select("ref").Where("ref=?", ref).Scan(&struct{ Rf string }{}).RecordNotFound()
	if nf {
		return true
	}
	return false
}

func extractFriendlyNameByContact(ctt []map[string]interface{}) string {
	for _, v := range ctt {
		if len(v) != 0 {
			return v["contact_id_friendlyname"].(string)
		}
	}
	return ""
}

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

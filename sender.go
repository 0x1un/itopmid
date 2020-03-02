package main

import (
	"github.com/0x1un/boxes/dingtalk/api"
	"errors"
	"fmt"
	"log"
)

func SendToProv(c *api.DingTalkClient, resp UserReqResponse) error {
	formValueArray := ConvertUserRequest(resp)
	for _, v := range formValueArray {
		response, err := c.SendProcessForTest(v)
		if response.ErrCode != 0 && err != nil {
			return errors.New(fmt.Sprintf("%s", response.ErrMsg))
		}
		log.Printf("Sent a ticket succeed! status code: %d", response.ErrCode)
	}
	return nil
}

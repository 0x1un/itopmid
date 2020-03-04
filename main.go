package main

import (
	_ "net/http/pprof"

	"github.com/0x1un/itopmid/iface"
	"github.com/0x1un/itopmid/support"
)

// 釘釘應用程序的agentid
const (
// iTopLocalURL = `http://localhost:8000/webservices/rest.php?version=1.3`
)

func init() {
	// load config
	config := &support.ItopMidConfig{}
	config.ReadConfigFile("itopmid.json")
	iface.CONFIG = config

	// load logger
	logger := &support.ItopMidLogger{}
	iface.LOGGER = logger
}

func main() {

	iface.LOGGER.Debug("这特么只是一个测试")
	iface.LOGGER.Debug("%s", iface.CONFIG.GetDatabaseURL())
	// mux := http.NewServeMux()
	// mux.HandleFunc("/custom_debug_path/profile", pprof.Profile)
	// log.Fatal(http.ListenAndServe(":7777", mux))
	// request_data, err := NewRestAPIAuthData("admin", "...")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// conn, err := db.NewDBConnect()
	// if err != nil {
	// 	panic(err)
	// }

	// 从itop中获取所有状态为开启的工单
	// resp := FetcheFromITOP(iTopLocalURL, request_data)
	// for _, v := range resp.Object {
	// 	fmt.Println(v.Filed)
	// 	if err := StoreTicketFromITOP(conn, v.Filed); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// client := api.NewClient(api.APPKEY, api.APPSECRET)
	// // 发送来自itop的工单
	// if err := SendToProv(client, resp); err != nil {
	// 	panic(err)
	// }
}

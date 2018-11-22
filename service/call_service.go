package service

import (
	"context"
	"file_pool_api/conf"
	"github.com/smallnest/rpcx/client"
	"helper_go/comhelper"
	"log"
)

const (
	// 正常接口返回
	SERVICE_SUCCESS = 0

	// 服务名称
	FilePoolService = "FilePoolService"
)

type Out struct {
	OutMsg  interface{} `json:"outMsg"`
	OutData interface{} `json:"outData"`
}

type Reply struct {
	Out *Out `msg:"Out"`
}

/**
 * 远程访问rpc获取数据
 * @param service string 服务名称
 * @param method string 方法名称
 * @param args map[string]interface{} 参数
 */
func CallService(service, method string, args map[string]interface{}) (Out, error) {
	rpc_addr := conf.GetConfig().ProxyServer

	d := client.NewPeer2PeerDiscovery("tcp@"+rpc_addr, "")
	xclient := client.NewXClient(service, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	reply := &Reply{}
	err := xclient.Call(context.Background(), method, args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
		return Out{}, err
	}
	return *reply.Out, nil
}

/**
 * 调用rpc获取数据
 */
func _call(service, method string, args map[string]interface{}) map[string]interface{} {
	ret, _ := CallService(service, method, args)
	var data map[string]interface{}
	data = make(map[string]interface{})
	if ret.OutMsg != "" {

		switch ret.OutData.(type) {
		case float64:
			data["error"] = int(ret.OutData.(float64))
		case string:
			data["error"] = comhelper.StringToInt(ret.OutData.(string))
		default:
			data["error"] = ret.OutData
		}
		data["data"] = ret.OutMsg
	} else {
		data["error"] = SERVICE_SUCCESS
		data["data"] = ret.OutData
	}
	if data["error"] == nil {
		data["error"] = conf.SERVER_ERROR
	}
	return data
}

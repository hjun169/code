package main

import (
	"fmt"
	"net/rpc"
	"utils"
)

func main(){
	address,err := utils.GetConfig("RPC_ADDRESS")//获取RPC服务地址
	utils.OutputError(err)
	
	port,err := utils.GetConfig("RPC_PORT")//获取RPC服务端口
	utils.OutputError(err)
	
	service := address + ":" + port
	
	client, err := rpc.Dial("tcp", service)	
	utils.OutputError(err)
	
	args := "test rpc"
	var reply utils.OutputData
	err = client.Call("Output.SetData", args, &reply)
	utils.OutputError(err)
	
	fmt.Println(args, reply)
}
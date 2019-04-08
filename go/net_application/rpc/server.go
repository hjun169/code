package main

import (
	"net/rpc"
	"net"
	"utils"
)

type Output utils.OutputData 

func (t *Output) SetData(val string, reply *Output) error {
	reply.Code = 200
	reply.Msg  = ""
	reply.Data = val
	return nil
}

func main(){
	address,err := utils.GetConfig("RPC_ADDRESS")//获取RPC服务地址
	utils.OutputError(err)
	
	port,err := utils.GetConfig("RPC_PORT")//获取RPC服务端口
	utils.OutputError(err)
	
	service := address + ":" + port
	
	output := new(Output)
	rpc.Register(output)
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	utils.OutputError(err)
	
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.OutputError(err)
	defer listener.Close()
	
	rpc.Accept(listener)
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		
		rpc.ServeConn(conn)
	}
}
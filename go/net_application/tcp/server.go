package main

import (
	"fmt"
	"net"
	"time"
	"utils"
)

func main() {
	address,err := utils.GetConfig("TCP_ADDRESS")//获取TCP服务地址
	utils.OutputError(err)
	
	port,err := utils.GetConfig("TCP_PORT")//获取TCP服务端口
	utils.OutputError(err)
	
	service := address + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	utils.OutputError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.OutputError(err)
	defer listener.Close()
	
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}
func handleConn(conn net.Conn) {
	defer conn.Close()
	i := 0	
	hb := make(chan int)
	defer close(hb)
	
	for{
		go heaterbeat(hb, conn) //心跳检测
		
		i++
		j := fmt.Sprintf("%d", i)
		
		buffer := make([]byte, 2048)
		n, err := conn.Read(buffer)		
		if err != nil {
			utils.LogMessage("error", err.Error())
			break
		}
		
		hb <- 1 //维持心跳
		
		output, err_decode := utils.JsonDecode(buffer[:n])
		code := 0
		msg  := ""
		data := ""
		
		if err_decode != nil {
			code = 500
			msg  = "接受到的第" + j + "条消息,解析消息出现错误:" + err_decode.Error()
			data = ""
		}else{
		    typeRight := 0
			var typeData interface{}
			typeData = output
			switch typeData.(type) {//检测类型
				case utils.OutputData:
					typeRight = 1		
			}
			
			if typeRight == 0 {
				code = 500
				msg  = "接受到的第" + j + "条消息,消息格式错误"
				data = ""	
			} else {			
			
				clientData := fmt.Sprintf("code:%d msg:%s data:%v", output.Code, output.Msg, output.Data)
				if output.Code == 200 {
					code = output.Code
					msg  = ""
					data = "接受到的第" + j + "条消息:" + clientData
				}else if output.Code == 600 {//客户端停止
					break
				}else{	
					code = output.Code
					msg  = "接受到的客户端传来第" + j + "条错误消息:" + clientData
					data = ""
				}
			}
		}

		encode, err_enocde := utils.JsonEncode(code, msg, data)
		if err_enocde == nil {
			conn.Write(encode)
		} else {
			fmt.Println(err_enocde)
		}
		
	}
	
}

//心跳检测函数
func heaterbeat(hb chan int, conn net.Conn) {
	select {
		case <-hb :
			return
		case <-time.After(time.Second * 30): //30秒无接受消息则停止
			conn.Close()		
	}
}

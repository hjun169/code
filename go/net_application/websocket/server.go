package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"utils"
)

type Output utils.OutputData 

func ReceiveOutput(conn *websocket.Conn){
	defer conn.Close()
	var output Output
	for{
		err := websocket.JSON.Receive(conn, &output)
		if err != nil {
			fmt.Println("Receive data Error:", err)
			break
		}else{
			var sendData Output
			sendData.Code = output.Code
			sendData.Msg  = ""
			sendData.Data = ""
			
			clientData := fmt.Sprintf("code:%d msg:%s data:%v", output.Code, output.Msg, output.Data)
			if sendData.Code == 200 {
				sendData.Data = clientData;
			} else {//code非200即错误消息
				sendData.Msg = clientData;
			}
			
			errSend := websocket.JSON.Send(conn, sendData)
			if errSend != nil {
				utils.LogMessage("error", errSend.Error())
				fmt.Println("Send data Error:", errSend)
				break
			}
		}
	}
}

func main(){
	address,err := utils.GetConfig("WS_ADDRESS")//获取websocket服务地址
	utils.OutputError(err)
	
	port,err := utils.GetConfig("WS_PORT")//获取websocket服务端口
	utils.OutputError(err)
	
	service := address + ":" + port
	
	http.Handle("/", websocket.Handler(ReceiveOutput))
	err_listen := http.ListenAndServe(service, nil)
	utils.OutputError(err_listen)
}
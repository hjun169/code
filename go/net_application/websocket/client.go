package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"utils"
	"os"
	"io"
	"strings"
	"bufio"
)

type Output utils.OutputData 

func main(){
	address,err := utils.GetConfig("WS_ADDRESS")//获取websocket服务地址
	utils.OutputError(err)
	
	port,err := utils.GetConfig("WS_PORT")//获取websocket服务端口
	utils.OutputError(err)
	
	service := address + ":" + port

	conn, err := websocket.Dial("ws://" + service, "", "http://" + service)
	utils.OutputError(err)
	defer conn.Close()
		
	fmt.Println("请按行输入字符,quit退出:")
	fh := bufio.NewReader(os.Stdin)
	for{
		line, err := fh.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break;
			}
			
			continue
		}
		
		str := strings.TrimSpace(line)//去除空格,包括\r
		if str == "quit" {
			break
		}
		
		var output Output
		output.Code = 200
		output.Msg = ""
		output.Data = str
		
		errSend := websocket.JSON.Send(conn, output)
		if errSend != nil {
			utils.LogMessage("error", errSend.Error())
			fmt.Println("Send data Error:", errSend)
			break
		}
		
		var receive Output
		err_receive := websocket.JSON.Receive(conn, &receive)
		if err_receive != nil {
			utils.LogMessage("error", err_receive.Error())
			fmt.Println("Receive data Error:", err_receive)
			break
		}else{
			fmt.Println("Serve reply:", receive)	
		}
	}
	
	os.Exit(0)
}
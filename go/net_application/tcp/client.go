package main

import (
	"fmt"
	"net"
	"os"
	"utils"
	"strings"
	"bufio"
	"io"
)

func main() {
	address,err := utils.GetConfig("TCP_ADDRESS")//获取TCP服务地址
	utils.OutputError(err)
	
	port,err := utils.GetConfig("TCP_PORT")//获取TCP服务端口
	utils.OutputError(err)
	
	service := address + ":" + port
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)//tcp
	utils.OutputError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	utils.OutputError(err)
	defer conn.Close()
	
	fmt.Println("请按行输入字符,quit退出或30秒无输入自动断开连接:")
	fh := bufio.NewReader(os.Stdin)
	for{
		line, err := fh.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				encode, err_encode := utils.JsonEncode(600, "stop by use", "")
				if err_encode == nil {
					conn.Write(encode)
				}
				break;
			}
			
			continue
		}
		
		str := strings.TrimSpace(line)//去除空格,包括\r
		if str == "quit" {
			encode, err_encode := utils.JsonEncode(600, "stop by use", "")
			if err_encode == nil {
				conn.Write(encode)
			}
			break
		}
		
		encode, err_encode := utils.JsonEncode(200, "", str)
		if err_encode == nil {
			_, err_write := conn.Write(encode)
			if err_write == nil {
				buffer := make([]byte, 2048)
				n, err_read := conn.Read(buffer)		
				if err_read == nil {
					output, err_decode := utils.JsonDecode(buffer[:n])
					if err_decode == nil {
						fmt.Println("server reply:", output)
					}
				} else if err_read == io.EOF {
					fmt.Println("server断开")
					break
				} else {
					fmt.Println("server reply:", err_read)
					break	
				}
				
			} else {
				fmt.Println("server reply:", err_write)
				break
			}			
		}
	}
	
	os.Exit(0)
}
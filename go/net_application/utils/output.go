//处理数据交互格式文件

package utils

import (	
	"encoding/json"
	"fmt"
)

type OutputData struct {
	Code int         //错误码,200为正确数据
	Msg  string      //错误码非200则非空错误消息
	Data interface{} //错误码200对应数据
}

//json encode数据
func JsonEncode(code int, msg string, data interface{}) ([]byte, error) {
	var output OutputData = OutputData{
		Code: code,
		Msg: msg,
		Data: data,
	}
	
	result,err := json.Marshal(output)
	
	if err != nil {
		errMsg := fmt.Sprintf("code:%d msg:%s data:%v", code, msg, data)	
		LogMessage("error", errMsg)
	}
	
	return result, err
}

//json decode数据
func JsonDecode(encode []byte) (OutputData, error) {
	output := OutputData{}
	err := json.Unmarshal(encode, &output)
	if err != nil {
		errMsg := string(encode)	
		LogMessage("error", errMsg)
	}
	
	return output, err
}

//处理配置文件

package utils

import (
	"os"
	"bufio"
	"io"
	"strings"
	"errors"
	"log"
)


type ConfigEnv map[string]string
var configs ConfigEnv = make(ConfigEnv, 10)

var errDateRecord error = errors.New("Env config has wrong data")
var errNoRecord error = errors.New("Not found env config")
var errNoKey error = errors.New("Not found config key")

//加载程序配置文件
func init(){
	err := LoadConfig("../.env")
	if err != nil {
		log.Fatalln("Failed to load config file", err)
	}
}

//加载配置文件到configs变量中
func LoadConfig(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {		
		return err
	}

	defer fh.Close()

	reader := bufio.NewReader(fh)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break;
			}
			
			return err
		}
		
		arrs := strings.Split(line, "#")
		str := strings.TrimSpace(arrs[0])//去除空格,包括\r
		if str == "" {//为#注释行
			continue
		}
		
		//获取每行记录对应的key value
		vals := strings.Split(str, "=")
		if len(vals) != 2 {
			return errDateRecord
		}
		
		vals[0] = strings.TrimSpace(vals[0]);
		vals[1] = strings.TrimSpace(vals[1]);
		configs[vals[0]] = vals[1];
	}
	
	if len(configs) < 1 {
		return errNoRecord
	}
	
	return nil
}

//获取配置文件
func GetConfig(key string) (string, error) {
	result, ok := configs[key]
	if ok != true {
		return result, errNoKey
	}
	
	return result, nil
}

//获取所有配置
func GetConfigs() map[string]string {
	return configs
}

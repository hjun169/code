//配置文件测试

package utils

import (
	. "utils"
	"testing"
	"fmt"
)

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("test.env")
	if err != nil {
		t.Errorf("LoadConfig test failed: got %s", err.Error())
	}
	
	configs := GetConfigs()
	fmt.Println("Right results:\n", configs);	
}

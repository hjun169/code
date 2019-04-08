//配置文件测试

package utils

import (
	. "utils"
	"testing"
)

func TestLogMessage(t *testing.T) {
	LogMessage("error", "testing error")
	LogMessage("debug", "testing debug")
	LogMessage("info", "testing info")
	LogMessage("warning", "testing warning")		
}

package task

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"yulong-hids/daemon/common"
)

// KillProcess 根据进程名结束进程
func KillProcess(processName string) string {
	var data string
	path := strings.Split(strings.Split(processName, " ")[0], "/")
	processName = path[len(path)-1]
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-1\.\-_]+$`, processName); !ok {
		return ""
	}
	if runtime.GOOS == "windows" {
		data, _ = common.CmdExec("taskkill.exe /f /im " + processName)
	} else {
		data, _ = common.CmdExec(fmt.Sprintf("kill -9 $(pidof %s)", processName))
	}
	return data
}

package task

import (
	"fmt"
	"runtime"
	"strings"
	"yulong-hids/daemon/common"
)

// KillProcess 根据进程名结束进程
func KillProcess(processName string) (string, error) {
	var data string
	var err error
	path := strings.Split(strings.Split(processName, " ")[0], "/")
	processName = path[len(path)-1]
	if runtime.GOOS == "windows" {
		data, err = common.CmdExec("taskkill.exe /f /im " + processName)
	} else {
		data, err = common.CmdExec(fmt.Sprintf("kill -9 $(pidof %s)", processName))
	}
	return data, err
}

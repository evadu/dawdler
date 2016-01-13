package ext

import (
	"testing"
	"fmt"
	"os/exec"
	"time"
)


func TestExt(t *testing.T) {
	loadIni()
	for 1==2{
		//如果当一个服务不小心挂掉了，需要重启动
		cmd:=exec.Command("tasklist")
		content,_:=cmd.Output()
		fmt.Println(string(content))
		time.Sleep(100*time.Second)
	}
}
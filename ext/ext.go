package ext
import(
	"../comm"
	"fmt"
	"time"
	"os/exec"
)
var ini comm.Ini
//扩展加载 都是长线程的服务型
func loadIni(){
	//读取需要开启的服务
	ini=comm.Ini{}
	ini.IniName="ext.ini"
	ini.ReadAll()
	go monitoring()
}
func monitoring(){
	for{
		if ini.CheckModf(){
			ini.ReadAll()
		}
		time.Sleep(time.Second*1)
	}
}

func main(){
	loadIni()
	for{
		//如果当一个服务不小心挂掉了，需要重启动
		cmd:=exec.Command("tasklist")
		fmt.Println(cmd.Output())
		time.Sleep(5*time.Second)
		
	}
}
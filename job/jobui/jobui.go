package jobui
import(
	"fmt"
	"net/http"
	"net"
)
//临时的东西搞搞，别那么认为
type JobUi struct{
	l net.Listener
	Runing bool
	load bool
}
func(this *JobUi) JobUiStart(Port string){
	if this.Runing==true{
		return
	}
	if this.load!=true{
		this.load=true
		http.HandleFunc("/jobui",router)
	}
	this.Runing=true
	server := &http.Server{Addr:":"+Port,Handler: nil}
	addr := ":"+Port
	if addr == "" {
		addr = ":http"
	}
	this.l, _ = net.Listen("tcp", addr)
	server.Serve(this.l)
}
func(this *JobUi) JobUIStop(){
	if this.Runing==false{
		return
	}
	fmt.Println("关闭了~ 应该收不到了吧！")
	this.l.Close()
	this.Runing=false
}
func router(writer http.ResponseWriter, req *http.Request) {
	path:=req.URL.RequestURI()
	fmt.Println(path)
//	frame := &view.Frame{Tpl: def.VIEW_PATH+"html"+path+def.VIEW_EXT}
//	if frame!=nil{
//		writer.Write(frame.Visible(true))
//	}
//	if path=="/savejob"{
//		req.ParseForm()
//	}
}


func saveJobIni(){
	
}
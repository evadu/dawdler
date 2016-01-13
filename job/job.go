package main
import(
//	"os/exec" //最重要的
	"./vo"
	"time"
	"../comm"
	"./jobui"
	"fmt"
	"os"
)
const(
	JOBUI="JOBUI" //JOBUI 是否开启
	JOBUI_PORT="JOBUI_PORT" //JOBUI的PORT
	JOB_DB_NAME="JOB_DB_NAME" //job保存的文件
	JOB_OS="JOB_OS" //操作系统 1为WINDOWS 2LINUX
)
var(
	job *Job
	JobUi *jobui.JobUi
	ini *comm.Ini
)

//启动
func main(){
	LoadInit()	//读取需要开启的服务
	NewJob()//真正的JOB线程
	InitThread()//保证JOB一直开启，并一直打开OR关闭JOBUI
}

func LoadInit(){
	JobUi=&jobui.JobUi{}
	ini=&comm.Ini{}
	ini.IniName="job.ini"
	ini.ReadAll()
	if ini.Values[JOB_DB_NAME]==""{
		ini.Values[JOB_DB_NAME]="./job.db"
	}
	go SetJobUI()
}
func InitThread(){
	for{
		if ini.CheckModf(){
			ini.ReadAll()
			go SetJobUI()
		}
		time.Sleep(time.Second*1)
	}
}
func SetJobUI(){
	if ini==nil{
		 return
	}
	status:=ini.Values[JOBUI]
	if status=="ON"{
		JobUi.JobUiStart(ini.Values[JOBUI_PORT])
	}else{
		JobUi.JobUIStop()
	}
}
//声明变量
type Job struct{
	currentDate vo.JobDate //当前时间
	JobList []vo.JobVo //开启的JOB
	NextRunTime int64 //下次执行的时间秒数	
}
func NewJob(){
	job=&Job{}
	job.currentDate.Month=int(time.Now().Month())
	job.currentDate.Week=int(time.Now().Weekday())
	job.currentDate.Day=time.Now().Day()
	job.currentDate.Hour=time.Now().Hour()
	job.currentDate.Minue=time.Now().Minute()
	job.currentDate.Second=time.Now().Second()
	//获取当前时间，进入需要执行的队列排序
	job.JobList=make([]vo.JobVo,0)
	job.LoadJob()
}
//需要跑的JOB 先简单点，就只有10个JOB，有100个劈死他
func (this *Job) LoadJob() (err error){
	//如果发现没有就需要创建 Job.db
	var(
		dbInfo os.FileInfo
	)
	filename:=ini.Values[JOB_DB_NAME]
	if dbInfo,err=os.Stat(filename);os.IsNotExist(err){
		_,err= os.Create(filename) //创建文件
		return
	}
	//不是新建的，读取JOBDB
	fmt.Println("TODO 去读取吧",dbInfo)
	return
}
//什么时候需要跑
func (this *Job) monitoring(){
	
	for{
		//文件监控
		//发生变化的时候,进行读入
		
		//时间排队 年 月 日 星期 小时 分 秒
		
	}
}
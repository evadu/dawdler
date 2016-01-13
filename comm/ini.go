package comm
import(
	"time"
	"bufio"
	"os"
	"io"
)
type Ini struct{
	IniName string//文件名称
	LastMdfTime time.Time//系统时间
	Values map[string]string//配置的值
}

func (this *Ini) ReadAll() (err error){
	initMap:=make(map[string]string)
	//CHECK文件是否存在
	var (
		fi os.FileInfo
	)
	if fi,err = os.Stat(this.IniName); os.IsNotExist(err) {
		return err
	}
	this.LastMdfTime=fi.ModTime()
	//CHECK文件是否能打开
	iniFile, err1 := os.Open(this.IniName)
	if err1!=nil{
		return err1
	}
	//逐行读取，返回到当前的Init
	rd := bufio.NewReader(iniFile)
    for {
        line, err := rd.ReadBytes(byte('\n')) //以'\n'为结束符读入一行        
		if line!=nil{
			key,val:=getVal(line)
			if val!=""{
				initMap[key]=val
			}
		}
        if err != nil || io.EOF == err {
            break
        }
		//将line解析
    }
	this.Values=initMap
	return
}
func getVal(bits []byte) (key string,value string){
	keys:=make([]byte,0)
	vals:=make([]byte,0)
	isKey:=true
	annotations:=0
	for _,bite:=range bits{
		if bite==byte('='){
			isKey=false
			continue
		}
		if bite==byte('-'){ //两个就是注释
			if annotations==1{
				break
			}
			annotations=1;
			continue
		}
		if annotations==1{
			annotations=0
			keys=append(keys,1)
		}
		if isKey{
			keys=append(keys,bite)
		}else{
			vals=append(vals,bite)
		}
	}
	//去掉空格
	keys=Trim(keys)
	vals=Trim(vals)
	return string(keys),string(vals)
}

func (this *Ini) CheckModf() bool{
	var (
		fi os.FileInfo
		err error
	)
	if fi,err = os.Stat(this.IniName); os.IsNotExist(err) {
		return false
	}
	return fi.ModTime().After(this.LastMdfTime)
}
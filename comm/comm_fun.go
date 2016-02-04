package comm
import(
	"reflect"
	"fmt"
	"strings"
)
//去前后空格
func Trim(vals []byte)[]byte{
	i:=len(vals)
	for ;i>0;i--{
		if(vals[i-1]!=byte(' ')){
			break
		}
	}
	j:=0
	for ;j<i;j++{
		if(vals[j]!=byte(' ')){
			break
		}
	}
	return vals[j:i]
}

func Struct2Map(obj interface{}) (data map[string]interface{}) {
	data = make(map[string]interface{})
	_type := reflect.TypeOf(obj)
	if strings.Contains(_type.String(), "map") {
		data:=reflect.ValueOf(obj).Value
		fmt.Println(obj,data)
//		for key,val:=range data{
//			data[key]=val
//		}
	}
	
//	for i := 0; i < t.NumField(); i++ {
//		data[t.Field(i).Name] = v.Field(i).Interface()
//	}
	return
}
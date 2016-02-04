package comm

import (
	"testing"
	"fmt"
)


func TestReadIni(t *testing.T) {
	ini:=Ini{IniName:"test"}
	ini.ReadAll()
	fmt.Println(ini)
}
func TestTrim(t *testing.T) {
	fmt.Println(string(Trim([]byte("   Hello World  "))),"_我去_","   Hello World  ")
}
type Tes struct{
	name string
}
func TestStruct2Map(t *testing.T){
	mymap:=map[string]string{"aaa":"bbb","bbb":"aaa"};
	Struct2Map(mymap);
}
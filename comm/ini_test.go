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
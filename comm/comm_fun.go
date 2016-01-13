package comm
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
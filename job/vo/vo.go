package vo

type JobDate struct{
	Month int //1~12
	Week int //1~7
	Day int //1~37
	Hour int //0~24
	Minue int //0~60
	Second int //0~60	
}

type JobVo struct{
	Date JobDate //执行时间
	Name string	//名称
	Path string //路径
	Status bool //状态
	Params map[string]interface{} //参数 JSON参数
}

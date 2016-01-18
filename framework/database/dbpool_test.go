package database
import (
	"testing"
	"fmt"
)


func TestDataBase(t *testing.T) {
	var mysql DB=&MySqlDb{}
	dbsource:=DataSource{Db:mysql}
	pool:=NewDbPollVo(dbsource,2,50)
	
	fmt.Println("拿到一个DATASOURCE",pool.pool)
	dSource:=pool.GetDataSource()
	dSource.Insert()
}

type MySqlDb struct{
	
}
func (this *MySqlDb) CreateDataSource(params ...interface{}) interface{}{
	return this
}
func (this *MySqlDb) Insert(params ...interface{}) interface{}{
	fmt.Println(params)
	fmt.Println("添加的参数一共有",len(params))
	//开始解析参数
	return 0
}
func (this *MySqlDb) Update(params ...interface{}) interface{}{
	return 0
}
func (this *MySqlDb) Delete(params ...interface{}) interface{}{
	return 0
}
func (this *MySqlDb) Get(params ...interface{}) interface{}{
	return 0
}
func (this *MySqlDb) Release(){
}
func(this *MySqlDb) Close(){	
}
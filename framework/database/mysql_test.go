package database

import (
	"fmt"
	"testing"
)

var pool *DBpoolVo

func TestDataBase(t *testing.T) {
	var mysql DB = &MySql{DBlink: "root:@/dawdler?charset=utf8"}
	dbsource := DataSource{Db: mysql}
	pool = NewDbPollVo(dbsource, 1, 50)
	//	dSource := pool.GetDataSource()
	//接口方法
	getTableInit("pub_template")
	//	fmt.Println(dSource.Insert("pub_template", map[string]interface{}{"code": "PUB_TEMPLATE_CODE", "name": "模版管理", "datasource": "mysql", "tablename": "pub_template"}))
	//自定义的方法
	//	mysqlSource := dSource.Db.(*MySql)
	//	fmt.Println(mysqlSource.InsertBatch("pub_template", []map[string]interface{}{{"username": "dudenglan", "pwd": "abcd1234", "sex": 1}}))

}
func getTableInit(table string) {
	dSource := pool.GetDataSource()
	//接口方法
	fields := dSource.Get("desc " + table).([]map[string]interface{})
	for _, fieldInfo := range fields {
		for key, val := range fieldInfo {
			fmt.Println(key, val)
		}
		fmt.Println()
	}
}

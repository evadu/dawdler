package database

import "fmt"

type DataSource struct {
	name   string
	status int
	Db     DB        //未知的一个数据库，也许是MYSQL 也许是MONGODB 也许是ORACEL
	pool   *DBpoolVo //如果有pool就加进去
}
type DB interface { //任何数据库都逃不过的五种方法 增删改查
	CreateDataSource(params ...interface{}) interface{}
	Insert(params ...interface{}) (err error, data interface{})
	Delete(params ...interface{}) interface{}
	Update(params ...interface{}) interface{}
	Get(args ...interface{}) interface{}
	Release() //释放
	Close()   //关闭
}

/***数据操作的实现***********************************/
func (dsource *DataSource) CreateDataSource(parames ...interface{}) DataSource {
	dsource.Db.CreateDataSource(parames) //去创建真正的DB去了
	return *dsource
}
func (dsource *DataSource) Insert(parames ...interface{}) (err error, data interface{}) {
	fmt.Println("这就是我的过程")
	return dsource.Db.Insert(parames...)
}
func (dsource *DataSource) Delete(parames ...interface{}) interface{} {
	return dsource.Db.Delete(parames...)
}
func (dsource *DataSource) Update(parames ...interface{}) interface{} {
	return dsource.Db.Update(parames...)
}
func (dsource *DataSource) Get(parames ...interface{}) interface{} {
	return dsource.Db.Get(parames...)
}
func (dsource *DataSource) Release() { //数据连接关闭之后
	//通过到它所属的POOL进行数据的删除
	if dsource.pool != nil {
		dsource.pool.busyNum--
		if dsource.pool.busyNum > dsource.pool.maxNum { //有问题，这个规则应该是 当空闲大于多少的时候，去关闭
			dsource.Close()
		}
	} else {
		dsource.Db.Close() //没有就直接关闭去
	}
}
func (dsource *DataSource) Close() {
	dsource.pool.currentNum--
	dsource.Close()
}

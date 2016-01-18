package database

//import "fmt"

type DBpoolVo struct{
	dataSource interface{}
	busyNum int
	currentNum int
	maxNum int
	minNum int
	name string //读的数据池，写的数据池,其它数据库的数据池
	pool map[int]DataSource //连接池
	selfPool map[string]DataSource //自定义的连接池
}

//以免很多数据库是没有人用的
type DBpool interface{
	InitDbPool()
	addDsource()//内部自动添加的数据源,当前的CURRENTNUM 空闲比较小的时候自动添加
	AddDataSource()/**外部添加 可能是别的数据源**/
	GetDataSource() DataSource/**使用后是否自动关闭**/
	GetDataSourceByName(name string) DataSource /**通过name去拿**/
	RemoveByDataSource(dsource DataSource)/**根据数据池去删除**/
	Close()
}
func NewDbPollVo(source DataSource,minNum int,maxNum int) *DBpoolVo{
	poolvo:=&DBpoolVo{dataSource:source,minNum:minNum,maxNum:maxNum}
	poolvo.pool=make(map[int]DataSource)
	poolvo.InitDbPool()
	return poolvo
}
/****数据池的实现**************************************/
//初始化数据池
func (dbpool *DBpoolVo) InitDbPool(){
	//初始化时加载最小的 
	for i:=0;i<dbpool.minNum;i++{
		dbpool.addDsource()
	}
}
func (dbpool *DBpoolVo) addDsource(){
	dSource:=dbpool.dataSource.(DataSource)
	dbpool.pool[dbpool.currentNum]=dSource.CreateDataSource()
	dSource.pool=dbpool
	dbpool.currentNum++ //数据连接池+1
}
func(dbpool *DBpoolVo)AddDataSource(dSource DataSource){
	//需要添加到自动管理吗? dbpool.currentNum++ //数据连接池+1
	dbpool.selfPool[dSource.name]=dSource
	dSource.pool=dbpool
}

func(dbpool *DBpoolVo)GetDataSource() DataSource{
	return dbpool.pool[dbpool.busyNum]
}
func(dbpool *DBpoolVo)GetDataSourceByName(name string) DataSource{
	return dbpool.selfPool[name]
}
func(dbpool *DBpoolVo)RemoveByDataSource(name string) DataSource{
	return dbpool.selfPool[name]
}
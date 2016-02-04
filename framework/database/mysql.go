package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type MySql struct {
	db     *sql.DB
	DBlink string
}

//创建数据源
func (this *MySql) CreateDataSource(params ...interface{}) interface{} {
	db, err := sql.Open("mysql", this.DBlink)
	if err != nil {
		db.Close()
		return err
	}
	//	db.SetMaxIdleConns(p.idleSize)
	//	db.SetMaxOpenConns(p.maxSize)
	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}
	this.db = db
	return this
}

//插入
func (this *MySql) Insert(params ...interface{}) (err error, data interface{}) {
	if len(params) < 2 {
		err = errors.New("params must be have ,tablenam,datas")
		return
	}
	table, ok := params[0].(string)                  //表
	values, ok := params[1].(map[string]interface{}) //MAP
	if ok == false {
		err = errors.New("params[1] type must be map[string]interface{}")
		return
	}
	values = this.validate(table, values) //得到验证数据
	if len(values) == 0 {
		err = errors.New("params[1] map data error")
		return
	}
	//解析数据
	buf := make([]byte, 0)
	buf = append(buf, "INSERT INTO "...)
	buf = append(buf, table...)
	buf = append(buf, " SET "...)
	for key, val := range values {
		buf = append(buf, " "...)
		buf = append(buf, key...)
		buf = append(buf, "='"...)
		buf = append(buf, fmt.Sprint(val)...)
		buf = append(buf, "',"...)
	}
	insertSql := string(buf[:len(buf)-1])
	//执行SQL语句
	res, err := this.db.Exec(insertSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	Id, _ := res.LastInsertId()
	values["id"] = Id
	affectRow, _ := res.RowsAffected()
	values["_AffectRow"] = affectRow
	data = values
	return
}

//批量插入
func (this *MySql) InsertBatch(params ...interface{}) (err error, data interface{}, rowAffect int64) {
	if len(params) < 2 {
		err = errors.New("params must be have ,tablenam,datas")
		return
	}
	table := params[0].(string)                        //表
	values, ok := params[1].([]map[string]interface{}) //数组
	if ok == false {
		err = errors.New("params[1] type must be []map[string]interface{}")
		return
	}
	for key, value := range values {
		values[key] = this.validate(table, value)
		if len(values[key]) == 0 {
			err = errors.New(table + "params[1] map data error" + fmt.Sprint(value))
			return
		}
	}
	if len(values) == 0 {
		err = errors.New("params[1] map data error")
		return
	}
	//解析数据
	buf := make([]byte, 0)
	buf = append(buf, "INSERT INTO "...)
	buf = append(buf, table...)
	buf = append(buf, "("...)

	keys := make([]string, 0)
	for key, _ := range values[0] {
		buf = append(buf, key+","...)
		keys = append(keys, key)
	}
	buf = buf[0 : len(buf)-1]
	buf = append(buf, ") "...)

	buf = append(buf, " values "...)
	for _, value := range values {
		buf = append(buf, " ("...)
		for index := 0; index < len(keys); index++ {
			buf = append(buf, fmt.Sprint("'", value[keys[index]], "',")...)
		}
		buf = buf[:len(buf)-1]
		buf = append(buf, "),"...)
	}
	insertSql := string(buf[:len(buf)-1])
	//执行SQL语句
	res, err := this.db.Exec(insertSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	rowAffect, err = res.RowsAffected() //插入行数
	data = values
	return
}
func (this *MySql) Update(params ...interface{}) interface{} {
	return 0
}
func (this *MySql) Delete(params ...interface{}) interface{} {
	return 0
}
func (this *MySql) Get(params ...interface{}) interface{} {
	sql := params[0].(string)
	rows, err := this.db.Query(sql)
	if err != nil {
		fmt.Println("出错啦")
		return 0
	}

	var records []map[string]interface{}
	cols, _ := rows.Columns()
	sanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
	for i, _ := range values {
		sanArgs[i] = &values[i]
	}
	for rows.Next() {
		record := make(map[string]interface{})
		rows.Scan(sanArgs...)
		for i, col := range values {
			if col != nil {
				record[strings.ToLower(cols[i])] = string(col.([]byte))
			}
		}
		if records == nil {
			records = make([]map[string]interface{}, 0)
		}
		records = append(records, record)

	}
	return records
}

func (this *MySql) Release() {
}

func (this *MySql) Close() {
	fmt.Printf("我去")
}

func (this *MySql) validate(table string, data map[string]interface{}) (ret_data map[string]interface{}) {
	sql := "SELECT COLUMN_NAME name  FROM information_schema.COLUMNS WHERE table_name = '" + table + "'"
	rows := this.Get(sql)
	column := rows.([]map[string]interface{})
	cols := []string{}
	for _, val := range column {
		cols = append(cols, val["name"].(string))
	}
	ret_data = map[string]interface{}{}
	for _, col := range cols {
		if data[col] != nil {
			ret_data[col] = data[col]
		}
	}
	return
}

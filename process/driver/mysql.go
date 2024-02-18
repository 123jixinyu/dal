package driver

import (
	"dal/process/util"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type MySQL struct {
	BaseDriver
}

// 获取与插件匹配的DSN
func (c *MySQL) getDsn() string {

	var str string
	if c.ConnConfig.User != "" {
		str += c.ConnConfig.User

		if c.ConnConfig.Password != "" {
			str += ":" + c.ConnConfig.Password
		}
		str += "@"
	}
	str += "tcp"

	str += fmt.Sprintf("(%s:%d)", c.ConnConfig.LocPorts[0].Loc, c.ConnConfig.LocPorts[0].Port)

	str += "/" + c.ConnConfig.Database

	return str
}

// 执行SQL
func (c *MySQL) Exec(query string) ([]map[string]interface{}, error) {

	db, err := sql.Open("mysql", c.getDsn())
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	columnTypeMap := make(map[string]reflect.Type, 0)

	for _, v := range columnTypes {
		columnTypeMap[v.Name()] = v.ScanType()
	}

	list := make([]map[string]interface{}, 0)

	for rows.Next() {
		var item = map[string]interface{}{}
		cache := make([]interface{}, len(columns))

		for index := range cache {
			var a interface{}
			cache[index] = &a
		}

		if err := rows.Scan(cache...); err != nil {
			return nil, err
		}

		for i, v := range cache {
			a := *v.(*interface{})
			converted, err := util.Convert(a, columnTypeMap[columns[i]])
			if err != nil {
				return nil, err
			}
			item[columns[i]] = converted
		}
		list = append(list, item)
	}
	return list, nil
}

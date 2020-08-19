package utils

import (
	"database/sql"
	"log"
	"sort"
	"strings"
	"time"
)

var DataSourceName = ""

func InsertData(datas []map[string]string) (sql.Result, error) {
	db, err := sql.Open("mysql", DataSourceName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	maxColumn := 0
	allKeys := make(map[string]bool)
	for _, data := range datas {
		for k := range data {
			if strings.Contains(k, ":") {
				k = strings.Replace(k, ":", "", -1)
			}
			if !allKeys[k] {
				allKeys[k] = true
			}
		}
	}
	keys := make([]string, 0, maxColumn)
	for k := range allKeys {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	maxColumn = len(allKeys)
	tableName := "logs_" + time.Now().Format("0102_150405")
	createTable := "create table " + tableName + " ("
	for _, k := range keys {
		createTable += k + ` varchar(1000) default '' not null,`
	}
	createTable = createTable[:len(createTable)-1] + ")"
	es, err := db.Exec(createTable)
	if err != nil {
		return es, err
	}
	insertSql := " insert into " + tableName + " ( "
	for _, k := range keys {
		insertSql += k + `,`
	}
	insertSql = insertSql[:len(insertSql)-1] + ") values  "
	vs := "(" + strings.Repeat("?,", maxColumn)
	vs = vs[:len(vs)-1] + "),"
	insertSql += strings.Repeat(vs, len(datas))
	insertSql = insertSql[:len(insertSql)-1] + ";"

	values := make([]interface{}, 0)
	for _, data := range datas {
		for _, key := range keys {
			values = append(values, data[key])
		}
	}
	return db.Exec(insertSql, values...)
}

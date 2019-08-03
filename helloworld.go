package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var corpid string
	var charge int32
	var discharge int32

	db, err := sql.Open("mysql", "admin:Neoplenty1226@tcp(mysql01.ckb1owlkbaxs.ap-northeast-1.rds.amazonaws.com:3306)/mysql")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("OK")
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM charge")
	rows.Scan(&corpid, &charge, &discharge)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(&corpid)
		fmt.Println(&charge)
		fmt.Println(&discharge)
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
}

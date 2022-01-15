// Package xsql extends for the sql package.
//power by <Joshua Conero>(2020)
package xsql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Query struct {
	DB       *sql.DB
	RowCount int //the count query result lines.
}

func (c *Query) Select(query string, args ...interface{}) ([]map[string]interface{}, error) {
	db := c.DB
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return c.RowsDick(rows)
}

// RowsDick reference link: https://blog.csdn.net/weimingjue/article/details/91042649
//get the rows dick
//database map to golang type:
//		INT    							int64
//		SMALLINT,TINYINT    			int
//		DECIMAL    						float64
//		DATETIME,DATE    				time.Time; <!temp make to string>
//		VARCHAR,CHAR,TEXT				string
func (c *Query) RowsDick(rows *sql.Rows) ([]map[string]interface{}, error) {
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	cTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	receive := make([]interface{}, len(columns))
	for index, _ := range receive {
		cType := cTypes[index]

		dataTypeStr := strings.ToUpper(cType.DatabaseTypeName())
		switch dataTypeStr {
		//@todo fail to map
		//case "DATETIME", "DATE":
		//	var a interface{}
		//	receive[index] = &a
		case "INT", "SMALLINT", "TINYINT", "BIT", "BOOL", "BOOLEAN", "MEDIUMINT", "INTEGER",
			"BIGINT", "YEAR":
			var a sql.NullInt64
			receive[index] = &a
		case "VARCHAR", "CHAR", "TEXT", "BINARY", "VARBINARY", "TINYBLOB", "TINYTEXT", "BLOB",
			"MEDIUMBLOB", "MEDIUMTEXT", "LONGBLOB", "LONGTEXT":
			var a sql.NullString
			receive[index] = &a
		case "DECIMAL", "FLOAT", "DOUBLE", "NUMERIC":
			var a sql.NullFloat64
			receive[index] = &a
		default:
			var a interface{}
			receive[index] = &a
		}
	}

	var dataList []map[string]interface{}
	var counter = 0
	for rows.Next() {
		if err := rows.Scan(receive...); err != nil {
			log.Fatal(err)
		}

		item := make(map[string]interface{})
		//get the true value
		for index, v := range receive {
			cType := cTypes[index]
			col := columns[index]

			dataTypeStr := strings.ToUpper(cType.DatabaseTypeName())
			switch dataTypeStr {
			case "INT", "SMALLINT", "TINYINT", "BIT", "BOOL", "BOOLEAN", "MEDIUMINT", "INTEGER",
				"BIGINT", "YEAR":
				var anyVal = *v.(*sql.NullInt64)
				if anyVal.Valid {
					item[col] = anyVal.Int64
				} else {
					item[col] = nil
				}
			case "DATETIME", "DATE", "TIME", "TIMESTAMP":
				tmpValue := *v.(*interface{})
				var timeStr string
				if tmpValue != nil {
					timeStr = fmt.Sprintf("%s", tmpValue)
				}
				item[col] = timeStr
			case "VARCHAR", "CHAR", "TEXT", "BINARY", "VARBINARY", "TINYBLOB", "TINYTEXT", "BLOB",
				"MEDIUMBLOB", "MEDIUMTEXT", "LONGBLOB", "LONGTEXT":
				var anyVal = *v.(*sql.NullString)
				if anyVal.Valid {
					item[col] = anyVal.String
				} else {
					item[col] = ""
				}
			case "DECIMAL", "FLOAT", "DOUBLE", "NUMERIC":
				var anyVal = *v.(*sql.NullFloat64)
				if anyVal.Valid {
					item[col] = anyVal.Float64
				} else {
					item[col] = nil
				}
			default:
				item[col] = *v.(*interface{})
			}

		}

		dataList = append(dataList, item)
		counter += 1
	}

	c.RowCount = counter
	return dataList, nil
}

func NewQuery(db *sql.DB) *Query {
	query := &Query{
		DB: db,
	}
	return query
}

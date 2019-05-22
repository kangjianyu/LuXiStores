package dao

import (
	"database/sql"
	log "github.com/jeanphorn/log4go"
)

func insert(sqlStr string, args ...interface{}) (lastID int64, err error) {
	stmtIns, err := mysqlClient.Prepare(sqlStr)
	if err != nil {
		log.Error("insert failed! sql: %s, args: %v, error: %v", sqlStr, args, err)
		return
	}
	defer stmtIns.Close()
	result, err := stmtIns.Exec(args...)
	if err != nil {
		log.Error("insert failed! sql: %s, args: %v, error: %v", sqlStr, args, err)
		return
	}

	log.Info("insert success sql: %s, args: %v", sqlStr, args)
	return result.LastInsertId()
}
func FetchRow(sqlStr string, args ...interface{}) (ret map[string]string, err error) {
	stmtOut, err := mysqlClient.Prepare(sqlStr)
	if err != nil {
		log.Error("fetchRow failed! sql: %s, args: %v, error: %v", sqlStr, args, err)
		return
	}
	defer stmtOut.Close()
	row, err := stmtOut.Query(args...)
	if err != nil {
		log.Error("fetchRow failed! sql: %s, args: %v, error: %v", sqlStr, args, err)
		return
	}
	columns, err := row.Columns()
	if err != nil {
		log.Error("fetchRow failed! sql: %s, args: %v, error: %v", sqlStr, args, err)
		return
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	ret = make(map[string]string, len(scanArgs))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for row.Next() {
		err = row.Scan(scanArgs...)
		if err != nil {
			log.Error("fetchRow failed! sql: %s, args: %v, error: %v", sqlStr, args, err)
			return
		}
		var value string

		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			ret[columns[i]] = value
		}
		break //get the first row only
	}
	return
}

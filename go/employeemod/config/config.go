package config

import (
	"database/sql"
	"fmt"
)
 
func Connect_sql() *sql.DB {
    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "SQL2019$"
    dbName := "employees"
 
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        fmt.Println(err.Error())
    }
    return db
}
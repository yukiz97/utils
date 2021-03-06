package dbcon

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetMySQLOpenConnectString return a connect string use for sql.Open()
func GetMySQLOpenConnectString(host string, userName string, password string, dbName string) string {
	return userName + ":" + password + "@tcp(" + host + ")/" + dbName+"?parseTime=true"
}

//InitDBMySQL return database connect depend on input string generate from GetMySQLOpenConnectString() function in this package
func InitDBMySQL(strConnect string) *sql.DB {
	db, err := sql.Open("mysql", strConnect)

	if err != nil {
		panic(err.Error())
	}

	return db
}

/*

package main

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/microsoft/go-mssqldb"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	query := url.Values{}
	query.Add("app name", "MyAppName")
	//u := "sqlserver://sa:vls021130@LAPTOP-AL5AFKOJ/SQL2019?app+name=MyAppName"
	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword("sa", "vls021130"),
		Host:   fmt.Sprintf("LAPTOP-AL5AFKOJ"),
		// Path:     "SQL2019", // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	// db, err := sql.Open("sqlserver", u.String())
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
	defer db.Close()

	//exec(db, "CREATE DATABASE TESTGO")
	exec(db, "USE TESTGO IF NOT EXISTS (select * from dbo.sysobjects where id = object_id(N'[dbo].[USERS]') and OBJECTPROPERTY(id, N'IsUserTable') = 1) BEGIN CREATE TABLE USERS (id INT IDENTITY(1,1) PRIMARY KEY, name VARCHAR(80)) END")
	// exec(db, "USE TESTGO")
	// exec(db, "IF EXISTS (select * from dbo.sysobjects where id = object_id(N'[dbo].[USERS]') and OBJECTPROPERTY(id, N'IsUserTable') = 1) BEGIN DROP TABLE USERS END")
	// exec(db, "IF NOT EXISTS (select * from dbo.sysobjects where id = object_id(N'[dbo].[USERS]') and OBJECTPROPERTY(id, N'IsUserTable') = 1) BEGIN CREATE TABLE USERS (id INT IDENTITY(1,1) PRIMARY KEY, name VARCHAR(80)) END")
	// exec(db, `CREATE TABLE USERS (
	//     id INT IDENTITY(1,1) PRIMARY KEY,
	//     name VARCHAR(80)
	// )`)
}

// func main() {
// 	db, err := sql.Open("mysql", "root:123456@/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	exec(db, "CREATE DATABASE IF NOT EXISTS TESTGO")
// 	exec(db, "USE TESTGO")
// 	exec(db, "DROP TABLE IF NOT EXISTS USERS")
// 	exec(db, `CRETE TABLE USERS (
//         id INTEGER AUTO_INCREMENT,
//         name VARCGAR(80),
//         PRIMARY KEY (id)
//     )`)

// }
*/

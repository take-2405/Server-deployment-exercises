package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

// Driver名
const driverName = "mysql"

// Conn 各repositoryで利用するDB接続(Connection)情報
var Conn *sql.DB
var datasource string

func init() {
	/* ===== データベースへ接続する. ===== */
	if os.Getenv("DATABASE_URL") != "" {
		// Heroku用
		datasource = "bd415c01f8fed4:4ccfd2ff@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_33fa03f071bb0e3?parseTime=true"
		var err error
		Conn, err = sql.Open(driverName, datasource)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(datasource)
	} else {
		// ローカル用
		// ユーザ
		user := os.Getenv("MYSQL_USER")
		// パスワード
		password := os.Getenv("MYSQL_PASSWORD")
		// 接続先ホスト
		host := os.Getenv("MYSQL_HOST")
		// 接続先ポート
		port := os.Getenv("MYSQL_PORT")
		// 接続先データベース
		database := os.Getenv("MYSQL_DATABASE")

		// 接続情報は以下のように指定する.
		// user:password@tcp(host:port)/database
		var err error
		Conn, err = sql.Open(driverName,
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
		if err != nil {
			log.Fatal(err)
		}
	}
}

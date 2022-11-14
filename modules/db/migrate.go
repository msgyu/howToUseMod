package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // これがないと20行目でエラーになる
)

var Db *sql.DB

var err error

// マイグレーションを実行
func Migrate() {
	Db, err = sql.Open(Config.SQLDriver, Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	migrateUser(Db)
	migrateTodo(Db)
}

// ユーザーテーブルを作成
func migrateUser(Db *sql.DB) {
	tableNameUser := "users"
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`, tableNameUser)

	_, err := Db.Exec(cmdU)

	if err != nil {
		log.Fatalln(err)
	}
}

// todoテーブルを作成
func migrateTodo(Db *sql.DB) {
	tableNameTodo := "todos"

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME
	)`, tableNameTodo)

	_, err := Db.Exec(cmdT)

	if err != nil {
		log.Fatalln(err)
	}
}

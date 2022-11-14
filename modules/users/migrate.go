package users

import (
	"database/sql"
	"fmt"
	"log"
)

func Migrate(Db *sql.DB) {
	tableName := "users"
	// SQLコマンド作成
	commond := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME
	)`, tableName)

	// SQLコマンド実行
	_, err := Db.Exec(commond)

	if err != nil {
		log.Fatalln(err)
	}
}

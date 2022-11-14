package db

import "fmt"

func init() {
	SetConfig()
	Migrate()
}

func FmtExec() {
	fmt.Println("Dbが呼び出されました")
}

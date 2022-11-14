package main

import (
	"fmt"
	"howToUseMod/modules/db"
	"howToUseMod/modules/routers"
	"howToUseMod/subpkg1"
	"howToUseMod/subpkg2"
)

func main() {
	subpkg1.ExecFmt()
	subpkg2.ExecFmt()
	fmt.Println(db.Config)
	routers.Main()
}

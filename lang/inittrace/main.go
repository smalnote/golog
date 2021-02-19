package main

import (
	"log"

	_ "github.com/smalnote/golog/lang/inittrace/bar"
	_ "github.com/smalnote/golog/lang/inittrace/foo"
)

func main() {
	log.Println("set GODEBUG=inittrace=1 to enable init trace")
	log.Println("package main import foo, bar anymously")
}

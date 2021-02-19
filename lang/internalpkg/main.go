package main

import (
	"log"

	"github.com/smalnote/golog/lang/internalpkg/foo"
	// "github.com/smalnote/golog/lang/internalpkg/foo/internal/bar" not allowed
)

func main() {
	log.Println(foo.NAME)
}

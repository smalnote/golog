package main

import (
	"context"
	"log"

	"github.com/smalnote/golog/lang/contextcfg/redis"
)

func main() {

	cfg := &redis.Config{
		Hostname: "9.134.18.91",
		Port:     6379,
		Password: "unknown",
	}

	ctx := cfg.Context(context.Background())

	c := redis.FromContext(ctx)
	log.Printf("%+v", c)

	c2 := redis.FromContext(context.Background())
	log.Printf("%+v", c2)
}

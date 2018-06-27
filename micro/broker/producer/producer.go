package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
)

var (
	topic = "go.micro.topic.foo"
)

func pub(exit <-chan os.Signal) {
	tick := time.NewTicker(time.Second)
	i := 0
	for {
		select {
		case <-tick.C:
			msg := &broker.Message{
				Header: map[string]string{
					"id": fmt.Sprintf("%d", i),
				},
				Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
			}
			if err := broker.Publish(topic, msg); err != nil {
				log.Printf("[pub] failed: %v", err)
			} else {
				fmt.Println("[pub] pubbed message:", string(msg.Body))
			}
			i++
		case <-exit:
			err := broker.Disconnect()
			if err != nil {
				log.Fatal("disconnect error: ", err)
			}
			log.Println("disconnected ")
			return
		}
	}
}

func main() {
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	pub(ch)
}

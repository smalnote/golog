package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
	// To enable rabbitmq plugin uncomment
	//_ "github.com/micro/go-plugins/broker/rabbitmq"
)

var (
	topic = "go.micro.topic.foo"
)

// Example of a shared subscription which receives a subset of messages
func sharedSub() {
	_, err := broker.Subscribe(topic, func(p broker.Publication) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
}

// Example of a subscription which receives all the messages
func sub(exit <-chan os.Signal) {
	_, err := broker.Subscribe(topic, func(p broker.Publication) error {
		fmt.Println("[sub] received message:", string(p.Message().Body), "header", p.Message().Header)
		return nil
	}, broker.Queue("consumer"))
	if err != nil {
		fmt.Println(err)
	}
	<-exit
	err = broker.Disconnect()
	if err != nil {
		log.Fatal("disconnect error: ", err)
	}
	log.Println("disconnected ")
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
	sub(ch)
}

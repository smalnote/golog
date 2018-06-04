package lang

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
)

func startDaemon() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		listener, err := net.Listen("tcp", "localhost:10086")
		if err != nil {
			log.Fatalf("cannot listen: %v ", err)
		}
		defer listener.Close()
		ch <- struct{}{}

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("cannot accept: %v\n", err)
				continue
			}
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return ch
}

func init() {
	ch := startDaemon()
	<-ch
}

func BenchmarkNetReq(b *testing.B) {
	for count := 0; count < b.N; count++ {
		conn, err := net.Dial("tcp", "localhost:10086")
		if err != nil {
			b.Fatalf("dial failed: %v\n", err)
			continue
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("read from conn failed: %v\n", err)
		}
		conn.Close()
	}
}

func BenchmarkPoolNetReq(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			conn, err := net.Dial("tcp", "localhost:10086")
			if err != nil {
				log.Fatalln("dail failed: ", err)
				return nil
			}
			return conn
		},
	}

	for count := 0; count < b.N; count++ {
		conn, ok := pool.Get().(net.Conn)
		if !ok || conn == nil {
			b.Fatalf("get conn from pool failed ")
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("read from conn failed: %v\n", err)
		}
		pool.Put(conn)
	}
}

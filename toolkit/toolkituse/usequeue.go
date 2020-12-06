package main

import (
	"log"
	"time"

	t "github.com/smalnote/golog/toolkit"
)

func main() {
	queue := t.NewQueue()
	for i := 0; i < 10; i++ {
		jobID := i
		err := queue.AddJob(func() {
			for i := 0; i <= jobID+1; i++ {
				<-time.After(100 * time.Millisecond)
				log.Printf("job #%d: tick #%d", jobID, i+1)
			}
		})
		if err != nil {
			log.Printf("add job #%d failed", jobID)
		}
	}
	queue.Start()
	<-time.After(time.Second)
	queue.Stop()
}

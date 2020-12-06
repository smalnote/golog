package toolkit

import (
	"fmt"
	"log"
	"time"
)

type job struct {
	handler func()
	timeout <-chan time.Time
}

func (job *job) run(done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	select {
	case <-job.timeout:
		log.Println("job is timeout")
	default:
		job.handler()
	}
}

// Queue is a job queue for runing job concurrent
type Queue interface {
	AddJob(func()) error
	Start()
	Stop()
}

type queue struct {
	maxConcurrentJobs int
	maxJobs           int
	execQueue         chan *job
	stop              chan struct{}
	token             chan struct{}
}

// NewQueue create a default job queue
func NewQueue() Queue {
	queue := queue{
		maxConcurrentJobs: 3,
		maxJobs:           10,
		execQueue:         make(chan *job, 10-3),
		token:             make(chan struct{}, 3),
		stop:              make(chan struct{}),
	}
	return &queue
}

// AddJob put the job to queue and wait for runing
func (q *queue) AddJob(handler func()) error {
	job := job{
		handler: handler,
		timeout: time.After(5 * time.Second),
	}
	select {
	case q.execQueue <- &job:
		return nil
	default:
		return fmt.Errorf("too many concurrent and queued jobs")
	}
}

func (q *queue) Start() {
	for i := 0; i < q.maxConcurrentJobs; i++ {
		q.token <- struct{}{}
	}
	go func() {
		for {
			select {
			case <-q.stop:
				log.Println("the queue stopped")
				break
			case job := <-q.execQueue:
				<-q.token
				go job.run(q.token)
			}
		}
	}()
}

func (q *queue) Stop() {
	q.stop <- struct{}{}
}

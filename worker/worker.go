package worker

import (
	"fmt"
	"sync"
	"time"
)

// Job contains ticker, job duration and status
type Job struct {
	Ticker   *time.Ticker
	Duration time.Duration
	Status   int // 0 - not started, 1 - in progress, 2 - waiting
}

// Worker schedules jobs and works on them
type Worker struct {
	Wg   sync.WaitGroup
	Jobs []*Job
}

// NewJob initializes job
func NewJob(p, d time.Duration) *Job {
	return &Job{Ticker: time.NewTicker(p), Duration: d, Status: 0}
}

// NewWorker initializes worker
func NewWorker() *Worker {
	return &Worker{}
}

// Do emulates work
func (j *Job) Do() {
	time.Sleep(j.Duration)
}

// Add adds jobs to worker
func (w *Worker) Add(p, d time.Duration) {
	w.Jobs = append(w.Jobs, NewJob(p, d))
	w.Wg.Add(1)
	fmt.Printf("Added job %v to worker: period = %v, duration = %v\n", len(w.Jobs)-1, p, d)
}

// Work starts work on a job until the ticker is stopped
func (w *Worker) Work(i int) {
	if w.Jobs[i].Status == 0 {
		for ; true; <-w.Jobs[i].Ticker.C {
			w.Jobs[i].Status = 1
			fmt.Println("Started working on job", i)

			w.Jobs[i].Do()

			if w.Jobs[i].Status == 1 {
				fmt.Printf("Job %v is waiting\n", i)
				w.Jobs[i].Status = 2
			} else {
				fmt.Println("Finished working on job", i)
				w.Wg.Done()
			}
		}
	}
}

// Start starts work on all jobs in goroutines
func (w *Worker) Start() {
	for i := 0; i < len(w.Jobs); i++ {
		go w.Work(i)
	}
}

// Stop stops job i from starting again
func (w *Worker) Stop(i int) {
	if w.Jobs[i].Status != 0 {
		w.Jobs[i].Ticker.Stop()

		if w.Jobs[i].Status == 1 {
			fmt.Println("Stopped repeating job", i)
			w.Jobs[i].Status = 2
		} else {
			fmt.Println("Stopped repeating and finished job", i)
			w.Wg.Done()
		}
	}
}

// Status prints status of worker's jobs
func (w *Worker) Status() {
	for i := 0; i < len(w.Jobs); i++ {
		fmt.Printf("Job %v status: %v\n", i, w.Jobs[i].Status)
	}
}

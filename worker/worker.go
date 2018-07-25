package worker

import (
	"fmt"
	"sync"
	"time"
)

// Job contains job duration and status
type Job struct {
	Duration time.Duration
	Status   int // 0 - not done, 1 - in progress, 2 - done
}

// Worker schedules jobs and works on them
type Worker struct {
	Wg    sync.WaitGroup
	Jobs  []*Job
	Timer *time.Timer
}

// NewJob initializes job
func NewJob(d time.Duration) *Job {
	return &Job{Duration: d * time.Second, Status: 0}
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
func (w *Worker) Add(delta int, d time.Duration) {
	for i := 0; i < delta; i++ {
		w.Jobs = append(w.Jobs, NewJob(d))
		fmt.Println("Added job to worker")
	}

	w.Wg.Add(delta)
}

// Work starts work on a job
func (w *Worker) Work(i int) {
	if w.Jobs[i].Status == 0 {
		defer w.Wg.Done()

		w.Jobs[i].Status = 1
		fmt.Println("Started working on job", i)

		w.Jobs[i].Do()

		w.Jobs[i].Status = 2
		fmt.Println("Finished working on job", i)
	}
}

// Start starts work on all jobs on schedule
func (w *Worker) Start(d time.Duration) {
	w.Timer = time.NewTimer(d * time.Second)

	for i := 0; i < len(w.Jobs); i++ {
		go w.Work(i)

		<-w.Timer.C
		w.Timer.Reset(d * time.Second)
	}

	w.Wg.Wait()
}

// Status prints status of worker's jobs
func (w *Worker) Status() {
	for i := 0; i < len(w.Jobs); i++ {
		fmt.Printf("Job %v: duration: %v, status: %v\n", i, w.Jobs[i].Duration.Seconds(), w.Jobs[i].Status)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	id     int
	status int // 0 - not done, 1 - in progress, 2 - done
}

type worker struct {
	wg   sync.WaitGroup
	jobs []*job
	t    *time.Timer
}

func (w *worker) add() {
	w.jobs = append(w.jobs, &job{len(w.jobs), 0})
	w.wg.Add(1)
	fmt.Println("Added job to worker")
}

func (w *worker) work(i int) {
	if w.jobs[i].status == 0 {
		defer w.wg.Done()

		w.jobs[i].status = 1
		fmt.Println("Started working on job", w.jobs[i].id)

		time.Sleep(5 * time.Second)

		w.jobs[i].status = 2
		fmt.Println("Finished working on job", w.jobs[i].id)
	} else {
		fmt.Println("Incorrect job", w.jobs[i].id)
	}
}

func (w *worker) start(interval time.Duration) {
	w.t = time.NewTimer(interval * time.Second)

	for i := 0; i < len(w.jobs); i++ {
		go w.work(i)

		<-w.t.C
		w.t.Reset(interval * time.Second)
	}

	w.wg.Wait()
}

func main() {
	var w worker

	for i := 0; i < 3; i++ {
		w.add()
	}

	w.start(3)
}

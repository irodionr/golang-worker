package main

import (
	"time"

	"./worker"
)

func main() {
	w := worker.NewWorker()

	// Add 2 jobs
	w.Add(10*time.Second, 5*time.Second) // Work intervals: 0-5, 10-15
	w.Add(16*time.Second, 8*time.Second) // Work intervals: 0-8, 16-24

	// Start working on both jobs
	w.Start()

	// Stop repeating job 0 at 12 seconds
	// Job 0 should be finished at 15 seconds
	time.Sleep(12 * time.Second)
	w.Stop(0)

	// Stop repeating job 1 at 26 seconds
	// Job 1 should be already finished at 24 seconds
	time.Sleep(14 * time.Second)
	w.Stop(1)

	w.Wg.Wait()
}

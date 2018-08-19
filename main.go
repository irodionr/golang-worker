package main

import (
	"time"

	"./worker"
)

func main() {
	w := worker.NewWorker()

	w.Add(10*time.Second, 5*time.Second)
	w.Add(16*time.Second, 8*time.Second)
	w.Status()

	w.Start()

	time.Sleep(12 * time.Second)
	w.Stop(0)
	w.Status()

	time.Sleep(12 * time.Second)
	w.Stop(1)
	w.Status()

	w.Wg.Wait()
}

package main

import "./worker"

func main() {
	w := worker.NewWorker()
	w.Add(3, 5)
	w.Status()
	w.Start(3)
	w.Status()
}

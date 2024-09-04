package asyncpatterns

import "sync"

type Worker struct {
	workers int
	jobs    chan func()
}

func NewWorker(workers int) Worker {
	return Worker{workers, make(chan func(), 50)}
}

func (w *Worker) Run() {
	var wg sync.WaitGroup

	wg.Add(w.workers)
	for i := 0; i < w.workers; i++ {
		go func() {
			for job := range w.jobs {
				job()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func (w *Worker) Close() {
	close(w.jobs)
}

func (w *Worker) Add(job func()) {
	w.jobs <- job
}

package asyncpatterns_test

import (
	"testing"

	asyncpatterns "github.com/kauebonfimm/go-algorithms/async_patterns"
)

func TestWorkerAsync(t *testing.T) {
	worker := asyncpatterns.NewWorker(9)
	go worker.Run()

	results := make(chan int, 300)
	jobs := 10000000
	sum := 0
	stop := make(chan struct{})

	go func() {
		i := 1
		for result := range results {
			sum += result
			if jobs <= i {
				close(results)
				stop <- struct{}{}
			}
			i++

		}
	}()

	for i := 1; i <= jobs; i++ {
		index := i
		worker.Add(func() {
			results <- index
		})
	}

	<-stop
	worker.Close()
	t.Log(sum)

}

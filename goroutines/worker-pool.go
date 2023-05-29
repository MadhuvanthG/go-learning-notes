package goroutines

/**
 * Thread-pool equivalent: a factory function to create a fixed number of go routines
 * These go routines will block waiting for a task to be done (passed through channel)
 */

type Task func() string

type TaskExecutor interface {
	Submit(task Task)
	GetResults() <-chan string
}

type workerPool struct {
	tasks   chan Task
	results chan string
}

func NewWorkerPool(numWorkers int, bufferCapacity int) TaskExecutor {
	wp := workerPool{
		tasks:   make(chan Task, bufferCapacity),
		results: make(chan string, bufferCapacity),
	}

	wp.spawnWorkers(numWorkers)

	return wp
}

func (w workerPool) spawnWorkers(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go func() {
			for task := range w.tasks {
				w.results <- task()
			}
		}()
	}
}

func (w workerPool) Submit(task Task) {
	w.tasks <- task
}

func (w workerPool) GetResults() <-chan string {
	return w.results
}

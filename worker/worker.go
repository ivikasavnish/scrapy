package worker

type worker struct{}

type WorkerManager struct {
	HowMany int
	Workers []*worker
}

func NewWorkerManager(howmany int) *WorkerManager {
	return &WorkerManager{
		HowMany: howmany,
	}
}

func (wm *WorkerManager) Start() {
	for i := 0; i < wm.HowMany; i++ {
		wm.Workers = append(wm.Workers, NewWorker())
	}

	for _, w := range wm.Workers {
		w.Start()
	}
}

func NewWorker() *worker {
	return &worker{}
}

// i need to implement a worker pool

func (w *worker) Start() {
	// pull the work from the queue
	// do the work
	// push the result to the queue

}

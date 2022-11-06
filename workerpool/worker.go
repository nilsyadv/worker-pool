package workerpool

type workerPool struct {
	maxWorker   int
	queuedTaskC chan func()
}

func NewWorkerPool(totalWorker int) *workerPool {
	return &workerPool{
		maxWorker:   totalWorker,
		queuedTaskC: make(chan func()),
	}
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for task := range wp.queuedTaskC {
				task()
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

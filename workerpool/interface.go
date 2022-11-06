package workerpool

// T is a type alias to accept any type.
type T = interface{}

// WorkerPool is a contract for Worker Pool implementation
type WorkerPool interface {
	Run()
	AddTask(task func())
}

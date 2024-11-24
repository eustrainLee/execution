package sr

type Receiver[T any] interface {
	SetValue(T)
	SetError(err error)
	SetStoped()
}

type Sender[T any] interface {
	Connect(Receiver[T]) OperationState
}

type Scheduler[T any] interface {
	Schedule() Sender[T]
}

type OperationState interface {
	Start()
}

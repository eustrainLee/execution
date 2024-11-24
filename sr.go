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

type integer interface {
	int | uint | int16 | uint16 | int32 | uint32 | int64 | uint64
}

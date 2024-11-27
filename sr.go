package sr

type SenderTag uint32

const (
	SenderTagNone         SenderTag = 0
	SenderTagMultiSend    SenderTag = 1 << iota // If the tag has set, It can be used multipy in each operation state.
	SenderTagMultiConnect                       // If the tag has set, It can be Connect to multiple receiver.
	// SenderTagStop                               // If the tag has set, It can be stoped by operation state.
)

type Receiver[T any] interface {
	SetValue(T)
	SetError(err error)
	SetStoped()
}

type Sender[T any] interface {
	Tag() SenderTag
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

type receiverOperation int

const (
	receiverOperationNone receiverOperation = iota
	receiverOperationHasValue
	receiverOperationHasError
	receiverOperationStoped
)

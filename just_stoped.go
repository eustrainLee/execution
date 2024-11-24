package sr

type justStopedSender[T any] struct{}

func JustStoped[T any]() Sender[T] {
	return justStopedSender[T]{}
}

func (s justStopedSender[T]) Connect(r Receiver[T]) OperationState {
	return justStopedSenderState[T]{r: r}
}

type justStopedSenderState[T any] struct {
	r Receiver[T]
}

func (state justStopedSenderState[T]) Start() {
	state.r.SetStoped()
}

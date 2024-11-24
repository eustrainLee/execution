package sr

type repeatSender[T any] struct {
	s Sender[T]
}

func Repeat[T any](s Sender[T]) Sender[T] {
	return repeatSender[T]{s: s}
}

func (s repeatSender[T]) Connect(r Receiver[T]) OperationState {
	return repeatSenderState[T]{s: s, r: r}
}

type repeatSenderState[T any] struct {
	s repeatSender[T]
	r Receiver[T]
}

func (state repeatSenderState[T]) Start() {
	for {
		state.s.s.Connect(state.r).Start()
	}
}

package sr

type repeatNSender[T any, I integer] struct {
	n I
	s Sender[T]
}

func RepeatN[T any, I integer](s Sender[T], n I) Sender[T] {
	return repeatNSender[T, I]{s: s, n: n}
}

func (s repeatNSender[T, I]) Connect(r Receiver[T]) OperationState {
	return repeatNSenderState[T, I]{s: s, r: r}
}

type repeatNSenderState[T any, I integer] struct {
	s repeatNSender[T, I]
	r Receiver[T]
}

func (state repeatNSenderState[T, I]) Start() {
	for i := I(0); i < state.s.n; i++ {
		state.s.s.Connect(state.r).Start()
	}
}

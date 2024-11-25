// sender factory: JustValue
package sr

type justValueSender[T any] struct {
	v T
}

func Just[T any](v T) Sender[T] {
	return &justValueSender[T]{v: v}
}

func (s *justValueSender[T]) Connect(r Receiver[T]) OperationState {
	return justValueSenderState[T]{s: s, r: r}
}

type justValueSenderState[T any] struct {
	s *justValueSender[T]
	r Receiver[T]
}

func (state justValueSenderState[T]) Start() {
	state.r.SetValue(state.s.v)
}

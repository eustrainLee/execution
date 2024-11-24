// sender factory: JustValue
package sr

type justValueSender[T any] struct {
	v T
}

func Just[T any](v T) Sender[T] {
	return &justValueSender[T]{v: v}
}

func JustResultOf[T any](f func() T) Sender[T] {
	return &justValueSender[T]{v: f()}
}

func (s *justValueSender[T]) Connect(r Receiver[T]) OperationState {
	return justValueSenderState[T]{}
}

type justValueSenderState[T any] struct {
	s *justValueSender[T]
	r Receiver[T]
}

func (state justValueSenderState[T]) Start() {
	state.r.SetValue(state.s.v)
}

package sr

type justErrorSender[T any] struct {
	err error
}

func JustError[T any](err error) Sender[T] {
	return &justErrorSender[T]{err: err}
}

func JustErrorResultOf[T any](func(err error)) Sender[T] {
	return nil
}

func (s *justErrorSender[T]) Connect(r Receiver[T]) OperationState {
	return justErrorSenderState[T]{}
}

type justErrorSenderState[T any] struct {
	s *justErrorSender[T]
	r Receiver[T]
}

func (state justErrorSenderState[T]) Start() {
	state.r.SetError(state.s.err)
}

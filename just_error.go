package sr

import "context"

type justErrorSender[T any] struct {
	err error
}

func JustError[T any](err error) Sender[T] {
	return justErrorSender[T]{err: err}
}

func (s justErrorSender[T]) Connect(r Receiver[T]) OperationState {
	return justErrorSenderState[T]{err: s.err, r: r}
}

func (s justErrorSender[T]) Tag() SenderTag {
	return SenderTagMultiShot
}

type justErrorSenderState[T any] struct {
	err error
	r   Receiver[T]
}

func (state justErrorSenderState[T]) Start(context.Context) {
	if state.r != nil {
		state.r.SetError(state.err)
	}
}

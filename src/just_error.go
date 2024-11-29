package src

import (
	"context"

	"github.com/eustrainLee/sr"
)

type justErrorSender[T any] struct {
	err error
}

func JustError[T any](err error) Sender[T] {
	return justErrorSender[T]{err: err}
}

func (s justErrorSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return justErrorSenderState[T]{err: s.err, r: r}
}

func (s justErrorSender[T]) Tag() sr.SenderTag {
	return sr.SenderTagMultiShot
}

type justErrorSenderState[T any] struct {
	err error
	r   sr.Receiver[T]
}

func (state justErrorSenderState[T]) Start(context.Context) {
	if state.r != nil {
		state.r.SetError(state.err)
	}
}

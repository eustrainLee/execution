package src

import (
	"context"

	"github.com/eustrainLee/sr"
)

type nonCtxSender[T any] struct {
	s sr.Sender[T]
}

func NonCtx[T any](s sr.Sender[T]) Sender[T] {
	return nonCtxSender[T]{s: s}
}

func (s nonCtxSender[T]) Tag() sr.SenderTag {
	return s.s.Tag()
}

func (s nonCtxSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return nonCtxOperationState[T]{s: s.s, r: r}
}

type nonCtxOperationState[T any] struct {
	s sr.Sender[T]
	r sr.Receiver[T]
}

func (os nonCtxOperationState[T]) Start(context.Context) {
	os.s.Connect(os.r).Start()
}

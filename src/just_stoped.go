package src

import (
	"context"

	"github.com/eustrainLee/sr"
)

type justStopedSender[T any] struct{}

func JustStoped[T any]() Sender[T] {
	return justStopedSender[T]{}
}

func (s justStopedSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return justStopedSenderState[T]{r: r}
}

func (s justStopedSender[T]) Tag() sr.SenderTag {
	return sr.SenderTagMultiShot
}

type justStopedSenderState[T any] struct {
	r sr.Receiver[T]
}

func (state justStopedSenderState[T]) Start(context.Context) {
	state.r.SetStoped()
}

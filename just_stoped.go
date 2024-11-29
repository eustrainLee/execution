package sr

import "context"

type justStopedSender[T any] struct{}

func JustStoped[T any]() Sender[T] {
	return justStopedSender[T]{}
}

func (s justStopedSender[T]) Connect(r Receiver[T]) OperationState {
	return justStopedSenderState[T]{r: r}
}

func (s justStopedSender[T]) Tag() SenderTag {
	return SenderTagMultiShot
}

type justStopedSenderState[T any] struct {
	r Receiver[T]
}

func (state justStopedSenderState[T]) Start(context.Context) {
	state.r.SetStoped()
}

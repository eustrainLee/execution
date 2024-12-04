package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
)

type fromChannelSender[T any] struct {
	ch <-chan T
}

func FromChannel[T any](ch <-chan T) Sender[T] {
	return fromChannelSender[T]{ch: ch}
}

func (s fromChannelSender[T]) Tag() sr.SenderTag {
	return sr.SenderTagMultiShot
}

func (s fromChannelSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return fromChannelOperationState[T]{ch: s.ch, r: r}
}

type fromDisposableChannelSender[T any] struct {
	ch <-chan T
}

func FromDisposableChannel[T any](ch <-chan T) Sender[T] {
	return fromDisposableChannelSender[T]{ch: ch}
}

func (s fromDisposableChannelSender[T]) Tag() sr.SenderTag {
	return sr.SenderTagNone
}

func (s fromDisposableChannelSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return fromChannelOperationState[T]{ch: s.ch, r: r}
}

type fromChannelOperationState[T any] struct {
	ch <-chan T
	r  sr.Receiver[T]
}

func (os fromChannelOperationState[T]) Start(context.Context) {
	v, ok := <-os.ch
	if ok {
		os.r.SetValue(v)
	} else {
		os.r.SetStoped()
	}
}

package sr

import (
	"context"
	"sync"
)

type splitSender[T any] struct {
	s    Sender[T]
	once sync.Once
	buff TrivalReceiver[T]
}

func Split[T any](s Sender[T]) Sender[T] {
	if s.Tag()&SenderTagMultiShot == 0 {
		return &splitSender[T]{s: s}
	}
	return s
}

func (s *splitSender[T]) Connect(r Receiver[T]) OperationState {
	return splitSenderState[T]{s: s, r: r}
}

func (s *splitSender[T]) Tag() SenderTag {
	return s.s.Tag() | SenderTagMultiShot
}

type splitSenderState[T any] struct {
	s *splitSender[T]
	r Receiver[T]
}

func (state splitSenderState[T]) Start(ctx context.Context) {
	state.s.once.Do(func() { state.s.s.Connect(&state.s.buff).Start(ctx) })
	state.s.buff.Forward(state.r)
}

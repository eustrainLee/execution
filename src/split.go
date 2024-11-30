package src

import (
	"context"
	"sync"

	"github.com/eustrainLee/execution/sr"
)

type splitSender[T any] struct {
	s    Sender[T]
	once sync.Once
	buff sr.TrivalReceiver[T]
}

func Split[T any](s Sender[T]) Sender[T] {
	if s.Tag()&sr.SenderTagMultiShot == 0 {
		return &splitSender[T]{s: s}
	}
	return s
}

func (s *splitSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return splitSenderState[T]{s: s, r: r}
}

func (s *splitSender[T]) Tag() sr.SenderTag {
	return s.s.Tag() | sr.SenderTagMultiShot
}

type splitSenderState[T any] struct {
	s *splitSender[T]
	r sr.Receiver[T]
}

func (state splitSenderState[T]) Start(ctx context.Context) {
	state.s.once.Do(func() { state.s.s.Connect(&state.s.buff).Start(ctx) })
	state.s.buff.Forward(state.r)
}

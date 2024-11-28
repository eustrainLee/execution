package sr

import "sync"

type repeatSender[T any] struct {
	s    Sender[T]
	once sync.Once
	buff TrivalReceiver[T]
}

func Repeat[T any](s Sender[T]) Sender[T] {
	return &repeatSender[T]{s: s}
}

func (s *repeatSender[T]) Connect(r Receiver[T]) OperationState {
	return repeatSenderState[T]{s: s, r: r}
}

func (s *repeatSender[T]) Tag() SenderTag {
	return s.s.Tag() | SenderTagMultiShot
}

type repeatSenderState[T any] struct {
	s *repeatSender[T]
	r Receiver[T]
}

func (state repeatSenderState[T]) Start() {
	state.s.once.Do(state.s.s.Connect(&state.s.buff).Start)
	state.s.buff.Forward(state.r)
}

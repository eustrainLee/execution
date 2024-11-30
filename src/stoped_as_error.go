package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
)

type stopedAsErrorSender[T any] struct {
	s   Sender[T]
	err error
}

func StopedAsError[T any](s Sender[T], err error) Sender[T] {
	return stopedAsErrorSender[T]{s: s}
}

func (s stopedAsErrorSender[T]) Tag() sr.SenderTag {
	return s.s.Tag()
}

func (s stopedAsErrorSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return stopedAsErrorOperationState[T]{s: s.s, r: r, err: s.err}
}

type stopedAsErrorOperationState[T any] struct {
	s   Sender[T]
	r   sr.Receiver[T]
	err error
}

func (os stopedAsErrorOperationState[T]) Start(ctx context.Context) {
	os.s.Connect(stopedAsErrorReceiver[T]{r: os.r, err: os.err}).Start(ctx)
}

type stopedAsErrorReceiver[T any] struct {
	err error
	r   sr.Receiver[T]
}

func (r stopedAsErrorReceiver[T]) SetValue(v T) {
	r.r.SetValue(v)
}

func (r stopedAsErrorReceiver[T]) SetError(err error) {
	r.r.SetError(err)
}

func (r stopedAsErrorReceiver[T]) SetStoped() {
	r.r.SetError(r.err)
}

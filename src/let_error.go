// sender factory: LetError
package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
)

type letErrorSender[T any] struct {
	s Sender[T]
	f func(context.Context, error, sr.Receiver[T])
}

func LetError[T any](s Sender[T], f func(context.Context, error, sr.Receiver[T])) Sender[T] {
	return letErrorSender[T]{s: s, f: f}
}

func (s letErrorSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return letErrorSenderState[T]{s: s.s, f: s.f, r: r}
}

func (s letErrorSender[T]) Tag() sr.SenderTag {
	return s.s.Tag()
}

type letErrorSenderState[T any] struct {
	s Sender[T]
	f func(context.Context, error, sr.Receiver[T])
	r sr.Receiver[T]
}

func (state letErrorSenderState[T]) Start(ctx context.Context) {
	state.s.Connect(letErrorReceiver[T]{ctx: ctx, f: state.f, r: state.r}).Start(ctx)
}

type letErrorReceiver[T any] struct {
	ctx context.Context
	f   func(context.Context, error, sr.Receiver[T])
	r   sr.Receiver[T]
}

func (ler letErrorReceiver[T]) SetValue(v T) {
	ler.r.SetValue(v)
}

func (ler letErrorReceiver[T]) SetError(err error) {
	ler.f(ler.ctx, err, ler.r)
}

func (ler letErrorReceiver[T]) SetStoped() {
	ler.r.SetStoped()
}

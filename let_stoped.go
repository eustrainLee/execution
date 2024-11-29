// sender factory: LetStoped
package sr

import "context"

type letStopedSender[T any] struct {
	s Sender[T]
	f func(context.Context, Receiver[T])
}

func LetStoped[T any](s Sender[T], f func(context.Context, Receiver[T])) Sender[T] {
	return letStopedSender[T]{s: s, f: f}
}

func (s letStopedSender[T]) Connect(r Receiver[T]) OperationState {
	return letStopedSenderState[T]{s: s.s, f: s.f, r: r}
}

func (s letStopedSender[T]) Tag() SenderTag {
	return s.s.Tag()
}

type letStopedSenderState[T any] struct {
	s Sender[T]
	f func(context.Context, Receiver[T])
	r Receiver[T]
}

func (state letStopedSenderState[T]) Start(ctx context.Context) {
	state.s.Connect(letStopedReceiver[T]{ctx: ctx, f: state.f, r: state.r}).Start(ctx)
}

type letStopedReceiver[T any] struct {
	ctx context.Context
	f   func(context.Context, Receiver[T])
	r   Receiver[T]
}

func (lsr letStopedReceiver[T]) SetValue(v T) {
	lsr.r.SetValue(v)
}

func (lsr letStopedReceiver[T]) SetError(err error) {
	lsr.r.SetError(err)
}

func (lsr letStopedReceiver[T]) SetStoped() {
	lsr.f(lsr.ctx, lsr.r)
}

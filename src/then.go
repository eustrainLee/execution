package src

import (
	"context"

	"github.com/eustrainLee/sr"
)

type thenSender[From any, To any] struct {
	s Sender[From]
	f func(ctx context.Context, v From) To
}

func Then[From any, To any](s Sender[From], f func(context.Context, From) To) Sender[To] {
	return thenSender[From, To]{s: s, f: f}
}

func (s thenSender[From, To]) Connect(r sr.Receiver[To]) OperationState {
	return thenSenderState[From, To]{s: s, r: r}
}

func (s thenSender[From, To]) Tag() sr.SenderTag {
	return s.s.Tag()
}

type thenSenderState[From any, To any] struct {
	s thenSender[From, To]
	r sr.Receiver[To]
}

func (state thenSenderState[From, To]) Start(ctx context.Context) {
	state.s.s.Connect(thenReceiver[From, To]{ctx: ctx, f: state.s.f, r: state.r}).Start(ctx)
}

type thenReceiver[From any, To any] struct {
	ctx context.Context
	f   func(ctx context.Context, v From) To
	r   sr.Receiver[To]
}

func (r thenReceiver[From, To]) SetValue(v From) {
	r.r.SetValue(r.f(r.ctx, v))
}
func (r thenReceiver[From, To]) SetError(err error) {
	r.r.SetError(err)
}
func (r thenReceiver[From, To]) SetStoped() {
	r.r.SetStoped()
}

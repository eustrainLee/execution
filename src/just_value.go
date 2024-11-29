// sender factory: JustValue
package src

import (
	"context"

	"github.com/eustrainLee/sr"
)

type justValueSender[T any] struct {
	v *T
}

func Just[T any](v T) Sender[T] {
	return justValueSender[T]{v: &v}
}

func (s justValueSender[T]) Connect(r sr.Receiver[T]) OperationState {
	return justValueSenderState[T]{s: s, r: r}
}

func (_ justValueSender[T]) Tag() sr.SenderTag {
	return sr.SenderTagMultiShot
}

type justValueSenderState[T any] struct {
	s justValueSender[T]
	r sr.Receiver[T]
}

func (state justValueSenderState[T]) Start(context.Context) {
	if state.r != nil {
		state.r.SetValue(*state.s.v)
	}
}

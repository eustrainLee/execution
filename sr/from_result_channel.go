package sr

import "github.com/samber/mo"

type fromResultChannelSender[T any] struct {
	ch <-chan mo.Result[T]
}

func FromResultChannel[T any](ch <-chan mo.Result[T]) Sender[T] {
	return fromResultChannelSender[T]{ch: ch}
}

func (s fromResultChannelSender[T]) Tag() SenderTag {
	return SenderTagMultiShot
}

func (s fromResultChannelSender[T]) Connect(r Receiver[T]) OperationState {
	return fromResultChannelOperationState[T]{ch: s.ch, r: r}
}

type fromDisposableResultChannelSender[T any] struct {
	ch <-chan mo.Result[T]
}

func FromDisposableResultChannel[T any](ch <-chan mo.Result[T]) Sender[T] {
	return fromDisposableResultChannelSender[T]{ch: ch}
}

func (s fromDisposableResultChannelSender[T]) Tag() SenderTag {
	return SenderTagNone
}

func (s fromDisposableResultChannelSender[T]) Connect(r Receiver[T]) OperationState {
	return fromResultChannelOperationState[T]{ch: s.ch, r: r}
}

type fromResultChannelOperationState[T any] struct {
	ch <-chan mo.Result[T]
	r  Receiver[T]
}

func (os fromResultChannelOperationState[T]) Start() {
	m, ok := <-os.ch
	if ok {
		if m.IsError() {
			os.r.SetError(m.Error())
		} else {
			v, _ := m.Get()
			os.r.SetValue(v)
		}
	} else {
		os.r.SetStoped()
	}
}

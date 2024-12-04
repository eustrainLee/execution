package sr

import "github.com/samber/mo"

func ConnectToResultChannel[T any](s Sender[T], ch chan<- mo.Result[T]) OperationState {
	return s.Connect(startOnChannelReceiver[T](ch))
}

type startOnChannelReceiver[T any] chan<- mo.Result[T]

func (r startOnChannelReceiver[T]) SetValue(v T) {
	r <- mo.Ok(v)
}

func (r startOnChannelReceiver[T]) SetError(err error) {
	r <- mo.Err[T](err)
}

func (r startOnChannelReceiver[T]) SetStoped() {
	close(r)
}

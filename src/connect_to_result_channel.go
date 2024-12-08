package src

import "github.com/samber/mo"

func ConnectToResultChannel[T any](s Sender[T], ch chan<- mo.Result[T], onStoped func()) OperationState {
	return s.Connect(startOnChannelReceiver[T]{
		resultCh: ch,
		onStoped: onStoped,
	})
}

type startOnChannelReceiver[T any] struct {
	resultCh chan<- mo.Result[T]
	onStoped func()
}

func (r startOnChannelReceiver[T]) SetValue(v T) {
	r.resultCh <- mo.Ok(v)
}

func (r startOnChannelReceiver[T]) SetError(err error) {
	r.resultCh <- mo.Err[T](err)
}

func (r startOnChannelReceiver[T]) SetStoped() {
	if r.onStoped != nil {
		r.onStoped()
	}
}

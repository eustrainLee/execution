package sr

import "github.com/samber/mo"

type ResultChannelReceiver[T any] chan<- mo.Result[T]

func (r ResultChannelReceiver[T]) SetValue(v T) {
	r <- mo.Ok(v)
}

func (r ResultChannelReceiver[T]) SetError(err error) {
	r <- mo.Err[T](err)
}

func (r ResultChannelReceiver[T]) SetStoped() {
	close(r)
}

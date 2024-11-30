package sr

type ChannelReceiver[T any] struct {
	ValueChan  chan T
	ErrorChan  chan error
	StopedChan chan struct{}
}

func (r ChannelReceiver[T]) SetValue(v T) {
	r.ValueChan <- v
}

func (r ChannelReceiver[T]) SetError(err error) {
	r.ErrorChan <- err
}

func (r ChannelReceiver[T]) SetStoped() {
	r.StopedChan <- struct{}{}
}

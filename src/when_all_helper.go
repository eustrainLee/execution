package src

type whenAllReceiver[T any] struct {
	ValueChan  chan T
	ErrorChan  chan error
	StopedChan chan struct{}
}

func (r whenAllReceiver[T]) SetValue(v T) {
	r.ValueChan <- v
}

func (r whenAllReceiver[T]) SetError(err error) {
	r.ErrorChan <- err
}

func (r whenAllReceiver[T]) SetStoped() {
	r.StopedChan <- struct{}{}
}

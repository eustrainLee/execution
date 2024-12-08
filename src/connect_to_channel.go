package src

func ConnectToChannel[T any](s Sender[T], valueChan chan<- T, errorChan chan<- error, onStoped func()) OperationState {
	if errorChan != nil {
		return s.Connect(connectToChannelWithErrorReceiver[T]{valueChan: valueChan, errorChan: errorChan, onStoped: onStoped})
	}
	return s.Connect(connectToChannelReceiver[T]{ch: valueChan, onStoped: onStoped})
}

type connectToChannelReceiver[T any] struct {
	ch       chan<- T
	onStoped func()
}

func (r connectToChannelReceiver[T]) SetValue(v T) {
	r.ch <- v
}

func (r connectToChannelReceiver[T]) SetError(err error) {
	panic(err)
}

func (r connectToChannelReceiver[T]) SetStoped() {
	r.onStoped()
}

type connectToChannelWithErrorReceiver[T any] struct {
	valueChan chan<- T
	errorChan chan<- error
	onStoped  func()
}

func (r connectToChannelWithErrorReceiver[T]) SetValue(v T) {
	r.valueChan <- v
}

func (r connectToChannelWithErrorReceiver[T]) SetError(err error) {
	r.errorChan <- err
}

func (r connectToChannelWithErrorReceiver[T]) SetStoped() {
	r.onStoped()
}

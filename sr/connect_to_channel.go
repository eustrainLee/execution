package sr

func ConnectToChannel[T any](s Sender[T], valueChan chan<- T, errorChan chan<- error) OperationState {
	if errorChan != nil {
		return s.Connect(connectToChannelWithErrorReceiver[T]{valueChan: valueChan, errorChan: errorChan})
	}
	return s.Connect(connectToChannelReceiver[T](valueChan))
}

type connectToChannelReceiver[T any] chan<- T

func (r connectToChannelReceiver[T]) SetValue(v T) {
	r <- v
}

func (r connectToChannelReceiver[T]) SetError(err error) {
	panic(err)
}

func (r connectToChannelReceiver[T]) SetStoped() {
	close(r)
}

type connectToChannelWithErrorReceiver[T any] struct {
	valueChan chan<- T
	errorChan chan<- error
}

func (r connectToChannelWithErrorReceiver[T]) SetValue(v T) {
	r.valueChan <- v
}

func (r connectToChannelWithErrorReceiver[T]) SetError(err error) {
	r.errorChan <- err
}

func (r connectToChannelWithErrorReceiver[T]) SetStoped() {
	close(r.valueChan)
	close(r.errorChan)
}

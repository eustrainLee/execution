package sr

type whenAllSliceSender[T any] struct {
	ss []Sender[T]
}

func WhenAllSlice[T any](ss ...Sender[T]) Sender[[]T] {
	return whenAllSliceSender[T]{ss: ss}
}

func (s whenAllSliceSender[T]) Tag() SenderTag {
	return SenderTagNone
}

func (s whenAllSliceSender[T]) Connect(r Receiver[[]T]) OperationState {
	return &whenAllSliceOperationState[T]{ss: s.ss, r: r}
}

type whenAllSliceOperationState[T any] struct {
	ss []Sender[T]
	r  Receiver[[]T]
}

func (os *whenAllSliceOperationState[T]) Start() {
	SenderCount := len(os.ss)
	result := make([]T, 0, SenderCount)
	valuesChan := make([]chan T, 0, SenderCount)
	errChan := make(chan error)
	stopedChan := make(chan struct{}, 2)
	for i := 0; i < SenderCount; i++ {
		valueChan := make(chan T, 1)
		valuesChan = append(valuesChan, valueChan)
		go os.ss[i].Connect(ChannelReceiver[T]{
			ValueChan:  valueChan,
			ErrorChan:  errChan,
			StopedChan: stopedChan,
		}).Start()
	}
	for i := 0; i < SenderCount; i++ {
		select {
		case v := <-valuesChan[i]:
			result = append(result, v)
		case err := <-errChan:
			os.r.SetError(err)
			return
		case <-stopedChan:
			os.r.SetStoped()
			return
		}
	}
	os.r.SetValue(result)
}

package sr

type stopedAsErrorSender[T any] struct {
	s   Sender[T]
	err error
}

func StopedAsError[T any](s Sender[T], err error) Sender[T] {
	return stopedAsErrorSender[T]{s: s}
}

func (s stopedAsErrorSender[T]) Tag() SenderTag {
	return s.s.Tag()
}

func (s stopedAsErrorSender[T]) Connect(r Receiver[T]) OperationState {
	return stopedAsErrorOperationState[T]{s: s.s, r: r, err: s.err}
}

type stopedAsErrorOperationState[T any] struct {
	s   Sender[T]
	r   Receiver[T]
	err error
}

func (os stopedAsErrorOperationState[T]) Start() {
	os.s.Connect(stopedAsErrorReceiver[T]{r: os.r, err: os.err}).Start()
}

type stopedAsErrorReceiver[T any] struct {
	err error
	r   Receiver[T]
}

func (r stopedAsErrorReceiver[T]) SetValue(v T) {
	r.r.SetValue(v)
}

func (r stopedAsErrorReceiver[T]) SetError(err error) {
	r.r.SetError(err)
}

func (r stopedAsErrorReceiver[T]) SetStoped() {
	r.r.SetError(r.err)
}

package sr

type asyncSender[T any] struct {
	s Sender[T]
}

func Async[T any](s Sender[T]) Sender[T] {
	return asyncSender[T]{s: s}
}

func (s asyncSender[T]) Tag() SenderTag {
	return SenderTagNone
}

func (s asyncSender[T]) Connect(r Receiver[T]) OperationState {
	ar := &asyncReceiver[T]{op: make(chan receiverOperation, 1)}
	go s.s.Connect(ar).Start()
	return asyncSenderState[T]{ar: ar, r: r}
}

type asyncSenderState[T any] struct {
	ar *asyncReceiver[T]
	r  Receiver[T]
}

func (state asyncSenderState[T]) Start() {
	switch <-state.ar.op {
	case receiverOperationHasValue:
		state.r.SetValue(state.ar.v)
	case receiverOperationHasError:
		state.r.SetError(state.ar.err)
	case receiverOperationStoped:
		state.r.SetStoped()
	default:
		panic("unknown async state")
	}
}

func (state asyncSenderState[T]) Stop() {
	panic("not supported")
}

type asyncReceiver[T any] struct {
	op     chan receiverOperation
	v      T
	err    error
	stoped bool
}

func (ar *asyncReceiver[T]) SetValue(v T) {
	ar.op <- receiverOperationHasValue
	ar.v = v
}
func (ar *asyncReceiver[T]) SetError(err error) {
	ar.op <- receiverOperationHasError
	ar.err = err
}
func (ar *asyncReceiver[T]) SetStoped() {
	ar.op <- receiverOperationStoped
	ar.stoped = true
}

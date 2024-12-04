package sr

type fromChannelSender[T any] struct {
	ch <-chan T
}

func FromChannel[T any](ch <-chan T) Sender[T] {
	return fromChannelSender[T]{ch: ch}
}

func (s fromChannelSender[T]) Tag() SenderTag {
	return SenderTagMultiShot
}

func (s fromChannelSender[T]) Connect(r Receiver[T]) OperationState {
	return fromChannelOperationState[T]{ch: s.ch, r: r}
}

type fromDisposableChannelSender[T any] struct {
	ch <-chan T
}

func FromDisposableChannel[T any](ch <-chan T) Sender[T] {
	return fromDisposableChannelSender[T]{ch: ch}
}

func (s fromDisposableChannelSender[T]) Tag() SenderTag {
	return SenderTagNone
}

func (s fromDisposableChannelSender[T]) Connect(r Receiver[T]) OperationState {
	return fromChannelOperationState[T]{ch: s.ch, r: r}
}

type fromChannelOperationState[T any] struct {
	ch <-chan T
	r  Receiver[T]
}

func (os fromChannelOperationState[T]) Start() {
	v, ok := <-os.ch
	if ok {
		os.r.SetValue(v)
	} else {
		os.r.SetStoped()
	}
}

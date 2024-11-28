// sender factory: LetError
package sr

type letErrorSender[T any] struct {
	tag SenderTag
	f   func() error
}

func LetError[T any](f func() error, tag SenderTag) Sender[T] {
	return letErrorSender[T]{f: f, tag: tag}
}

func (s letErrorSender[T]) Connect(r Receiver[T]) OperationState {
	return letErrorSenderState[T]{f: s.f, r: r}
}

func (s letErrorSender[T]) Tag() SenderTag {
	return s.tag
}

type letErrorSenderState[T any] struct {
	f func() error
	r Receiver[T]
}

func (state letErrorSenderState[T]) Start() {
	state.r.SetError(state.f())
}

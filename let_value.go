// sender factory: JustResultOf
package sr

type letValueSender[T any] struct {
	tag SenderTag
	f   func() T
}

func LetValue[T any](f func() T, tag SenderTag) Sender[T] {
	return letValueSender[T]{f: f, tag: tag}
}

func (s letValueSender[T]) Connect(r Receiver[T]) OperationState {
	return letValueSenderState[T]{f: s.f, r: r}
}

func (s letValueSender[T]) Tag() SenderTag {
	return s.tag
}

type letValueSenderState[T any] struct {
	f func() T
	r Receiver[T]
}

func (state letValueSenderState[T]) Start() {
	state.r.SetValue(state.f())
}

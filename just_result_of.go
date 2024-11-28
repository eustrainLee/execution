// sender factory: JustResultOf
package sr

type justResultOfSender[T any] struct {
	tag SenderTag
	f   func() T
}

func JustResultOf[T any](f func() T, tag SenderTag) Sender[T] {
	return justResultOfSender[T]{f: f, tag: tag}
}

func (s justResultOfSender[T]) Connect(r Receiver[T]) OperationState {
	return justResultOfSenderState[T]{f: s.f, r: r}
}

func (s justResultOfSender[T]) Tag() SenderTag {
	return s.tag
}

type justResultOfSenderState[T any] struct {
	f func() T
	r Receiver[T]
}

func (state justResultOfSenderState[T]) Start() {
	state.r.SetValue(state.f())
}

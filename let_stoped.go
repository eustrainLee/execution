// sender factory: LetStoped
package sr

type letStopedSender[T any] struct {
	tag SenderTag
	f   func()
}

func LetStoped[T any](f func(), tag SenderTag) Sender[T] {
	return letStopedSender[T]{f: f, tag: tag}
}

func (s letStopedSender[T]) Connect(r Receiver[T]) OperationState {
	return letStopedSenderState[T]{f: s.f, r: r}
}

func (s letStopedSender[T]) Tag() SenderTag {
	return s.tag
}

type letStopedSenderState[T any] struct {
	f func()
	r Receiver[T]
}

func (state letStopedSenderState[T]) Start() {
	state.r.SetStoped()
}

// sender factory: JustResultOf
package sr

type justResultOfSender[T any] struct {
	f func() T
}

func JustResultOf[T any](f func() T) Sender[T] {
	return justResultOfSender[T]{f: f}
}

func (s justResultOfSender[T]) Connect(r Receiver[T]) OperationState {
	return justResultOfSenderState[T]{f: s.f, r: r}
}

func (_ justResultOfSender[T]) Tag() SenderTag {
	return SenderTagMultiSend | SenderTagMultiConnect
}

type justResultOfSenderState[T any] struct {
	f func() T
	r Receiver[T]
}

func (state justResultOfSenderState[T]) Start() {
	state.r.SetValue(state.f())
}

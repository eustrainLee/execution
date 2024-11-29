// sender factory: LetStoped
package sr

type letStopedSender[T any] struct {
	s Sender[T]
	f func(Receiver[T])
}

func LetStoped[T any](s Sender[T], f func(Receiver[T])) Sender[T] {
	return letStopedSender[T]{s: s, f: f}
}

func (s letStopedSender[T]) Connect(r Receiver[T]) OperationState {
	return letStopedSenderState[T]{s: s.s, f: s.f, r: r}
}

func (s letStopedSender[T]) Tag() SenderTag {
	return s.s.Tag()
}

type letStopedSenderState[T any] struct {
	s Sender[T]
	f func(Receiver[T])
	r Receiver[T]
}

func (state letStopedSenderState[T]) Start() {
	state.s.Connect(letStopedReceiver[T]{f: state.f, r: state.r}).Start()
}

type letStopedReceiver[T any] struct {
	f func(Receiver[T])
	r Receiver[T]
}

func (lsr letStopedReceiver[T]) SetValue(v T) {
	lsr.r.SetValue(v)
}

func (lsr letStopedReceiver[T]) SetError(err error) {
	lsr.r.SetError(err)
}

func (lsr letStopedReceiver[T]) SetStoped() {
	lsr.f(lsr.r)
}

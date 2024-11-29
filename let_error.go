// sender factory: LetError
package sr

type letErrorSender[T any] struct {
	s Sender[T]
	f func(error, Receiver[T])
}

func LetError[T any](s Sender[T], f func(error, Receiver[T])) Sender[T] {
	return letErrorSender[T]{s: s, f: f}
}

func (s letErrorSender[T]) Connect(r Receiver[T]) OperationState {
	return letErrorSenderState[T]{s: s.s, f: s.f, r: r}
}

func (s letErrorSender[T]) Tag() SenderTag {
	return s.s.Tag()
}

type letErrorSenderState[T any] struct {
	s Sender[T]
	f func(error, Receiver[T])
	r Receiver[T]
}

func (state letErrorSenderState[T]) Start() {
	state.s.Connect(letErrorReceiver[T]{f: state.f, r: state.r}).Start()
}

type letErrorReceiver[T any] struct {
	f func(error, Receiver[T])
	r Receiver[T]
}

func (ler letErrorReceiver[T]) SetValue(v T) {
	ler.r.SetValue(v)
}

func (ler letErrorReceiver[T]) SetError(err error) {
	ler.f(err, ler.r)
}

func (ler letErrorReceiver[T]) SetStoped() {
	ler.r.SetStoped()
}

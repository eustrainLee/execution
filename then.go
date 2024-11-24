package sr

type thenSender[From any, To any] struct {
	s Sender[From]
	f func(v From) To
}

func Then[From any, To any](s Sender[From], f func(v From) To) Sender[To] {
	return thenSender[From, To]{s: s, f: f}
}

func (s thenSender[From, To]) Connect(r Receiver[To]) OperationState {
	return thenSenderState[From, To]{s: s, r: r}
}

type thenSenderState[From any, To any] struct {
	s thenSender[From, To]
	r Receiver[To]
}

func (state thenSenderState[From, To]) Start() {
	state.s.s.Connect(thenReceiver[From, To]{f: state.s.f, r: state.r}).Start()
}

type thenReceiver[From any, To any] struct {
	f func(v From) To
	r Receiver[To]
}

func (r thenReceiver[From, To]) SetValue(v From) {
	r.r.SetValue(r.f(v))
}
func (r thenReceiver[From, To]) SetError(err error) {
	r.r.SetError(err)
}
func (r thenReceiver[From, To]) SetStoped() {
	r.r.SetStoped()
}

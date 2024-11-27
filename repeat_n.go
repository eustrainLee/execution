package sr

type repeatNSender[T any, I integer] struct {
	n I
	s Sender[T]
}

func RepeatN[T any, I integer](s Sender[T], n I) Sender[T] {
	return repeatNSender[T, I]{s: s, n: n}
}

func (s repeatNSender[T, I]) Connect(r Receiver[T]) OperationState {
	return repeatNSenderState[T, I]{s: s, r: r}
}

func (s repeatNSender[T, I]) Tag() SenderTag {
	return s.s.Tag() | SenderTagMultiSend
}

type repeatNSenderState[T any, I integer] struct {
	s repeatNSender[T, I]
	r Receiver[T]
}

func (state repeatNSenderState[T, I]) Start() {
	recv := TrivalReceiver[T]{}
	state.s.s.Connect(&recv).Start()
	switch {
	case recv.Error != nil:
		for i := I(0); i < state.s.n; i++ {
			state.r.SetError(recv.Error)
		}
	case recv.Stoped:
		state.r.SetStoped()
	default:
		for i := I(0); i < state.s.n; i++ {
			state.r.SetValue(recv.Value)
		}
	}
}

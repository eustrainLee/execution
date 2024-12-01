package sr

import (
	"github.com/samber/lo"
)

type whenAll2Sender[T1, T2 any] struct {
	s1 Sender[T1]
	s2 Sender[T2]
}

func WhenAll2[T1, T2 any](s1 Sender[T1], s2 Sender[T2]) Sender[lo.Tuple2[T1, T2]] {
	return whenAll2Sender[T1, T2]{s1: s1, s2: s2}
}

func (s whenAll2Sender[T1, T2]) Tag() SenderTag {
	return SenderTagNone
}

func (s whenAll2Sender[T1, T2]) Connect(r Receiver[lo.Tuple2[T1, T2]]) OperationState {
	return whenAll2OperationState[T1, T2]{s: s, r: r}
}

type whenAll2OperationState[T1, T2 any] struct {
	s whenAll2Sender[T1, T2]
	r Receiver[lo.Tuple2[T1, T2]]
}

func (os whenAll2OperationState[T1, T2]) Start() {
	result := lo.Tuple2[T1, T2]{}
	errChan := make(chan error)
	stopedChan := make(chan struct{}, 2)
	v1Chan := make(chan T1, 1)
	go os.s.s1.Connect(ChannelReceiver[T1]{
		ValueChan:  v1Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	v2Chan := make(chan T2, 1)
	go os.s.s2.Connect(ChannelReceiver[T2]{
		ValueChan:  v2Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	const SenderCount = 2
	for i := 0; i < SenderCount; i++ {
		select {
		case v := <-v1Chan:
			result.A = v
		case v := <-v2Chan:
			result.B = v
		case err := <-errChan:
			os.r.SetError(err)
			return
		case <-stopedChan:
			os.r.SetStoped()
			return
		}
	}
	os.r.SetValue(result)
}

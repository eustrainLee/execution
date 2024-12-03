package sr

import (
	"github.com/samber/lo"
)

type whenAll4Sender[T1, T2, T3, T4 any] struct {
	s1 Sender[T1]
	s2 Sender[T2]
	s3 Sender[T3]
	s4 Sender[T4]
}

func WhenAll4[T1, T2, T3, T4 any](s1 Sender[T1], s2 Sender[T2], s3 Sender[T3], s4 Sender[T4]) Sender[lo.Tuple4[T1, T2, T3, T4]] {
	return whenAll4Sender[T1, T2, T3, T4]{s1: s1, s2: s2, s3: s3, s4: s4}
}

func (s whenAll4Sender[T1, T2, T3, T4]) Tag() SenderTag {
	return SenderTagNone
}

func (s whenAll4Sender[T1, T2, T3, T4]) Connect(r Receiver[lo.Tuple4[T1, T2, T3, T4]]) OperationState {
	return whenAll4OperationState[T1, T2, T3, T4]{s: s, r: r}
}

type whenAll4OperationState[T1, T2, T3, T4 any] struct {
	s whenAll4Sender[T1, T2, T3, T4]
	r Receiver[lo.Tuple4[T1, T2, T3, T4]]
}

func (os whenAll4OperationState[T1, T2, T3, T4]) Start() {
	const SenderCount = 4
	result := lo.Tuple4[T1, T2, T3, T4]{}
	errChan := make(chan error)
	stopedChan := make(chan struct{}, SenderCount)
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
	v3Chan := make(chan T3, 1)
	go os.s.s3.Connect(ChannelReceiver[T3]{
		ValueChan:  v3Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	v4Chan := make(chan T4, 1)
	go os.s.s4.Connect(ChannelReceiver[T4]{
		ValueChan:  v4Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	for i := 0; i < SenderCount; i++ {
		select {
		case result.A = <-v1Chan:
		case result.B = <-v2Chan:
		case result.C = <-v3Chan:
		case result.D = <-v4Chan:
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

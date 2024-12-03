package sr

import (
	"github.com/samber/lo"
)

type whenAll7Sender[T1, T2, T3, T4, T5, T6, T7 any] struct {
	s1 Sender[T1]
	s2 Sender[T2]
	s3 Sender[T3]
	s4 Sender[T4]
	s5 Sender[T5]
	s6 Sender[T6]
	s7 Sender[T7]
}

func WhenAll7[T1, T2, T3, T4, T5, T6, T7 any](s1 Sender[T1], s2 Sender[T2], s3 Sender[T3], s4 Sender[T4], s5 Sender[T5], s6 Sender[T6], s7 Sender[T7]) Sender[lo.Tuple7[T1, T2, T3, T4, T5, T6, T7]] {
	return whenAll7Sender[T1, T2, T3, T4, T5, T6, T7]{s1: s1, s2: s2, s3: s3, s4: s4, s5: s5, s6: s6, s7: s7}
}

func (s whenAll7Sender[T1, T2, T3, T4, T5, T6, T7]) Tag() SenderTag {
	return SenderTagNone
}

func (s whenAll7Sender[T1, T2, T3, T4, T5, T6, T7]) Connect(r Receiver[lo.Tuple7[T1, T2, T3, T4, T5, T6, T7]]) OperationState {
	return whenAll7OperationState[T1, T2, T3, T4, T5, T6, T7]{s: s, r: r}
}

type whenAll7OperationState[T1, T2, T3, T4, T5, T6, T7 any] struct {
	s whenAll7Sender[T1, T2, T3, T4, T5, T6, T7]
	r Receiver[lo.Tuple7[T1, T2, T3, T4, T5, T6, T7]]
}

func (os whenAll7OperationState[T1, T2, T3, T4, T5, T6, T7]) Start() {
	const SenderCount = 7
	result := lo.Tuple7[T1, T2, T3, T4, T5, T6, T7]{}
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
	v5Chan := make(chan T5, 1)
	go os.s.s5.Connect(ChannelReceiver[T5]{
		ValueChan:  v5Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	v6Chan := make(chan T6, 1)
	go os.s.s6.Connect(ChannelReceiver[T6]{
		ValueChan:  v6Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	v7Chan := make(chan T7, 1)
	go os.s.s7.Connect(ChannelReceiver[T7]{
		ValueChan:  v7Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start()
	for i := 0; i < SenderCount; i++ {
		select {
		case result.A = <-v1Chan:
		case result.B = <-v2Chan:
		case result.C = <-v3Chan:
		case result.D = <-v4Chan:
		case result.E = <-v5Chan:
		case result.F = <-v6Chan:
		case result.G = <-v7Chan:
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

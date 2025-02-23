package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
	"github.com/samber/lo"
)

type whenAll6Sender[T1, T2, T3, T4, T5, T6 any] struct {
	s1 Sender[T1]
	s2 Sender[T2]
	s3 Sender[T3]
	s4 Sender[T4]
	s5 Sender[T5]
	s6 Sender[T6]
}

func WhenAll6[T1, T2, T3, T4, T5, T6 any](s1 Sender[T1], s2 Sender[T2], s3 Sender[T3], s4 Sender[T4], s5 Sender[T5], s6 Sender[T6]) Sender[lo.Tuple6[T1, T2, T3, T4, T5, T6]] {
	return whenAll6Sender[T1, T2, T3, T4, T5, T6]{s1: s1, s2: s2, s3: s3, s4: s4, s5: s5, s6: s6}
}

func (s whenAll6Sender[T1, T2, T3, T4, T5, T6]) Tag() sr.SenderTag {
	return sr.SenderTagNone
}

func (s whenAll6Sender[T1, T2, T3, T4, T5, T6]) Connect(r sr.Receiver[lo.Tuple6[T1, T2, T3, T4, T5, T6]]) OperationState {
	return whenAll6OperationState[T1, T2, T3, T4, T5, T6]{s: s, r: r}
}

type whenAll6OperationState[T1, T2, T3, T4, T5, T6 any] struct {
	s whenAll6Sender[T1, T2, T3, T4, T5, T6]
	r sr.Receiver[lo.Tuple6[T1, T2, T3, T4, T5, T6]]
}

func (os whenAll6OperationState[T1, T2, T3, T4, T5, T6]) Start(ctx context.Context) {
	const SenderCount = 6
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	result := lo.Tuple6[T1, T2, T3, T4, T5, T6]{}
	errChan := make(chan error)
	stopedChan := make(chan struct{}, SenderCount)
	v1Chan := make(chan T1, 1)
	go os.s.s1.Connect(sr.ChannelReceiver[T1]{
		ValueChan:  v1Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v2Chan := make(chan T2, 1)
	go os.s.s2.Connect(sr.ChannelReceiver[T2]{
		ValueChan:  v2Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v3Chan := make(chan T3, 1)
	go os.s.s3.Connect(sr.ChannelReceiver[T3]{
		ValueChan:  v3Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v4Chan := make(chan T4, 1)
	go os.s.s4.Connect(sr.ChannelReceiver[T4]{
		ValueChan:  v4Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v5Chan := make(chan T5, 1)
	go os.s.s5.Connect(sr.ChannelReceiver[T5]{
		ValueChan:  v5Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v6Chan := make(chan T6, 1)
	go os.s.s6.Connect(sr.ChannelReceiver[T6]{
		ValueChan:  v6Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	for i := 0; i < SenderCount; i++ {
		select {
		case result.A = <-v1Chan:
		case result.B = <-v2Chan:
		case result.C = <-v3Chan:
		case result.D = <-v4Chan:
		case result.E = <-v5Chan:
		case result.F = <-v6Chan:
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

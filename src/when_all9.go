package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
	"github.com/samber/lo"
)

type whenAll9Sender[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	s1 Sender[T1]
	s2 Sender[T2]
	s3 Sender[T3]
	s4 Sender[T4]
	s5 Sender[T5]
	s6 Sender[T6]
	s7 Sender[T7]
	s8 Sender[T8]
	s9 Sender[T9]
}

func WhenAll9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](s1 Sender[T1], s2 Sender[T2], s3 Sender[T3], s4 Sender[T4], s5 Sender[T5], s6 Sender[T6], s7 Sender[T7], s8 Sender[T8], s9 Sender[T9]) Sender[lo.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]] {
	return whenAll9Sender[T1, T2, T3, T4, T5, T6, T7, T8, T9]{s1: s1, s2: s2, s3: s3, s4: s4, s5: s5, s6: s6, s7: s7, s8: s8, s9: s9}
}

func (s whenAll9Sender[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Tag() sr.SenderTag {
	return sr.SenderTagNone
}

func (s whenAll9Sender[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Connect(r sr.Receiver[lo.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]) OperationState {
	return whenAll9OperationState[T1, T2, T3, T4, T5, T6, T7, T8, T9]{s: s, r: r}
}

type whenAll9OperationState[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	s whenAll9Sender[T1, T2, T3, T4, T5, T6, T7, T8, T9]
	r sr.Receiver[lo.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]]
}

func (os whenAll9OperationState[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Start(ctx context.Context) {
	const SenderCount = 9
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	result := lo.Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{}
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
	v7Chan := make(chan T7, 1)
	go os.s.s7.Connect(sr.ChannelReceiver[T7]{
		ValueChan:  v7Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v8Chan := make(chan T8, 1)
	go os.s.s8.Connect(sr.ChannelReceiver[T8]{
		ValueChan:  v8Chan,
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}).Start(ctx)
	v9Chan := make(chan T9, 1)
	go os.s.s9.Connect(sr.ChannelReceiver[T9]{
		ValueChan:  v9Chan,
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
		case result.G = <-v7Chan:
		case result.H = <-v8Chan:
		case result.I = <-v9Chan:
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

package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
	"github.com/samber/lo"
)

type whenAll2Sender[T1, T2 any] struct {
	s1 Sender[T1]
	s2 Sender[T2]
}

func WhenAll2[T1, T2 any](s1 Sender[T1], s2 Sender[T2]) Sender[lo.Tuple2[T1, T2]] {
	return whenAll2Sender[T1, T2]{s1: s1, s2: s2}
}

func (s whenAll2Sender[T1, T2]) Tag() sr.SenderTag {
	return sr.SenderTagNone
}

func (s whenAll2Sender[T1, T2]) Connect(r sr.Receiver[lo.Tuple2[T1, T2]]) OperationState {
	return whenAll2OperationState[T1, T2]{s: s, r: r}
}

type whenAll2OperationState[T1, T2 any] struct {
	s whenAll2Sender[T1, T2]
	r sr.Receiver[lo.Tuple2[T1, T2]]
}

func (os whenAll2OperationState[T1, T2]) Start(ctx context.Context) {
	const SenderCount = 2
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	result := lo.Tuple2[T1, T2]{}
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
	for i := 0; i < SenderCount; i++ {
		select {
		case result.A = <-v1Chan:
		case result.B = <-v2Chan:
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

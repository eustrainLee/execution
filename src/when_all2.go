package src

import (
	"context"

	"github.com/eustrainLee/sr"
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
	return &whenAll2OperationState[T1, T2]{s: s, r: r}
}

type whenAll2OperationState[T1, T2 any] struct {
	s  whenAll2Sender[T1, T2]
	r  sr.Receiver[lo.Tuple2[T1, T2]]
	br whenAll2Receiver[T1, T2]
}

func (os *whenAll2OperationState[T1, T2]) Start(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	result := lo.Tuple2[T1, T2]{}
	errChan := make(chan error)
	stopedChan := make(chan struct{}, 2)
	os.br.r1 = whenAllReceiver[T1]{
		ValueChan:  make(chan T1, 1),
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}
	go os.s.s1.Connect(os.br.r1).Start(ctx)
	os.br.r2 = whenAllReceiver[T2]{
		ValueChan:  make(chan T2, 1),
		ErrorChan:  errChan,
		StopedChan: stopedChan,
	}
	go os.s.s2.Connect(os.br.r2).Start(ctx)
	const SenderCount = 2
	for i := 0; i < SenderCount; i++ {
		select {
		case v := <-os.br.r1.ValueChan:
			result.A = v
		case v := <-os.br.r2.ValueChan:
			result.B = v
		case err := <-errChan:
			os.r.SetError(err)
		case <-stopedChan:
			os.r.SetStoped()
		}
	}
	os.r.SetValue(result)
}

type whenAll2Receiver[T1, T2 any] struct {
	r1 whenAllReceiver[T1]
	r2 whenAllReceiver[T2]
}

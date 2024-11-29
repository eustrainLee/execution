package sr

import "errors"

type TrivalReceiver[T any] struct {
	Op     ReceiverOperation
	Value  T
	Error  error
	Stoped bool
}

func (tr *TrivalReceiver[T]) SetValue(v T) {
	tr.Value = v
}
func (tr *TrivalReceiver[T]) SetError(err error) {
	tr.Error = err
}
func (tr *TrivalReceiver[T]) SetStoped() {
	tr.Stoped = true
}

func (tr *TrivalReceiver[T]) Forward(r Receiver[T]) {
	switch tr.Op {
	case ReceiverOperationStoped:
		r.SetStoped()
	case ReceiverOperationHasError:
		r.SetError(tr.Error)
	case ReceiverOperationHasValue:
		r.SetValue(tr.Value)
	default:
		panic(errors.New("receiver not ready"))
	}
}

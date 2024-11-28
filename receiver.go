package sr

import "errors"

type TrivalReceiver[T any] struct {
	Op     receiverOperation
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
	case receiverOperationStoped:
		r.SetStoped()
	case receiverOperationHasError:
		r.SetError(tr.Error)
	case receiverOperationHasValue:
		r.SetValue(tr.Value)
	default:
		panic(errors.New("receiver not ready"))
	}
}

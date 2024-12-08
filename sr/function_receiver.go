package sr

type FunctionReceiver[T any] struct {
	ValueFunc  func(T)
	ErrorFunc  func(error)
	StopedFunc func()
}

func (r FunctionReceiver[T]) SetValue(v T) {
	r.ValueFunc(v)
}

func (r FunctionReceiver[T]) SetError(err error) {
	r.ErrorFunc(err)
}

func (r FunctionReceiver[T]) SetStoped() {
	r.StopedFunc()
}

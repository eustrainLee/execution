package sr

type TrivalReceiver[T any] struct {
	Value  T
	Error  error
	Stoped bool
}

func (r *TrivalReceiver[T]) SetValue(v T) {
	r.Value = v
}
func (r *TrivalReceiver[T]) SetError(err error) {
	r.Error = err
}
func (r *TrivalReceiver[T]) SetStoped() {
	r.Stoped = true
}

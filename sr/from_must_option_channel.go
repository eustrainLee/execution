package sr

import "github.com/samber/mo"

type fromMustOptionChannelSender[T any] struct {
	ch  <-chan mo.Option[T]
	err error
}

func FromMustOptionChannel[T any](ch <-chan mo.Option[T], err error) Sender[T] {
	return fromMustOptionChannelSender[T]{ch: ch, err: err}
}

func (s fromMustOptionChannelSender[T]) Tag() SenderTag {
	return SenderTagMultiShot
}

func (s fromMustOptionChannelSender[T]) Connect(r Receiver[T]) OperationState {
	return fromMustOptionChannelOperationState[T]{ch: s.ch, err: s.err, r: r}
}

type fromDisposableMustOptionChannelSender[T any] struct {
	ch  <-chan mo.Option[T]
	err error
}

func FromDisposableOptionChannel[T any](ch <-chan mo.Option[T], err error) Sender[T] {
	return fromDisposableMustOptionChannelSender[T]{ch: ch, err: err}
}

func (s fromDisposableMustOptionChannelSender[T]) Tag() SenderTag {
	return SenderTagNone
}

func (s fromDisposableMustOptionChannelSender[T]) Connect(r Receiver[T]) OperationState {
	return fromMustOptionChannelOperationState[T]{ch: s.ch, err: s.err, r: r}
}

type fromMustOptionChannelOperationState[T any] struct {
	ch  <-chan mo.Option[T]
	err error
	r   Receiver[T]
}

func (os fromMustOptionChannelOperationState[T]) Start() {
	m, ok := <-os.ch
	if ok {
		v, ok := m.Get()
		if ok {
			os.r.SetValue(v)
		} else {
			os.r.SetError(os.err)
		}
	} else {
		os.r.SetStoped()
	}
}

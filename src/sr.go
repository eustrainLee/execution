package src

import (
	"context"

	"github.com/eustrainLee/execution/sr"
)

type Sender[T any] interface {
	Tag() sr.SenderTag
	Connect(sr.Receiver[T]) OperationState
}

type OperationState interface {
	Start(context.Context)
}

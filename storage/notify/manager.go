package notify

import (
	"context"
	"fmt"
)

type Manager interface {
	Add(...Provider)
	Run(context.Context) error
}

type Provider interface {
	fmt.Stringer
	// Start 启动
	Start(context.Context) error
	// Attempt 是否允许启动
	Attempt() bool
}

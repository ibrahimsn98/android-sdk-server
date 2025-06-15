package ports

import "context"

type Executor interface {
	RunCommand(ctx context.Context, name string, args ...string) (*Output, error)
}

type Output struct {
	Stdout string
	Stderr string
}

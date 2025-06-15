package ports

import "context"

type Emulator interface {
	Start(ctx context.Context, avdName string, args ...string) (*Output, error)
}

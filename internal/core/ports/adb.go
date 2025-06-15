package ports

import "context"

type ADB interface {
	Devices(ctx context.Context) (*Output, error)
	StartServer(ctx context.Context) (*Output, error)
	StopServer(ctx context.Context) (*Output, error)
}

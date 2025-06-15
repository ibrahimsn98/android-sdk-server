package ports

import "context"

type ADB interface {
	Devices(ctx context.Context) (*Output, error)
	StopDevice(ctx context.Context, deviceSerial string) (*Output, error)
	RestartDevice(ctx context.Context, deviceSerial string) (*Output, error)
	StartServer(ctx context.Context) (*Output, error)
	StopServer(ctx context.Context) (*Output, error)
}

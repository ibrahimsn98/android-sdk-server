package ports

import "context"

type AVDManager interface {
	CreateAVD(
		ctx context.Context,
		avdManagerArgs AVDManagerArgs,
		name string,
		packagePath string,
		options ...string,
	) (*Output, error)
	DeleteAVD(ctx context.Context, avdManagerArgs AVDManagerArgs, name string) (*Output, error)
	ListAVDs(ctx context.Context, avdManagerArgs AVDManagerArgs) (*Output, error)
}

type AVDManagerArgs struct {
	SDKVersion string
}

package ports

import "context"

type AVDManager interface {
	CreateAVD(ctx context.Context, name string, packagePath string, options ...string) (*Output, error)
	DeleteAVD(ctx context.Context, name string) (*Output, error)
	ListAVDs(ctx context.Context) (*Output, error)
}

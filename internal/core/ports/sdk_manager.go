package ports

import "context"

type SDKManager interface {
	UpdateAll(ctx context.Context) (*Output, error)
	ListPackages(ctx context.Context) (*Output, error)
	InstallPackages(ctx context.Context, packages []string) (*Output, error)
}

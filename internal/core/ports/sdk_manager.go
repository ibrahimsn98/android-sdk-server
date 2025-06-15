package ports

import "context"

type SDKManager interface {
	UpdateAll(ctx context.Context, sdkManagerArgs SDKManagerArgs) (*Output, error)
	ListPackages(ctx context.Context, sdkManagerArgs SDKManagerArgs) (*Output, error)
	InstallPackages(ctx context.Context, sdkManagerArgs SDKManagerArgs, packages []string) (*Output, error)
}

type SDKManagerArgs struct {
	SDKVersion string
}

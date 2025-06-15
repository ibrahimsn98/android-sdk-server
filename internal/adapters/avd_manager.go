package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
)

type AVDManager struct {
	executor ports.Executor
	binPath  string
}

func NewAVDManager(exec ports.Executor, binPath string) *AVDManager {
	return &AVDManager{executor: exec, binPath: binPath}
}

func (a *AVDManager) CreateAVD(ctx context.Context, name string, packagePath string, options ...string) (*ports.Output, error) {
	args := []string{"create", "avd", "--name", name, "--package", packagePath}
	args = append(args, options...)
	return a.executor.RunCommand(ctx, a.binPath, args...)
}

func (a *AVDManager) DeleteAVD(ctx context.Context, name string) (*ports.Output, error) {
	args := []string{"delete", "avd", "--name", name}
	return a.executor.RunCommand(ctx, a.binPath, args...)
}

func (a *AVDManager) ListAVDs(ctx context.Context) (*ports.Output, error) {
	args := []string{"list", "avd"}
	return a.executor.RunCommand(ctx, a.binPath, args...)
}

package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
)

type SDKManager struct {
	executor ports.Executor
	binPath  string
}

func NewSDKManager(exec ports.Executor, binPath string) *SDKManager {
	return &SDKManager{executor: exec, binPath: binPath}
}

func (s *SDKManager) UpdateAll(ctx context.Context) (*ports.Output, error) {
	return s.executor.RunCommand(ctx, s.binPath, "--update")
}

func (s *SDKManager) ListPackages(ctx context.Context) (*ports.Output, error) {
	out, err := s.executor.RunCommand(ctx, s.binPath, "--list")
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SDKManager) InstallPackages(ctx context.Context, packages []string) (*ports.Output, error) {
	args := append([]string{"--install"}, packages...)
	return s.executor.RunCommand(ctx, s.binPath, args...)
}

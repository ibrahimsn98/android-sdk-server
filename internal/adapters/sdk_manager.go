package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
	"path/filepath"
)

type SDKManager struct {
	executor ports.Executor
	sdkPath  string
}

func NewSDKManager(exec ports.Executor, sdkPath string) *SDKManager {
	return &SDKManager{executor: exec, sdkPath: sdkPath}
}

func (s *SDKManager) getBinPath(sdkManagerArgs *ports.SDKManagerArgs) string {
	return filepath.Join(s.sdkPath, "cmdline-tools", sdkManagerArgs.SDKVersion, "bin", "sdkmanager")
}

func (s *SDKManager) UpdateAll(ctx context.Context, sdkManagerArgs *ports.SDKManagerArgs) (*ports.Output, error) {
	return s.executor.RunCommand(ctx, s.getBinPath(sdkManagerArgs), []string{"--update"}, []string{})
}

func (s *SDKManager) ListPackages(ctx context.Context, sdkManagerArgs *ports.SDKManagerArgs) (*ports.Output, error) {
	out, err := s.executor.RunCommand(ctx, s.getBinPath(sdkManagerArgs), []string{"--list"}, []string{})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s *SDKManager) InstallPackages(
	ctx context.Context,
	sdkManagerArgs *ports.SDKManagerArgs,
	packages []string,
) (*ports.Output, error) {
	args := append([]string{"--install"}, packages...)
	return s.executor.RunCommand(ctx, s.getBinPath(sdkManagerArgs), args, []string{})
}

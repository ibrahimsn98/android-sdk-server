package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
	"path/filepath"
)

type AVDManager struct {
	executor ports.Executor
	sdkPath  string
}

func NewAVDManager(exec ports.Executor, sdkPath string) *AVDManager {
	return &AVDManager{executor: exec, sdkPath: sdkPath}
}

func (a *AVDManager) getBinPath(avdManagerArgs *ports.AVDManagerArgs) string {
	return filepath.Join(a.sdkPath, "cmdline-tools", avdManagerArgs.SDKVersion, "bin", "avdmanager")
}

func (a *AVDManager) CreateAVD(
	ctx context.Context,
	avdManagerArgs *ports.AVDManagerArgs,
	name string,
	packagePath string,
	options []string,
) (*ports.Output, error) {
	args := []string{"create", "avd", "--name", name, "--package", packagePath}
	args = append(args, options...)
	return a.executor.RunCommand(ctx, a.getBinPath(avdManagerArgs), args, []string{"no"})
}

func (a *AVDManager) DeleteAVD(
	ctx context.Context,
	avdManagerArgs *ports.AVDManagerArgs,
	name string,
) (*ports.Output, error) {
	args := []string{"delete", "avd", "--name", name}
	return a.executor.RunCommand(ctx, a.getBinPath(avdManagerArgs), args, []string{})
}

func (a *AVDManager) ListAVDs(
	ctx context.Context,
	avdManagerArgs *ports.AVDManagerArgs,
) (*ports.Output, error) {
	args := []string{"list", "avd"}
	return a.executor.RunCommand(ctx, a.getBinPath(avdManagerArgs), args, []string{})
}

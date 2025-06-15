package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
	"fmt"
)

type AVDManager struct {
	executor ports.Executor
	sdkPath  string
}

func NewAVDManager(exec ports.Executor, sdkPath string) *AVDManager {
	return &AVDManager{executor: exec, sdkPath: sdkPath}
}

func (a *AVDManager) getBinPath(avdManagerArgs ports.AVDManagerArgs) string {
	return fmt.Sprintf("%s/cmdline-tools/%s/bin/avdmanager", a.sdkPath, avdManagerArgs.SDKVersion)
}

func (a *AVDManager) CreateAVD(
	ctx context.Context,
	avdManagerArgs ports.AVDManagerArgs,
	name string,
	packagePath string,
	options ...string,
) (*ports.Output, error) {
	args := []string{"create", "avd", "--name", name, "--package", packagePath}
	args = append(args, options...)
	return a.executor.RunCommand(ctx, a.getBinPath(avdManagerArgs), args...)
}

func (a *AVDManager) DeleteAVD(
	ctx context.Context,
	avdManagerArgs ports.AVDManagerArgs,
	name string,
) (*ports.Output, error) {
	args := []string{"delete", "avd", "--name", name}
	return a.executor.RunCommand(ctx, a.getBinPath(avdManagerArgs), args...)
}

func (a *AVDManager) ListAVDs(
	ctx context.Context,
	avdManagerArgs ports.AVDManagerArgs,
) (*ports.Output, error) {
	args := []string{"list", "avd"}
	return a.executor.RunCommand(ctx, a.getBinPath(avdManagerArgs), args...)
}

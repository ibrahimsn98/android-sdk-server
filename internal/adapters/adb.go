package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
)

type Adb struct {
	executor ports.Executor
	binPath  string
}

func NewAdb(exec ports.Executor, binPath string) *Adb {
	return &Adb{executor: exec, binPath: binPath}
}

func (a *Adb) Devices(ctx context.Context) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, "devices")
}

func (a *Adb) StartServer(ctx context.Context) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, "start-server")
}

func (a *Adb) StopServer(ctx context.Context) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, "kill-server")
}

func (a *Adb) InstallAPK(ctx context.Context, deviceID string, apkPath string) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, "-s", deviceID, "install", apkPath)
}

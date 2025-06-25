package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
	"path/filepath"
)

type Adb struct {
	executor ports.Executor
	binPath  string
}

func NewAdb(exec ports.Executor, sdkPath string) *Adb {
	return &Adb{executor: exec, binPath: filepath.Join(sdkPath, "platform-tools", "adb")}
}

func (a *Adb) Devices(ctx context.Context) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, []string{"devices"}, []string{})
}

func (a *Adb) StopDevice(ctx context.Context, deviceSerial string) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, []string{"-s", deviceSerial, "emu", "kill"}, []string{})
}

func (a *Adb) RestartDevice(ctx context.Context, deviceSerial string) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, []string{"-s", deviceSerial, "emu", "restart"}, []string{})
}

func (a *Adb) StartServer(ctx context.Context) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, []string{"start-server"}, []string{})
}

func (a *Adb) StopServer(ctx context.Context) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, []string{"kill-server"}, []string{})
}

func (a *Adb) InstallAPK(ctx context.Context, deviceID string, apkPath string) (*ports.Output, error) {
	return a.executor.RunCommand(ctx, a.binPath, []string{"-s", deviceID, "install", apkPath}, []string{})
}

package adapters

import (
	"android-cmd-server/internal/core/ports"
	"context"
	"path/filepath"
)

type Emulator struct {
	executor ports.Executor
	binPath  string
}

func NewEmulator(exec ports.Executor, sdkPath string) *Emulator {
	return &Emulator{executor: exec, binPath: filepath.Join(sdkPath, "emulator", "emulator")}
}

func (e *Emulator) Start(ctx context.Context, avdName string, args []string) (*ports.Output, error) {
	allArgs := append([]string{"-avd", avdName}, args...)
	return e.executor.RunCommand(ctx, e.binPath, allArgs, []string{})
}

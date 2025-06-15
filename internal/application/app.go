package application

import (
	"android-cmd-server/internal/adapters"
	"android-cmd-server/internal/api"
	"android-cmd-server/internal/infrastructure/shell"
	"android-cmd-server/internal/infrastructure/system"
	"context"
)

type App interface {
	Startup(context.Context, Module) error
}

type Application struct {
}

func (a *Application) Startup(ctx context.Context, app Module) (err error) {
	executor := shell.NewExecutor()

	sdkPath, err := system.FindAndroidSDKPath()
	if err != nil {
		return err
	}

	sdkManager := adapters.NewSDKManager(executor, sdkPath)
	avdManager := adapters.NewAVDManager(executor, sdkPath)
	adb := adapters.NewAdb(executor, sdkPath)
	emulator := adapters.NewEmulator(executor, sdkPath)

	api.NewAVDController(app.Api(), avdManager)
	api.NewSDKController(app.Api(), sdkManager)
	api.NewADBController(app.Api(), adb)
	api.NewEmulatorController(app.Api(), emulator)
	return nil
}

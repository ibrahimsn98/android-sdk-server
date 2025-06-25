package application

import (
	"android-cmd-server/internal/adapters"
	"android-cmd-server/internal/api"
	"android-cmd-server/internal/infrastructure/config"
	"android-cmd-server/internal/infrastructure/shell"
	"context"
)

type App interface {
	Startup(context.Context, Module) error
}

type Application struct {
}

func (a *Application) Startup(_ context.Context, app Module) (err error) {
	executor := shell.NewExecutor()

	cfg, err := config.InitiateConfig()
	if err != nil {
		return err
	}

	sdkManager := adapters.NewSDKManager(executor, cfg.SdkPath)
	avdManager := adapters.NewAVDManager(executor, cfg.SdkPath)
	adb := adapters.NewAdb(executor, cfg.SdkPath)
	emulator := adapters.NewEmulator(executor, cfg.SdkPath)

	api.NewAVDController(app.Api(), avdManager)
	api.NewSDKController(app.Api(), sdkManager)
	api.NewADBController(app.Api(), adb)
	api.NewEmulatorController(app.Api(), emulator)
	return nil
}

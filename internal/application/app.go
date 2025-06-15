package application

import (
	"android-cmd-server/internal/adapters"
	"android-cmd-server/internal/api"
	"android-cmd-server/internal/infrastructure/shell"
	"context"
	"fmt"
	"path/filepath"
)

type App interface {
	Startup(context.Context, Module) error
}

type Application struct {
}

func (a *Application) Startup(ctx context.Context, app Module) (err error) {
	executor := shell.NewExecutor()

	androidHome := "/Users/ibrahim.suren/Library/Android/sdk"
	version := "17.0"

	sdkManager := adapters.NewSDKManager(
		executor,
		filepath.Join(androidHome, "cmdline-tools", version, "bin", "sdkmanager"),
	)
	avdManager := adapters.NewAVDManager(
		executor,
		filepath.Join(androidHome, "cmdline-tools", version, "bin", "avdmanager"),
	)
	adb := adapters.NewAdb(
		executor,
		filepath.Join(androidHome, "platform-tools", "adb"),
	)
	emulator := adapters.NewEmulator(
		executor,
		filepath.Join(androidHome, "tools", "emulator"),
	)

	fmt.Println(emulator)

	api.NewAVDController(app.Api(), avdManager)
	api.NewSDKController(app.Api(), sdkManager)
	api.NewADBController(app.Api(), adb)
	return nil
}

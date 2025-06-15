package main

import (
	"android-cmd-server/internal/application"
	"android-cmd-server/internal/infrastructure/logger"
	"android-cmd-server/internal/infrastructure/waiter"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
	"time"
)

type module struct {
	logger *logger.Logger
	api    *echo.Echo
	app    application.App
	waiter waiter.Waiter
}

func (a *module) Logger() *logger.Logger {
	return a.logger
}

func (a *module) Api() *echo.Echo {
	return a.api
}

func (a *module) Waiter() waiter.Waiter {
	return a.waiter
}

func (a *module) waitForApi(ctx context.Context) error {
	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		a.Logger().Info("[Rest] Server is started...")
		defer a.Logger().Info("[Rest] Server shutdown.")

		if err := a.api.Start("0.0.0.0:8001"); err != nil {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		a.Logger().Info("[Rest] Server to be shutdown...")
		stopped := make(chan struct{})
		go func() {
			err := a.api.Shutdown(ctx)
			if err != nil {
				return
			}
			close(stopped)
		}()
		timeout := time.NewTimer(2000 * time.Millisecond)
		select {
		case <-timeout.C:
			err := a.api.Shutdown(ctx)
			if err != nil {
				return err
			}
			return fmt.Errorf("server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})
	return group.Wait()
}

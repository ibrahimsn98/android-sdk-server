package api

import (
	"android-cmd-server/internal/core/domain"
	"android-cmd-server/internal/core/domain/adb"
	"android-cmd-server/internal/core/ports"
	"android-cmd-server/internal/infrastructure/server"
	"github.com/labstack/echo/v4"
)

type ADBController struct {
	adb ports.ADB
}

func NewADBController(
	echo *echo.Echo,
	adb ports.ADB,
) {
	c := &ADBController{
		adb: adb,
	}
	echo.GET("adb/devices", server.Handle(c.devices()))
	echo.POST("adb/stop-device", server.Handle(c.stopDevice()))
	echo.POST("adb/restart-device", server.Handle(c.restartDevice()))
	echo.POST("adb/start-server", server.Handle(c.startServer()))
	echo.POST("adb/stop-server", server.Handle(c.stopServer()))
}

func (c *ADBController) devices() server.HandlerFunc[server.Empty, domain.Response] {
	return func(ctx server.ApiContext, req *server.Empty) (*domain.Response, error) {
		output, err := c.adb.Devices(ctx.Context())
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *ADBController) stopDevice() server.HandlerFunc[adb.StopDevice, domain.Response] {
	return func(ctx server.ApiContext, req *adb.StopDevice) (*domain.Response, error) {
		output, err := c.adb.StopDevice(ctx.Context(), req.DeviceSerial)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *ADBController) restartDevice() server.HandlerFunc[adb.RestartDevice, domain.Response] {
	return func(ctx server.ApiContext, req *adb.RestartDevice) (*domain.Response, error) {
		output, err := c.adb.RestartDevice(ctx.Context(), req.DeviceSerial)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *ADBController) startServer() server.HandlerFunc[server.Empty, domain.Response] {
	return func(ctx server.ApiContext, req *server.Empty) (*domain.Response, error) {
		output, err := c.adb.StartServer(ctx.Context())
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *ADBController) stopServer() server.HandlerFunc[server.Empty, domain.Response] {
	return func(ctx server.ApiContext, req *server.Empty) (*domain.Response, error) {
		output, err := c.adb.StopServer(ctx.Context())
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

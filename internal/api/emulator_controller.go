package api

import (
	"android-cmd-server/internal/core/domain"
	"android-cmd-server/internal/core/domain/emulator"
	"android-cmd-server/internal/core/ports"
	"android-cmd-server/internal/infrastructure/server"
	"github.com/labstack/echo/v4"
)

type EmulatorController struct {
	emu ports.Emulator
}

func NewEmulatorController(
	echo *echo.Echo,
	emu ports.Emulator,
) {
	c := &EmulatorController{
		emu: emu,
	}
	echo.POST("emulator/start", server.Handle(c.start()))
}

func (c *EmulatorController) start() server.HandlerFunc[emulator.Start, domain.Response] {
	return func(ctx server.ApiContext, req *emulator.Start) (*domain.Response, error) {
		output, err := c.emu.Start(ctx.Context(), req.AVDName, req.Args...)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

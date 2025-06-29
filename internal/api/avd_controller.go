package api

import (
	"android-cmd-server/internal/core/domain"
	"android-cmd-server/internal/core/domain/avd"
	"android-cmd-server/internal/core/ports"
	"android-cmd-server/internal/infrastructure/server"
	"github.com/labstack/echo/v4"
)

type AVDController struct {
	avdManager ports.AVDManager
}

func NewAVDController(
	echo *echo.Echo,
	avdManager ports.AVDManager,
) {
	c := &AVDController{
		avdManager: avdManager,
	}
	echo.POST("avd-manager/create-avd", server.Handle(c.createAVD()))
	echo.GET("avd-manager/list-avd", server.Handle(c.listAVDs()))
	echo.DELETE("avd-manager/delete-avd", server.Handle(c.deleteAVD()))
}

func (c *AVDController) createAVD() server.HandlerFunc[avd.CreateAVD, domain.Response] {
	return func(ctx server.ApiContext, req *avd.CreateAVD) (*domain.Response, error) {
		args, err := c.createAVDManagerArgs(ctx)
		if err != nil {
			return nil, err
		}
		output, err := c.avdManager.CreateAVD(ctx.Context(), args, req.Name, req.PackagePath, req.Options)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *AVDController) listAVDs() server.HandlerFunc[server.Empty, domain.Response] {
	return func(ctx server.ApiContext, req *server.Empty) (*domain.Response, error) {
		args, err := c.createAVDManagerArgs(ctx)
		if err != nil {
			return nil, err
		}
		output, err := c.avdManager.ListAVDs(ctx.Context(), args)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *AVDController) deleteAVD() server.HandlerFunc[avd.DeleteAVD, domain.Response] {
	return func(ctx server.ApiContext, req *avd.DeleteAVD) (*domain.Response, error) {
		args, err := c.createAVDManagerArgs(ctx)
		if err != nil {
			return nil, err
		}
		output, err := c.avdManager.DeleteAVD(ctx.Context(), args, req.Name)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *AVDController) createAVDManagerArgs(ctx server.ApiContext) (*ports.AVDManagerArgs, error) {
	sdkVersion, err := domain.CheckSdkVersion(ctx.Header("sdk-version"))
	if err != nil {
		return nil, err
	}
	return &ports.AVDManagerArgs{
		SDKVersion: sdkVersion,
	}, nil
}

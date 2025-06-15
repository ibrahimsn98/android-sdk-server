package api

import (
	"android-cmd-server/internal/core/domain"
	"android-cmd-server/internal/core/domain/sdk"
	"android-cmd-server/internal/core/ports"
	"android-cmd-server/internal/infrastructure/server"
	"github.com/labstack/echo/v4"
)

type SDKController struct {
	sdkManager ports.SDKManager
}

func NewSDKController(
	echo *echo.Echo,
	sdkManager ports.SDKManager,
) {
	c := &SDKController{
		sdkManager: sdkManager,
	}
	echo.POST("sdk-manager/update-all", server.Handle(c.updateAll()))
	echo.GET("sdk-manager/list-packages", server.Handle(c.listPackages()))
	echo.POST("sdk-manager/install-packages", server.Handle(c.installPackages()))
}

func (c *SDKController) updateAll() server.HandlerFunc[server.Empty, domain.Response] {
	return func(ctx server.ApiContext, req *server.Empty) (*domain.Response, error) {
		args, err := c.createSDKManagerArgs(ctx)
		if err != nil {
			return nil, err
		}
		output, err := c.sdkManager.UpdateAll(ctx.Context(), args)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *SDKController) listPackages() server.HandlerFunc[server.Empty, domain.Response] {
	return func(ctx server.ApiContext, req *server.Empty) (*domain.Response, error) {
		args, err := c.createSDKManagerArgs(ctx)
		if err != nil {
			return nil, err
		}
		output, err := c.sdkManager.ListPackages(ctx.Context(), args)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *SDKController) installPackages() server.HandlerFunc[sdk.InstallPackages, domain.Response] {
	return func(ctx server.ApiContext, req *sdk.InstallPackages) (*domain.Response, error) {
		args, err := c.createSDKManagerArgs(ctx)
		if err != nil {
			return nil, err
		}
		output, err := c.sdkManager.InstallPackages(ctx.Context(), args, req.Packages)
		if err != nil {
			return nil, err
		}
		return &domain.Response{
			Stdout: output.Stdout,
			Stderr: output.Stderr,
		}, nil
	}
}

func (c *SDKController) createSDKManagerArgs(ctx server.ApiContext) (*ports.SDKManagerArgs, error) {
	sdkVersion, err := domain.CheckSdkVersion(ctx.Header("sdk-version"))
	if err != nil {
		return nil, err
	}
	return &ports.SDKManagerArgs{
		SDKVersion: sdkVersion,
	}, nil
}

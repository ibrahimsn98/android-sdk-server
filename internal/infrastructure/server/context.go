package server

import (
	"context"

	"github.com/labstack/echo/v4"
)

type apiContext struct {
	ec echo.Context
}

func NewApiContext(ec echo.Context) ApiContext {
	return &apiContext{ec: ec}
}

func (a *apiContext) Context() context.Context {
	return a.ec.Request().Context()
}

func (a *apiContext) Param(name string) string {
	return a.ec.Param(name)
}

func (a *apiContext) Set(key string, value any) {
	a.ec.Set(key, value)
}

func (a *apiContext) Get(key string) any {
	return a.ec.Get(key)
}

func (a *apiContext) Header(name string) string {
	return a.ec.Request().Header.Get(name)
}

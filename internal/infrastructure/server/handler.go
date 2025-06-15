package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type HandlerFunc[Req any, Res any] func(ctx ApiContext, req *Req) (*Res, error)

type ApiContext interface {
	Context() context.Context

	Param(name string) string

	Set(key string, value any)
	Get(key string) any
}

type Empty struct{}

func Handle[Req any, Res any](fn HandlerFunc[Req, Res]) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req Req

		reqType := reflect.TypeOf(req)
		if reqType.Kind() != reflect.Struct || reqType.NumField() > 0 {
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "invalid request body",
				})
			}

			if err := c.Validate(&req); err != nil {
				return c.JSON(http.StatusUnprocessableEntity, map[string]string{
					"error": err.Error(),
				})
			}
		}

		appCtx := NewApiContext(c)

		resp, err := fn(appCtx, &req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

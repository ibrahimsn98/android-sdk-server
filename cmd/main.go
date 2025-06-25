package main

import (
	"android-cmd-server/internal/application"
	"android-cmd-server/internal/infrastructure/waiter"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func run() (err error) {
	m := module{}
	m.api = echo.New()
	m.api.Validator = &CustomValidator{validator: validator.New()}
	m.api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${method} ${uri} (${status})\n",
	}))

	m.waiter = waiter.New(waiter.CatchSignals())
	m.app = &application.Application{}
	err = m.app.Startup(m.waiter.Context(), &m)
	if err != nil {
		return
	}

	m.waiter.Add(m.waitForApi)
	return m.waiter.Wait()
}

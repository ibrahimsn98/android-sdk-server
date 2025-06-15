package application

import (
	"android-cmd-server/internal/infrastructure/waiter"
	"github.com/labstack/echo/v4"
)

type Module interface {
	Api() *echo.Echo
	Waiter() waiter.Waiter
}

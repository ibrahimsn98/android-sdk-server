package application

import (
	"android-cmd-server/internal/infrastructure/logger"
	"android-cmd-server/internal/infrastructure/waiter"
	"github.com/labstack/echo/v4"
)

type Module interface {
	Logger() *logger.Logger
	Api() *echo.Echo
	Waiter() waiter.Waiter
}

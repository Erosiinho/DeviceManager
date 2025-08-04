package routes

import (
	"github.com/labstack/echo/v4"
	"wakeapi/internal/controllers"
)

func Register(apiGrp *echo.Group, controllerCtx *controllers.ControllerCtx) {
	deviceGroup := apiGrp.Group("/devices")
	deviceGroup.GET("", controllerCtx.GetAllDevices)
	deviceGroup.GET("/:id/start", controllerCtx.StartDevice)
	deviceGroup.GET("/:id/stop", controllerCtx.StopDevice)
	deviceGroup.GET("/:id/ping", controllerCtx.PingDevice)
}

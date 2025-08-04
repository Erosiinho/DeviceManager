package routes

import (
	"github.com/labstack/echo/v4"
	"wakeapi/internal/controllers"
)

func Register(apiGrp *echo.Group, controllerCtx *controllers.ControllerCtx) {
	deviceGroup := apiGrp.Group("/devices")
	deviceGroup.GET("", controllerCtx.GetAllDevices)
	deviceGroup.POST("/:id/start", controllerCtx.StartDevice)
	deviceGroup.POST("/:id/stop", controllerCtx.StopDevice)
	deviceGroup.GET("/:id/ping", controllerCtx.PingDevice)
}

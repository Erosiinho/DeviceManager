package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"wakeapi/internal/commands"
	"wakeapi/pkg/response"
)

func (ctx *ControllerCtx) GetAllDevices(c echo.Context) error {
	return c.JSON(http.StatusOK, ctx.DeviceManager.GetAllDevices())
}

func (ctx *ControllerCtx) StartDevice(c echo.Context) error {
	// retrieve params
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, response.NewApiErrorResponse(
			http.StatusBadRequest,
			"Missing id params",
			"",
		))
	}

	// get specific device
	device, found := ctx.DeviceManager.GetDevice(id)
	if !found {
		return c.JSON(http.StatusNotFound, response.NewApiErrorResponse(
			http.StatusNotFound,
			"Device not found",
			"",
		))
	}

	// ping to know if already up
	isUp := commands.Ping(device.IPAddress)
	if isUp {
		return c.JSON(http.StatusConflict, response.NewApiErrorResponse(
			http.StatusConflict,
			"Device already connected",
			"",
		))
	}

	// wol on device
	err := commands.WakeOnLan(device.MacAddress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewApiErrorResponse(
			http.StatusInternalServerError,
			"Failed to send Wake-on-LAN",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "WoL command sent successfully"})
}

func (ctx *ControllerCtx) StopDevice(c echo.Context) error {
	// retrieve params
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, response.NewApiErrorResponse(
			http.StatusBadRequest,
			"Missing id params",
			"",
		))
	}

	// get specific device
	device, found := ctx.DeviceManager.GetDevice(id)
	if !found {
		return c.JSON(http.StatusNotFound, response.NewApiErrorResponse(
			http.StatusNotFound,
			"Device not found",
			"",
		))
	}

	// shutdown on device
	err := commands.ShutdownViaSSH(
		device.IPAddress,
		device.SSHUsername,
		device.SSHPassword,
		strconv.Itoa(device.SSHPort),
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.NewApiErrorResponse(
			http.StatusInternalServerError,
			"Failed to send shutdown command",
			err.Error(),
		))
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Shutdown command sent successfully"})
}

func (ctx *ControllerCtx) PingDevice(c echo.Context) error {
	// retrieve params
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, response.NewApiErrorResponse(
			http.StatusBadRequest,
			"Missing id params",
			"",
		))
	}

	// get specific device
	device, found := ctx.DeviceManager.GetDevice(id)
	if !found {
		return c.JSON(http.StatusNotFound, response.NewApiErrorResponse(
			http.StatusNotFound,
			"Device not found",
			"",
		))
	}

	// ping
	isUp := commands.Ping(device.IPAddress)
	status := "offline"
	if isUp {
		status = "online"
	}
	return c.JSON(http.StatusOK, map[string]string{"status": status})
}

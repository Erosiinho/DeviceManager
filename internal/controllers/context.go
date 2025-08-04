package controllers

import "wakeapi/internal/config"

type ControllerCtx struct {
	DeviceManager *config.DeviceManager
}

func NewContext(deviceManager *config.DeviceManager) *ControllerCtx {
	return &ControllerCtx{
		DeviceManager: deviceManager,
	}
}

package config

import "wakeapi/pkg/response"

type Config struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	ID          string `json:"id"`
	MacAddress  string `json:"mac_address"`
	IPAddress   string `json:"ip_address"`
	SSHPort     int    `json:"ssh_port"`
	SSHUsername string `json:"ssh_username"`
	SSHPassword string `json:"ssh_password"`
}

type IDeviceManager interface {
	Load(string) error
	GetPublicDevice(string) (*response.DevicePublic, bool)
	GetDevice(string) (*Device, bool)
	GetAllDevices() map[string]*response.DevicePublic
	ListDeviceIDs() []string
	PrintDevices()
}

type DeviceManager struct {
	devices map[string]*Device
}

func NewDeviceManager() *DeviceManager {
	return &DeviceManager{
		devices: make(map[string]*Device),
	}
}

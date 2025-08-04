package config

import (
	"encoding/json"
	"fmt"
	"os"
	"wakeapi/pkg/response"
)

func (dm *DeviceManager) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error when reading file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("erreur when parsing JSON: %w", err)
	}

	for i := range config.Devices {
		device := &config.Devices[i]
		dm.devices[device.ID] = device
	}
	return nil
}

func (dm *DeviceManager) GetDevice(id string) (*Device, bool) {
	device, exists := dm.devices[id]
	return device, exists
}

func (dm *DeviceManager) GetPublicDevice(id string) (*response.DevicePublic, bool) {
	device, exists := dm.devices[id]
	if !exists {
		return nil, false
	}
	return &response.DevicePublic{
		ID:          device.ID,
		MacAddress:  device.MacAddress,
		IPAddress:   device.IPAddress,
		SSHPort:     device.SSHPort,
		SSHUsername: device.SSHUsername,
	}, true
}

func (dm *DeviceManager) GetAllDevices() map[string]*response.DevicePublic {
	devices := make(map[string]*response.DevicePublic, len(dm.devices))
	for _, d := range dm.devices {
		devices[d.ID] = &response.DevicePublic{
			ID:          d.ID,
			MacAddress:  d.MacAddress,
			IPAddress:   d.IPAddress,
			SSHPort:     d.SSHPort,
			SSHUsername: d.SSHUsername,
		}
	}
	return devices
}

func (dm *DeviceManager) ListDeviceIDs() []string {
	ids := make([]string, 0, len(dm.devices))
	for id := range dm.devices {
		ids = append(ids, id)
	}
	return ids
}

func (dm *DeviceManager) PrintDevices() {
	if len(dm.devices) == 0 {
		fmt.Println("No device")
		return
	}
	fmt.Printf("=== Devices (%d) ===\n", len(dm.devices))
	for id, device := range dm.devices {
		fmt.Printf("ID: %s\n", id)
		fmt.Printf("  MAC: %s\n", device.MacAddress)
		fmt.Printf("  IP: %s:%d\n", device.IPAddress, device.SSHPort)
		fmt.Printf("  SSH: %s@%s\n", device.SSHUsername, device.IPAddress)
		fmt.Println("  ---")
	}
}

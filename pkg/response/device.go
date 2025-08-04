package response

type DevicePublic struct {
	ID          string `json:"id"`
	MacAddress  string `json:"mac_address"`
	IPAddress   string `json:"ip_address"`
	SSHPort     int    `json:"ssh_port"`
	SSHUsername string `json:"ssh_username"`
}

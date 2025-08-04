package commands

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"os/exec"
	"time"
)

func ShutdownViaSSH(ip, user, password string, port string) error {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", ip, port), config)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	cmd := "shutdown /s /t 0"
	return session.Run(cmd)
}

func WakeOnLan(macAddr string) error {
	hwAddr, err := net.ParseMAC(macAddr)
	if err != nil {
		return err
	}

	packet := make([]byte, 102)
	for i := 0; i < 6; i++ {
		packet[i] = 0xFF
	}
	for i := 6; i < 102; i += 6 {
		copy(packet[i:], hwAddr)
	}

	addr := &net.UDPAddr{IP: net.IPv4bcast, Port: 9}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	return err
}

func Ping(ip string) bool {
	_, err := exec.Command("ping", "-c", "1", "-W", "1", ip).Output()
	fmt.Printf("error when pinging: %v\n", err)
	return err == nil
}

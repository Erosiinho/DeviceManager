# WakeAPI 🔌💻

**WakeAPI** is a minimalist REST API to control the power state of your machines using their unique identifier.
It allows you to **power on** (via Wake-on-LAN), **shut down** (via SSH), and **check the status** (via ping) of a configured device.

---

## 🚀 Features

* 🔸 Power on a machine using Wake-on-LAN
* 🔸 Gracefully shut down a machine via SSH
* 🔸 Check if a machine is up using a simple ping
* 🔸 List all configured machines

---

## 🧀 Configuration

Devices are listed in a JSON config file using the following format:

```json
{
  "devices": [
    {
      "id": "work-pc",
      "mac_address": "AA:BB:CC:DD:EE:FF",
      "ip_address": "192.168.1.100",
      "ssh_port": 22,
      "ssh_username": "user",
      "ssh_password": "password"
    }
  ]
}
```

The path to this config file is defined via the `DEFAULT_CONFIG_PATH` environment variable.

---

## 🔗 API Routes

All routes are prefixed with `/api/v1`.

| Method | Route                       | Description                          |
|--------| --------------------------- | ------------------------------------ |
| GET    | `/api/v1/devices`           | Returns the list of devices          |
| GET    | `/api/v1/devices/:id/start` | Powers on the device via Wake-on-LAN |
| GET    | `/api/v1/devices/:id/stop`  | Shuts down the device via SSH        |
| GET    | `/api/v1/devices/:id/ping`  | Checks if the device is online       |

> ⚠️ The `id` field in the URL must match the one defined in the JSON config.

---

## ⚙️ Required Environment Variables

| Variable              | Description                         |
| --------------------- | ----------------------------------- |
| `DEFAULT_CONFIG_PATH` | Path to the JSON configuration file |
| `HTTP_PORT`           | HTTP port the API will listen on    |

---

## 🐳 Docker

The API is dockerized and ready to use. Just define the two required environment variables, mount your config, and run:

```yaml
services:
  wakeapi:
    image: wakeapi:latest
    environment:
      - DEFAULT_CONFIG_PATH=/app/config.json
      - HTTP_PORT=8089
    volumes:
      - ./config.json:/app/config.json
    network_mode: "host"
```

Then start the service:

```bash
docker compose up -d
```
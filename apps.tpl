{
    "default": {
        "rfid": {
            "name": "rfid",
            "label": "RFID",
            "description": "RFID Apps",
            "icon": "rfid",
            "image": "ghcr.io/peramic/apps:latest",
            "trust": true,
            "state": "STARTED",
            "devices": [
                {
                    "path": "/dev/ttyNUR0"
                },
                {
                    "path": "/dev/mica_gpio"
                }
            ],
            "mounts": [
                {
                    "source": "conf",
                    "destination": "/opt/havis-apps/conf"
                },
                {
                    "source": "depot",
                    "destination": "/opt/depot/files"
                }
            ]
        },
        "gpio": {
            "name": "gpio",
            "label": "GPIO",
            "description": "GPIO Monitor",
            "icon": "gpio",
            "image": "ghcr.io/peramic/gpio:latest",
            "trust": true,
            "state": "STARTED",
            "devices": [
                {
                    "path": "/dev/mica_gpio"
                }
            ],
            "mounts": [
                {
                    "source": "conf",
                    "destination": "/opt/peramic-app/conf"
                }
            ]
        }
    },
    "system": {
        "runtime": {
            "name": "runtime",
            "label": "Runtime",
            "description": "Runtime",
            "icon": "settings_applications",
            "image": "ghcr.io/peramic/runtime:${VARIANT}-latest",
            "trust": true,
            "state": "UPGRADED"
        },
        "market": {
            "name": "market",
            "label": "Market",
            "description": "Market",
            "icon": "shopping_basket",
            "image": "ghcr.io/peramic/market:${VARIANT}-latest",
            "trust": true,
            "state": "STARTED"
        },
        "auth": {
            "name": "auth",
            "label": "Auth",
            "description": "Auth",
            "icon": "login",
            "image": "ghcr.io/peramic/auth:latest",
            "trust": true,
            "state": "STARTED",
            "mounts": [
                {
                    "source": "conf",
                    "destination": "/opt/peramic-app/conf"
                }
            ]
        },
        "httpd": {
            "name": "httpd",
            "label": "Httpd",
            "description": "Httpd",
            "icon": "http",
            "image": "ghcr.io/peramic/httpd:latest",
            "trust": true,
            "namespaces": [{"type" : "network"}],
            "state": "STARTED"
        },
        "log": {
            "name": "log",
            "label": "Log",
            "description": "Logging",
            "icon": "opcua_gw",
            "image": "ghcr.io/peramic/log:latest",
            "trust": true,
            "state": "STARTED"
        },
        "mqtt": {
            "name": "mqtt",
            "label": "MQTT",
            "description": "MQTT broker",
            "icon": "wifi_tethering",
            "image": "ghcr.io/peramic/mqtt:latest",
            "trust": true,
            "mounts": [
                {
                    "source": "conf",
                    "destination": "/etc/mosquitto/conf.d"
                }
            ],
            "state": "STARTED"
        },
        "vpn": {
            "name": "vpn",
            "label": "VPN",
            "description": "OpenVPN",
            "icon": "verified_user",
            "image": "ghcr.io/peramic/vpn:latest",
            "trust": true,
            "namespaces": [{"type" : "network"}],
            "devices": [
                {
                    "path": "/dev/net/tun"
                }
            ],
            "mounts": [
                {
                    "source": "conf",
                    "destination": "/opt/peramic-app/conf"
                }
            ],
            "capabilities": [
                "CAP_SYS_ADMIN",
                "CAP_NET_ADMIN"
            ],
            "state": "STARTED"
        }
    }
}

# Huawei Modem Go Library

The Huawei Modem Go library provides a simple and convenient way to interact with Huawei LTE WiFi routers (MiFi) through a series of HTTP API endpoints. This library enables users to perform various operations, such as logging in, retrieving device status, sending SMS messages, and more.

## Features

- **Login and Logout**: Authenticate and manage sessions with the modem.
- **Device Status**: Retrieve comprehensive status information, including signal strength, battery level, and network status.
- **SMS Management**: Send, read, and delete SMS messages.
- **Device Information**: Get detailed information about the device.
- **Network and Signal Information**: Obtain current network type, signal strength, and more.
- **Control Operations**: Reboot the device and manage various settings.

## Installation

To install the Huawei Modem Go library, use the following command:

```sh
go get github.com/lagarciag/huaweimoden
```

## Usage
### Basic Usage Example
Here's a basic example of how to use the library to log in to the modem, get device status, and send an SMS message.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/lagarciag/huaweimoden/device"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Initialize the device
	modem := device.Device{
		l:        sugar,
		client:   &http.Client{},
		deviceIP: "192.168.8.1",
		user:     "admin",
		password: "yourpassword",
	}

	// Login to the modem
	err := modem.Login()
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	defer modem.Logout()

	// Get device status
	status, err := modem.GetDeviceStatus()
	if err != nil {
		log.Fatalf("Failed to get device status: %v", err)
	}
	fmt.Printf("Device Status: %+v\n", status)

	// Send an SMS
	err = modem.SendSMS("1234567890", "Hello from Go!")
	if err != nil {
		log.Fatalf("Failed to send SMS: %v", err)
	}
	fmt.Println("SMS sent successfully")
}
```



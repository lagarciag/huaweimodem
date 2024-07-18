package huaweimodem

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceStatus represents the status of the device, including signal strength and battery level.
type DeviceStatus struct {
	XMLName        xml.Name `xml:"response"`       // XMLName is the XML element name for the response.
	SignalStrength string   `xml:"SignalIcon"`     // SignalStrength is the signal strength of the device.
	BatteryLevel   string   `xml:"BatteryPercent"` // BatteryLevel is the battery level of the device.
	DeviceName     string   `xml:"DeviceName"`     // DeviceName is the name of the device.
	SerialNumber   string   `xml:"SerialNumber"`   // SerialNumber is the serial number of the device.
	// Add other fields as needed
}

func (d *Device) GetDeviceStatus() (*DeviceStatus, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf(UrlDeviceStatus, d.deviceIP), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create status request: %w", err)
	}
	req.Header.Set("Cookie", d.sessionID)
	req.Header.Set("__RequestVerificationToken", d.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send status request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read status response: %w", err)
	}

	var status DeviceStatus
	if err := xml.Unmarshal(body, &status); err != nil {
		return nil, fmt.Errorf("failed to unmarshal status response: %w", err)
	}

	return &status, nil
}

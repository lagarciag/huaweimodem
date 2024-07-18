package huaweimodem

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceStatus represents the status of the device, including signal strength and battery level.
type DeviceStatus struct {
	XMLName              xml.Name `xml:"response"`             // XMLName is the XML element name for the response.
	WifiConnectionStatus int      `xml:"WifiConnectionStatus"` // WifiConnectionStatus indicates the current Wi-Fi connection status.
	SignalStrength       int      `xml:"SignalStrength"`       // SignalStrength represents the strength of the signal.
	SignalIcon           int      `xml:"SignalIcon"`           // SignalIcon provides a visual representation of signal strength.
	CurrentNetworkType   int      `xml:"CurrentNetworkType"`   // CurrentNetworkType indicates the type of the current network.
	CurrentServiceDomain int      `xml:"CurrentServiceDomain"` // CurrentServiceDomain indicates the current service domain.
	RoamingStatus        int      `xml:"RoamingStatus"`        // RoamingStatus indicates if the device is roaming.
	BatteryStatus        int      `xml:"BatteryStatus"`        // BatteryStatus indicates the status of the battery.
	BatteryLevel         int      `xml:"BatteryLevel"`         // BatteryLevel represents the battery level.
	BatteryPercent       int      `xml:"BatteryPercent"`       // BatteryPercent indicates the battery percentage.
	SimlockStatus        int      `xml:"simlockStatus"`        // SimlockStatus indicates if the SIM is locked.
	WanIPAddress         string   `xml:"WanIPAddress"`         // WanIPAddress is the WAN IP address of the device.
	WanIPv6Address       string   `xml:"WanIPv6Address"`       // WanIPv6Address is the WAN IPv6 address of the device.
	PrimaryDns           string   `xml:"PrimaryDns"`           // PrimaryDns is the primary DNS address.
	SecondaryDns         string   `xml:"SecondaryDns"`         // SecondaryDns is the secondary DNS address.
	PrimaryIPv6Dns       string   `xml:"PrimaryIPv6Dns"`       // PrimaryIPv6Dns is the primary IPv6 DNS address.
	SecondaryIPv6Dns     string   `xml:"SecondaryIPv6Dns"`     // SecondaryIPv6Dns is the secondary IPv6 DNS address.
	CurrentWifiUser      int      `xml:"CurrentWifiUser"`      // CurrentWifiUser is the number of current Wi-Fi users.
	TotalWifiUser        int      `xml:"TotalWifiUser"`        // TotalWifiUser is the total number of Wi-Fi users.
	CurrentTotalWifiUser int      `xml:"currenttotalwifiuser"` // CurrentTotalWifiUser is the current total number of Wi-Fi users.
	ServiceStatus        int      `xml:"ServiceStatus"`        // ServiceStatus indicates the status of the service.
	SimStatus            int      `xml:"SimStatus"`            // SimStatus indicates the status of the SIM card.
	WifiStatus           int      `xml:"WifiStatus"`           // WifiStatus indicates the status of the Wi-Fi.
	CurrentNetworkTypeEx int      `xml:"CurrentNetworkTypeEx"` // CurrentNetworkTypeEx indicates the extended current network type.
	WanPolicy            int      `xml:"WanPolicy"`            // WanPolicy indicates the WAN policy.
	MaxSignal            int      `xml:"maxsignal"`            // MaxSignal indicates the maximum signal strength.
	WifiIndoorOnly       int      `xml:"wifiindooronly"`       // WifiIndoorOnly indicates if the Wi-Fi is for indoor use only.
	WifiFrequence        int      `xml:"wififrequence"`        // WifiFrequence indicates the frequency of the Wi-Fi.
	Classify             string   `xml:"classify"`             // Classify indicates the classification of the device.
	FlyMode              int      `xml:"flymode"`              // FlyMode indicates if the device is in flight mode.
	CellRoam             int      `xml:"cellroam"`             // CellRoam indicates if the device is in cell roaming mode.
}

// GetWifiConnectionStatus returns the Wi-Fi connection status.
func (d *DeviceStatus) GetWifiConnectionStatus() int {
	return d.WifiConnectionStatus
}

// GetSignalStrength returns the signal strength.
func (d *DeviceStatus) GetSignalStrength() int {
	return d.SignalStrength
}

// GetSignalIcon returns the signal icon.
func (d *DeviceStatus) GetSignalIcon() int {
	return d.SignalIcon
}

// GetCurrentNetworkType returns the current network type.
func (d *DeviceStatus) GetCurrentNetworkType() int {
	return d.CurrentNetworkType
}

// GetCurrentServiceDomain returns the current service domain.
func (d *DeviceStatus) GetCurrentServiceDomain() int {
	return d.CurrentServiceDomain
}

// GetRoamingStatus returns the roaming status.
func (d *DeviceStatus) GetRoamingStatus() int {
	return d.RoamingStatus
}

// GetBatteryStatus returns the battery status.
func (d *DeviceStatus) GetBatteryStatus() int {
	return d.BatteryStatus
}

// GetBatteryLevel returns the battery level.
func (d *DeviceStatus) GetBatteryLevel() int {
	return d.BatteryLevel
}

// GetBatteryPercent returns the battery percentage.
func (d *DeviceStatus) GetBatteryPercent() int {
	return d.BatteryPercent
}

// GetSimlockStatus returns the SIM lock status.
func (d *DeviceStatus) GetSimlockStatus() int {
	return d.SimlockStatus
}

// GetWanIPAddress returns the WAN IP address.
func (d *DeviceStatus) GetWanIPAddress() string {
	return d.WanIPAddress
}

// GetWanIPv6Address returns the WAN IPv6 address.
func (d *DeviceStatus) GetWanIPv6Address() string {
	return d.WanIPv6Address
}

// GetPrimaryDns returns the primary DNS address.
func (d *DeviceStatus) GetPrimaryDns() string {
	return d.PrimaryDns
}

// GetSecondaryDns returns the secondary DNS address.
func (d *DeviceStatus) GetSecondaryDns() string {
	return d.SecondaryDns
}

// GetPrimaryIPv6Dns returns the primary IPv6 DNS address.
func (d *DeviceStatus) GetPrimaryIPv6Dns() string {
	return d.PrimaryIPv6Dns
}

// GetSecondaryIPv6Dns returns the secondary IPv6 DNS address.
func (d *DeviceStatus) GetSecondaryIPv6Dns() string {
	return d.SecondaryIPv6Dns
}

// GetCurrentWifiUser returns the number of current Wi-Fi users.
func (d *DeviceStatus) GetCurrentWifiUser() int {
	return d.CurrentWifiUser
}

// GetTotalWifiUser returns the total number of Wi-Fi users.
func (d *DeviceStatus) GetTotalWifiUser() int {
	return d.TotalWifiUser
}

// GetCurrentTotalWifiUser returns the current total number of Wi-Fi users.
func (d *DeviceStatus) GetCurrentTotalWifiUser() int {
	return d.CurrentTotalWifiUser
}

// GetServiceStatus returns the service status.
func (d *DeviceStatus) GetServiceStatus() int {
	return d.ServiceStatus
}

// GetSimStatus returns the SIM status.
func (d *DeviceStatus) GetSimStatus() int {
	return d.SimStatus
}

// GetWifiStatus returns the Wi-Fi status.
func (d *DeviceStatus) GetWifiStatus() int {
	return d.WifiStatus
}

// GetCurrentNetworkTypeEx returns the extended current network type.
func (d *DeviceStatus) GetCurrentNetworkTypeEx() int {
	return d.CurrentNetworkTypeEx
}

// GetWanPolicy returns the WAN policy.
func (d *DeviceStatus) GetWanPolicy() int {
	return d.WanPolicy
}

// GetMaxSignal returns the maximum signal strength.
func (d *DeviceStatus) GetMaxSignal() int {
	return d.MaxSignal
}

// GetWifiIndoorOnly returns whether the Wi-Fi is for indoor use only.
func (d *DeviceStatus) GetWifiIndoorOnly() int {
	return d.WifiIndoorOnly
}

// GetWifiFrequence returns the Wi-Fi frequency.
func (d *DeviceStatus) GetWifiFrequence() int {
	return d.WifiFrequence
}

// GetClassify returns the classification of the device.
func (d *DeviceStatus) GetClassify() string {
	return d.Classify
}

// GetFlyMode returns whether the device is in flight mode.
func (d *DeviceStatus) GetFlyMode() int {
	return d.FlyMode
}

// GetCellRoam returns whether the device is in cell roaming mode.
func (d *DeviceStatus) GetCellRoam() int {
	return d.CellRoam
}

// DeviceStatus retrieves the status of the device, including signal strength and battery level.
// It first checks if the user is logged in by verifying the sessionID.
// If not logged in, it returns an error.
// Then it refreshes the session and token information, and sends a request to the device status endpoint.
// The response is parsed and unmarshalled into a DeviceStatus struct, which is returned.
//
// Returns:
//   - A pointer to the DeviceStatus struct containing the device status.
//   - An error if any step in the process fails.
func (d *Device) DeviceStatus() (*DeviceStatus, error) {
	if d.sessionID == "" {
		return nil, fmt.Errorf("you must login first")
	}

	err := d.getSesTokInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get SesTokInfo: %w", err)
	}

	client := d.client
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

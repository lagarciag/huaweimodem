package huaweimodem

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceStatus represents the status of the device, including signal strength and battery level.
type DeviceStatus struct {
	XMLName              xml.Name `xml:"response"` // XMLName is the XML element name for the response.
	WifiConnectionStatus int      `xml:"WifiConnectionStatus"`
	SignalStrength       int      `xml:"SignalStrength"`
	SignalIcon           int      `xml:"SignalIcon"`
	CurrentNetworkType   int      `xml:"CurrentNetworkType"`
	CurrentServiceDomain int      `xml:"CurrentServiceDomain"`
	RoamingStatus        int      `xml:"RoamingStatus"`
	BatteryStatus        int      `xml:"BatteryStatus"`
	BatteryLevel         int      `xml:"BatteryLevel"`
	BatteryPercent       int      `xml:"BatteryPercent"`
	SimlockStatus        int      `xml:"simlockStatus"`
	WanIPAddress         string   `xml:"WanIPAddress"`
	WanIPv6Address       string   `xml:"WanIPv6Address"`
	PrimaryDns           string   `xml:"PrimaryDns"`
	SecondaryDns         string   `xml:"SecondaryDns"`
	PrimaryIPv6Dns       string   `xml:"PrimaryIPv6Dns"`
	SecondaryIPv6Dns     string   `xml:"SecondaryIPv6Dns"`
	CurrentWifiUser      int      `xml:"CurrentWifiUser"`
	TotalWifiUser        int      `xml:"TotalWifiUser"`
	CurrentTotalWifiUser int      `xml:"currenttotalwifiuser"`
	ServiceStatus        int      `xml:"ServiceStatus"`
	SimStatus            int      `xml:"SimStatus"`
	WifiStatus           int      `xml:"WifiStatus"`
	CurrentNetworkTypeEx int      `xml:"CurrentNetworkTypeEx"`
	WanPolicy            int      `xml:"WanPolicy"`
	MaxSignal            int      `xml:"maxsignal"`
	WifiIndoorOnly       int      `xml:"wifiindooronly"`
	WifiFrequence        int      `xml:"wififrequence"`
	Classify             string   `xml:"classify"`
	FlyMode              int      `xml:"flymode"`
	CellRoam             int      `xml:"cellroam"`
	// Add other fields as needed
}

func (d *DeviceStatus) GetWifiConnectionStatus() int {
	return d.WifiConnectionStatus
}

func (d *DeviceStatus) GetSignalStrength() int {
	return d.SignalStrength
}

func (d *DeviceStatus) GetSignalIcon() int {
	return d.SignalIcon
}

func (d *DeviceStatus) GetCurrentNetworkType() int {
	return d.CurrentNetworkType
}

func (d *DeviceStatus) GetCurrentServiceDomain() int {
	return d.CurrentServiceDomain
}

func (d *DeviceStatus) GetRoamingStatus() int {
	return d.RoamingStatus
}

func (d *DeviceStatus) GetBatteryStatus() int {
	return d.BatteryStatus
}

func (d *DeviceStatus) GetBatteryLevel() int {
	return d.BatteryLevel
}

func (d *DeviceStatus) GetBatteryPercent() int {
	return d.BatteryPercent
}

func (d *DeviceStatus) GetSimlockStatus() int {
	return d.SimlockStatus
}

func (d *DeviceStatus) GetWanIPAddress() string {
	return d.WanIPAddress
}

func (d *DeviceStatus) GetWanIPv6Address() string {
	return d.WanIPv6Address
}

func (d *DeviceStatus) GetPrimaryDns() string {
	return d.PrimaryDns
}

func (d *DeviceStatus) GetSecondaryDns() string {
	return d.SecondaryDns
}

func (d *DeviceStatus) GetPrimaryIPv6Dns() string {
	return d.PrimaryIPv6Dns
}

func (d *DeviceStatus) GetSecondaryIPv6Dns() string {
	return d.SecondaryIPv6Dns
}

func (d *DeviceStatus) GetCurrentWifiUser() int {
	return d.CurrentWifiUser
}

func (d *DeviceStatus) GetTotalWifiUser() int {
	return d.TotalWifiUser
}

func (d *DeviceStatus) GetCurrentTotalWifiUser() int {
	return d.CurrentTotalWifiUser
}

func (d *DeviceStatus) GetServiceStatus() int {
	return d.ServiceStatus
}

func (d *DeviceStatus) GetSimStatus() int {
	return d.SimStatus
}

func (d *DeviceStatus) GetWifiStatus() int {
	return d.WifiStatus
}

func (d *DeviceStatus) GetCurrentNetworkTypeEx() int {
	return d.CurrentNetworkTypeEx
}

func (d *DeviceStatus) GetWanPolicy() int {
	return d.WanPolicy
}

func (d *DeviceStatus) GetMaxSignal() int {
	return d.MaxSignal
}

func (d *DeviceStatus) GetWifiIndoorOnly() int {
	return d.WifiIndoorOnly
}

func (d *DeviceStatus) GetWifiFrequence() int {
	return d.WifiFrequence
}

func (d *DeviceStatus) GetClassify() string {
	return d.Classify
}

func (d *DeviceStatus) GetFlyMode() int {
	return d.FlyMode
}

func (d *DeviceStatus) GetCellRoam() int {
	return d.CellRoam
}

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

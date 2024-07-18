package huaweimodem

import (
	"encoding/xml"
	"go.uber.org/zap"
	"net/http"
	"net/http/cookiejar"
)

// Constants for content type and URLs
const (
	httpContentType = "application/x-www-form-urlencoded; charset=UTF-8"
)

// List of API endpoints discovered so far, not all of them are implemented
const (
	// UrlLogin is the endpoint to authenticate and login to the Huawei modem.
	UrlLogin = "http://%s/api/user/login"

	// UrlLogout is the endpoint to log out from the Huawei modem.
	UrlLogout = "http://%s/api/user/logout"

	// UrlSesTokInfo is the endpoint to get session and token information.
	UrlSesTokInfo = "http://%s/api/webserver/SesTokInfo"

	// UrlDeviceStatus is the endpoint to retrieve the current status of the device, including signal strength, battery level, and network status.
	UrlDeviceStatus = "http://%s/api/monitoring/status"

	// UrlDeviceInformation is the endpoint to retrieve detailed information about the device, such as device name, serial number, and IMEI.
	UrlDeviceInformation = "http://%s/api/device/information"

	// UrlSMSList is the endpoint to get a list of SMS messages.
	UrlSMSList = "http://%s/api/sms/sms-list"

	// UrlSendSMS is the endpoint to send an SMS message.
	UrlSendSMS = "http://%s/api/sms/send-sms"

	// UrlDeleteSMS is the endpoint to delete an SMS message.
	UrlDeleteSMS = "http://%s/api/sms/delete-sms"

	// UrlSetSMSRead is the endpoint to mark an SMS message as read.
	UrlSetSMSRead = "http://%s/api/sms/set-read"

	// UrlCurrentPLMN is the endpoint to get information about the current network provider (PLMN).
	UrlCurrentPLMN = "http://%s/api/net/current-plmn"

	// UrlConvergedStatus is the endpoint to get the converged status of the modem, including signal and network status.
	UrlConvergedStatus = "http://%s/api/monitoring/converged-status"

	// UrlMonthStatistics is the endpoint to retrieve monthly statistics about data usage.
	UrlMonthStatistics = "http://%s/api/monitoring/month_statistics"

	// UrlStartDate is the endpoint to get the start date of the data monitoring period.
	UrlStartDate = "http://%s/api/monitoring/start_date"

	// UrlHostList is the endpoint to retrieve a list of devices currently connected to the modem via Wi-Fi.
	UrlHostList = "http://%s/api/wlan/host-list"

	// UrlCheckNotifications is the endpoint to check for new notifications, such as unread SMS messages.
	UrlCheckNotifications = "http://%s/api/monitoring/check-notifications"

	// UrlControlDevice is the endpoint to control the device, such as rebooting it.
	UrlControlDevice = "http://%s/api/device/control"

	// UrlGetSysSettings is the endpoint to retrieve system settings.
	UrlGetSysSettings = "http://%s/api/settings/get-sys"

	// UrlSetSysSettings is the endpoint to set or update system settings.
	UrlSetSysSettings = "http://%s/api/settings/set-sys"

	// UrlDHCPSettings is the endpoint to retrieve or update DHCP settings.
	UrlDHCPSettings = "http://%s/api/dhcp/settings"

	// UrlFirewallSwitch is the endpoint to enable or disable the firewall.
	UrlFirewallSwitch = "http://%s/api/security/firewall-switch"
)

// ErrorResponse represents a generic error response from the API.
type ErrorResponse struct {
	XMLName   xml.Name `xml:"error"`   // XMLName is the XML element name for the error.
	ErrorCode string   `xml:"code"`    // ErrorCode is the code returned by the API, indicating the type of error.
	Message   string   `xml:"message"` // Message is the error message returned by the API.
}

// Device represents the device information and authentication details.
type Device struct {
	l            *zap.SugaredLogger // Logger instance for logging.
	client       *http.Client       // HTTP client for making requests.
	sessionID    string             // Session ID for the current session.
	token        string             // Token for authentication.
	deviceIP     string             // IP address of the device.
	user         string             // Username for authentication.
	password     string             // Password for authentication.
	deviceStatus *DeviceStatus
}

// DeviceIP returns the IP address of the device.
func (d *Device) DeviceIP() string {
	return d.deviceIP
}

// User returns the username used to authenticate with the device.
func (d *Device) User() string {
	return d.user
}

// NewDevice creates a new instance of Device with the specified logger, device IP, username, and password.
// It initializes the Device struct and sets up an HTTP client with a cookie jar to manage session cookies.
//
// Parameters:
//   - l: A SugaredLogger instance from the zap logging package.
//   - deviceIP: The IP address of the device to connect to.
//   - user: The username for authentication.
//   - password: The password for authentication.
//
// Returns:
//   - A pointer to the initialized Device instance.
//   - An error if the device could not be created.
func NewDevice(l *zap.SugaredLogger, deviceIP, user, password string) (*Device, error) {
	d := Device{
		l:        l,
		deviceIP: deviceIP,
		user:     user,
	}

	// Hash and encode the password
	d.password = d.hashAndEncodePassword(password)

	// Initialize HTTP client with a cookie jar
	client := &http.Client{
		Jar: nil, // cookie jar to store and manage cookies
	}
	client.Jar, _ = cookiejar.New(nil)

	d.client = client

	return &d, nil
}

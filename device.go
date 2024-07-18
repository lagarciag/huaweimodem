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

const (
	UrlLogin        = "http://%s/api/user/login"
	UrlSesTokInfo   = "http://%s/api/webserver/SesTokInfo"
	UrlDeviceStatus = "http://%s/api/monitoring/status"
	UrlSMSList      = "http://%s/api/sms/sms-list"
	UrlSendSMS      = "http://%s/api/sms/send-sms"
)

// ErrorResponse represents a generic error response from the API.
type ErrorResponse struct {
	XMLName   xml.Name `xml:"error"`   // XMLName is the XML element name for the error.
	ErrorCode string   `xml:"code"`    // ErrorCode is the code returned by the API, indicating the type of error.
	Message   string   `xml:"message"` // Message is the error message returned by the API.
}

// Device represents the device information and authentication details.
type Device struct {
	l         *zap.SugaredLogger // Logger instance for logging.
	client    *http.Client       // HTTP client for making requests.
	sessionID string             // Session ID for the current session.
	token     string             // Token for authentication.
	deviceIP  string             // IP address of the device.
	user      string             // Username for authentication.
	password  string             // Password for authentication.
}

func NewDevice(l *zap.SugaredLogger, deviceIP, user, password string) (*Device, error) {
	d := Device{
		l:        l,
		deviceIP: deviceIP,
		user:     user,
	}

	d.password = d.hashAndEncodePassword(password)

	client := &http.Client{
		Jar: nil, // cookie jar to store and manage cookies
	}
	client.Jar, _ = cookiejar.New(nil)

	d.client = client

	return &d, nil
}

package huaweimodem

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// SesTokInfo represents the session and token information required for authentication.
type SesTokInfo struct {
	XMLName xml.Name `xml:"response"` // XMLName is the XML element name for the response.
	SesInfo string   `xml:"SesInfo"`  // SesInfo is the session information.
	TokInfo string   `xml:"TokInfo"`  // TokInfo is the token information.
}

// Login authenticates with the device by obtaining session and token information,
// hashing the combined token, and sending a login request.
func (d *Device) Login() (err error) {

	// Get session and token information
	err = d.getSesTokInfo()
	if err != nil {
		return fmt.Errorf("failed to get SesTokInfo: %w", err)
	}

	// Combine user, password, and token, and hash the result
	combinedToken := fmt.Sprintf("%s%s%s", d.user, d.password, d.token)
	hashedCombinedToken := d.hashAndEncodePassword(combinedToken)

	// Create login payload
	loginPayload := fmt.Sprintf(`<request><Username>%s</Username><Password>%s</Password><password_type>4</password_type></request>`, d.user, hashedCombinedToken)
	req, err := http.NewRequest("POST", fmt.Sprintf(UrlLogin, d.deviceIP), bytes.NewBuffer([]byte(loginPayload)))
	if err != nil {
		return fmt.Errorf("failed to create login request: %w", err)
	}

	// Set headers for the request
	req.Header.Set("Content-Type", httpContentType)
	req.Header.Set("__RequestVerificationToken", d.token)
	req.Header.Set("Cookie", d.sessionID)

	// Send the request
	resp, err := d.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send login request: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed with status code %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for error response
	var errorResponse ErrorResponse
	if err := xml.Unmarshal(body, &errorResponse); err == nil {
		return fmt.Errorf("error code %s", errorResponse.ErrorCode)
	}

	d.l.Debug("login successfully")
	d.deviceStatus, err = d.DeviceStatus()
	if err != nil {
		return fmt.Errorf("failed to get device status: %w", err)
	}

	return nil
}

// getSesTokInfo fetches the session and token information required for authentication.
func (d *Device) getSesTokInfo() error {
	client := d.client
	deviceIP := d.deviceIP

	// Create the request URL
	requestUrl := fmt.Sprintf(UrlSesTokInfo, deviceIP)
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create SesTokInfo request: %w", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get SesTokInfo: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read SesTokInfo response: %w", err)
	}

	// Unmarshal the response into SesTokInfo struct
	var sesTokInfo SesTokInfo
	if err := xml.Unmarshal(body, &sesTokInfo); err != nil {
		return fmt.Errorf("failed to unmarshal SesTokInfo: %w", err)
	}

	// Set the session ID and token
	d.sessionID = sesTokInfo.SesInfo
	d.token = sesTokInfo.TokInfo

	d.l.Debug("sessionID: ", d.sessionID)
	d.l.Debug("token: ", d.token)

	return nil
}

// hashAndEncodePassword hashes the provided password using SHA-256 and then encodes the hash in base64 format.
func (d *Device) hashAndEncodePassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hasher.Sum(nil)
	hashedPasswordAsString := hex.EncodeToString(hashedPassword)
	encodedPassword := base64.URLEncoding.EncodeToString([]byte(hashedPasswordAsString))
	return encodedPassword
}

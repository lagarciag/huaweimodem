package huaweimodem

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// SMS represents the structure of an SMS request to be sent.
type SMS struct {
	XMLName  xml.Name `xml:"request"`  // XMLName is the XML element name for the request.
	Index    int      `xml:"Index"`    // Index is the message index, typically set to -1 for new messages.
	Phones   Phones   `xml:"Phones"`   // Phones contains a list of phone numbers to send the SMS to.
	Content  string   `xml:"Content"`  // Content is the text content of the SMS.
	Length   int      `xml:"Length"`   // Length is the length of the SMS content.
	Reserved int      `xml:"Reserved"` // Reserved is a reserved field, often set to 1.
	Date     string   `xml:"Date"`     // Date is the date the SMS is sent.
}

// Phones represents a list of phone numbers for the SMS.
type Phones struct {
	Phone []string `xml:"Phone"` // Phone is a list of phone numbers.
}

// SMSResponse represents the response received after sending an SMS.
type SMSResponse struct {
	XMLName   xml.Name `xml:"response"` // XMLName is the XML element name for the response.
	ErrorCode string   `xml:"code"`     // ErrorCode is the code returned by the API, indicating success or error.
	Message   string   `xml:"message"`  // Message is the message returned by the API, typically empty on success.
}

// SMSList represents the list of SMS messages retrieved from the device.
type SMSList struct {
	XMLName  xml.Name     `xml:"response"`         // XMLName is the XML element name for the response.
	Messages []SMSMessage `xml:"Messages>Message"` // Messages is a list of SMS messages.
}

// SMSMessage represents a single SMS message.
type SMSMessage struct {
	XMLName xml.Name `xml:"Message"` // XMLName is the XML element name for the message.
	Index   int      `xml:"Index"`   // Index is the index of the message.
	Phone   string   `xml:"Phone"`   // Phone is the phone number the message was sent from or to.
	Content string   `xml:"Content"` // Content is the content of the message.
	Date    string   `xml:"Date"`    // Date is the date the message was sent or received.
}

// DeleteSMSRequest represents the XML request to delete an SMS message.
type DeleteSMSRequest struct {
	XMLName xml.Name `xml:"request"`
	Index   int      `xml:"Index"`
}

// DeleteSMSResponse represents the XML response after deleting an SMS message.
type DeleteSMSResponse struct {
	XMLName xml.Name `xml:"response"`
	Result  string   `xml:"result"`
}

func (d *Device) ReadSMSInbox() (*SMSList, error) {

	if d.sessionID == "" {
		return nil, fmt.Errorf("you must login first")
	}

	err := d.getSesTokInfo()
	if err != nil {
		return nil, fmt.Errorf("failed to get SesTokInfo: %w", err)
	}

	reqBody := `<?xml version="1.0" encoding="UTF-8"?><request><PageIndex>1</PageIndex><ReadCount>20</ReadCount><BoxType>1</BoxType><SortType>0</SortType><Ascending>0</Ascending><UnreadPreferred>0</UnreadPreferred></request>`

	client := d.client
	req, err := http.NewRequest("POST", fmt.Sprintf(UrlSMSList, d.deviceIP), bytes.NewBufferString(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create SMS list request: %w", err)
	}

	req.Header.Set("Content-Type", httpContentType)
	req.Header.Set("Cookie", d.sessionID)
	req.Header.Set("__RequestVerificationToken", d.token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send SMS list request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read SMS list response: %w", err)
	}

	var smsList SMSList
	var errorResponse ErrorResponse
	if err := xml.Unmarshal(body, &smsList); err != nil {
		if err := xml.Unmarshal(body, &errorResponse); err == nil {
			return nil, fmt.Errorf("error code %s", errorResponse.ErrorCode)
		}
		return nil, fmt.Errorf("failed to unmarshal SMS list: %w", err)
	}

	return &smsList, nil
}

func (d *Device) SendSMS(phoneNumber, message string) error {
	if d.sessionID == "" {
		return fmt.Errorf("you must login first")
	}

	err := d.getSesTokInfo()
	if err != nil {
		return fmt.Errorf("failed to get SesTokInfo: %w", err)
	}

	sms := SMS{
		Index:    -1,
		Phones:   Phones{Phone: []string{phoneNumber}},
		Content:  message,
		Length:   len(message),
		Reserved: 1,
		Date:     time.Now().String(),
	}

	xmlData, err := xml.Marshal(sms)
	if err != nil {
		return fmt.Errorf("failed to marshal SMS request: %w", err)
	}

	client := d.client
	req, err := http.NewRequest("POST", fmt.Sprintf(UrlSendSMS, d.deviceIP), bytes.NewBuffer(xmlData))
	if err != nil {
		return fmt.Errorf("failed to create SMS request: %w", err)
	}
	req.Header.Set("Content-Type", httpContentType)
	req.Header.Set("Cookie", d.sessionID)
	req.Header.Set("__RequestVerificationToken", d.token)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send SMS request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read SMS response: %w", err)
	}

	var smsResponse SMSResponse
	var errorResponse ErrorResponse

	if err := xml.Unmarshal(body, &smsResponse); err == nil {
		if smsResponse.ErrorCode != "" {
			return fmt.Errorf("error code %s", smsResponse.ErrorCode)
		}
		d.l.Debug("SMS sent successfully")
		return nil
	} else if err := xml.Unmarshal(body, &errorResponse); err == nil {
		return fmt.Errorf("error code %s", errorResponse.ErrorCode)
	} else {
		return fmt.Errorf("unexpected response format")
	}
}

// DeleteSMSWithIndex deletes the first SMS message in the inbox.
func (d *Device) DeleteSMSWithIndex(index int) error {
	if d.sessionID == "" {
		return fmt.Errorf("you must login first")
	}

	if messages, err := d.ReadSMSInbox(); err == nil {
		if len(messages.Messages) == 0 {
			return fmt.Errorf("no messages to delete")
		}

		foundIndex := false
		for _, message := range messages.Messages {
			if message.Index == index {
				foundIndex = true
				break
			}
		}
		if !foundIndex {
			return fmt.Errorf("message with index %d not found", index)
		}

	} else {
		return fmt.Errorf("failed to read SMS inbox: %w", err)
	}

	err := d.getSesTokInfo()
	if err != nil {
		return fmt.Errorf("failed to get SesTokInfo: %w", err)
	}

	deleteRequest := DeleteSMSRequest{Index: index}
	xmlData, err := xml.Marshal(deleteRequest)
	if err != nil {
		return fmt.Errorf("failed to marshal delete request: %w", err)
	}

	d.l.Debug("xmlData: ", string(xmlData))

	url := fmt.Sprintf(UrlDeleteSMS, d.deviceIP)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(xmlData))
	if err != nil {
		return fmt.Errorf("failed to create delete SMS request: %w", err)
	}

	req.Header.Set("Content-Type", httpContentType)
	req.Header.Set("Cookie", d.sessionID)
	req.Header.Set("__RequestVerificationToken", d.token)

	resp, err := d.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send delete SMS request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("delete SMS request failed with status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read delete SMS response: %w", err)
	}

	var deleteResponse DeleteSMSResponse
	if err := xml.Unmarshal(body, &deleteResponse); err != nil {
		var errorResponse ErrorResponse
		if err = xml.Unmarshal(body, &errorResponse); err == nil {
			return fmt.Errorf("error code %s, message: %s", errorResponse.ErrorCode, errorResponse.Message)
		}

		return fmt.Errorf("failed to unmarshal delete SMS response: %w", err)
	}

	if deleteResponse.Result != "" {
		return fmt.Errorf("failed to delete SMS, result: %s", deleteResponse.Result)
	}

	d.l.Debug("SMS deleted successfully")

	return nil
}

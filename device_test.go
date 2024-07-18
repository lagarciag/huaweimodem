package huaweimodem

import (
	"github.com/k0kubun/pp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestSendSMS(t *testing.T) {
	logger := NewLogger()
	l := logger.Sugar()
	defer logger.Sync()

	t.Log("Starting test...")

	deviceIP := "192.168.8.1" // Replace with your device's IP address
	username := "admin"       // Replace with your username
	password := "250200.Ab"   // Replace with your password
	//phoneNumber := "+50688928380" // Ensure the phone number is in international format
	//message := "Hello, this is a test message."

	device, err := NewDevice(l, deviceIP, username, password)
	if err != nil {
		t.Fatalf("Error creating device: %v", err)
	}
	err = device.Login()
	if err != nil {
		t.Fatalf("Error getting session ID: %v", err)
	}

	t.Log("loggin successfull")

	smss, err := device.ReadSMSInbox()

	if err != nil {
		t.Fatalf("Error reading SMS inbox: %v", err)
	}

	pp.Println(smss)

	var status *DeviceStatus
	if status, err = device.GetDeviceStatus(); err != nil {
		t.Fatalf("Error getting device status: %v", err)
	}

	pp.Println(status)

	err = device.SendSMS("+50688928380", "hola mundo cruel 2")
	if err != nil {
		t.Fatalf("Error sending SMS: %v", err)
	}

	t.Log("sent SMS successfully")

}

// NewLogger creates a new zap.Logger with the provided options, colorized console output, and caller info
func NewLogger(options ...zap.Option) *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder, // Enable colors for log levels
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // Encode caller information
	}
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.DebugLevel)
	return zap.New(core, zap.AddCaller()).WithOptions(options...) // AddCaller to include caller information
}

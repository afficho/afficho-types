package types

// ScreenshotResponse is sent from device to cloud with a captured screenshot.
type ScreenshotResponse struct {
	// DeviceID identifies the device that captured the screenshot.
	DeviceID string `json:"device_id"`
	// ImageBase64 is the base64-encoded PNG image data.
	ImageBase64 string `json:"image_base64"`
	// Timestamp is the capture time in RFC 3339 format.
	Timestamp string `json:"timestamp"`
}

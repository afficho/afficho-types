package types

// ProofOfPlayReport is a batch of playback records sent from device to cloud.
type ProofOfPlayReport struct {
	// DeviceID identifies the reporting device.
	DeviceID string `json:"device_id"`
	// Records contains the individual playback events.
	Records []ProofOfPlayRecord `json:"records"`
}

// ProofOfPlayRecord is a single playback event proving that a content item
// was displayed for a given duration.
type ProofOfPlayRecord struct {
	// ContentID is the identifier of the content item that was played.
	ContentID string `json:"content_id"`
	// StartedAt is the UTC time the content began displaying (RFC 3339).
	StartedAt string `json:"started_at"`
	// DurationS is how long the content was displayed, in seconds.
	DurationS int `json:"duration_s"`
}

package types

// AlertMessage is sent from cloud to device to display an alert banner
// overlaying the current content.
type AlertMessage struct {
	// Text is the alert message shown on screen.
	Text string `json:"text"`
	// TTLS is the time-to-live in seconds. Zero means the alert stays
	// until manually cleared via a TypeClearAlert message.
	TTLS int `json:"ttl_s,omitempty"`
	// Urgency controls the visual style: "info", "warning", or "critical".
	Urgency string `json:"urgency"`
}

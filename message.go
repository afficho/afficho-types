// Package types defines the wire-format types shared between afficho-client
// and afficho-cloud. It contains no business logic and depends only on the
// standard library.
package types

import "encoding/json"

// WebSocket message type constants. These appear in the "type" field of every
// WSMessage exchanged between client and cloud.
const (
	// TypeCurrent tells the device which content item to display now.
	TypeCurrent = "current"
	// TypeReload instructs the device to reload its current view.
	TypeReload = "reload"
	// TypeAlert pushes an alert banner to the device.
	TypeAlert = "alert"
	// TypeClearAlert removes the current alert banner from the device.
	TypeClearAlert = "clear_alert"
	// TypeTicket is reserved for support-ticket messaging.
	TypeTicket = "ticket"
	// TypeSettings pushes updated device settings.
	TypeSettings = "settings"
	// TypeHeartbeatAck acknowledges a heartbeat from the device.
	TypeHeartbeatAck = "heartbeat_ack"
	// TypeSyncContent pushes content items for the device to download.
	TypeSyncContent = "sync_content"
	// TypeSyncPlaylist pushes a full playlist replacement to the device.
	TypeSyncPlaylist = "sync_playlist"
)

// WSMessage is the top-level WebSocket envelope. Every message exchanged
// between client and cloud is wrapped in this structure.
type WSMessage struct {
	// Type identifies the kind of message (see Type* constants).
	Type string `json:"type"`
	// Payload carries the type-specific data as raw JSON.
	Payload json.RawMessage `json:"payload"`
}

// Package types defines the wire-format types shared between afficho-client
// and afficho-cloud. It contains no business logic and depends only on the
// standard library.
package types

import "encoding/json"

// WebSocket message type constants. These appear in the "type" field of every
// WSMessage exchanged between client and cloud.
const (
	// Display control (local + cloud → device).

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

	// Cloud → device sync and commands.

	// TypeSyncContent pushes content items for the device to download.
	TypeSyncContent = "sync_content"
	// TypeSyncPlaylist pushes a full playlist replacement to the device.
	TypeSyncPlaylist = "sync_playlist"
	// TypeSyncSchedule pushes schedule definitions to the device.
	TypeSyncSchedule = "sync_schedule"
	// TypeCommand sends a command to the device (reload, reboot, update, screenshot).
	TypeCommand = "command"
	// TypeHeartbeatAck acknowledges a heartbeat from the device.
	TypeHeartbeatAck = "heartbeat_ack"

	// Device → cloud.

	// TypeRegister is sent by the device on first connect with DeviceRegistration payload.
	TypeRegister = "register"
	// TypeHeartbeat is sent periodically by the device with Heartbeat payload.
	TypeHeartbeat = "heartbeat"
	// TypeSyncAck acknowledges a completed content or playlist sync.
	TypeSyncAck = "sync_ack"
	// TypeProofOfPlay sends batched playback records to the cloud.
	TypeProofOfPlay = "proof_of_play"
	// TypeScreenshot sends a screenshot in response to a command.
	TypeScreenshot = "screenshot"
)

// WSMessage is the top-level WebSocket envelope. Every message exchanged
// between client and cloud is wrapped in this structure.
type WSMessage struct {
	// Type identifies the kind of message (see Type* constants).
	Type string `json:"type"`
	// Payload carries the type-specific data as raw JSON.
	Payload json.RawMessage `json:"payload"`
}

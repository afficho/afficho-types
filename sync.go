package types

// SyncAck is sent from device to cloud to acknowledge a completed sync.
type SyncAck struct {
	// SyncType identifies what was synced: "content" or "playlist".
	SyncType string `json:"sync_type"`
}

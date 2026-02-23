package types

// SyncAck is sent from device to cloud to acknowledge a completed sync.
type SyncAck struct {
	// SyncType identifies what was synced: "content", "playlist", or "schedule".
	SyncType string `json:"sync_type"`
	// Success indicates whether the sync completed without error.
	Success bool `json:"success"`
	// Error describes what went wrong if Success is false.
	Error string `json:"error,omitempty"`
}

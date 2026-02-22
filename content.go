package types

// ContentSyncItem represents a single piece of content pushed from cloud to
// device during a content sync. The device downloads and caches the content
// locally.
type ContentSyncItem struct {
	// ID is the unique identifier for this content item.
	ID string `json:"id"`
	// Name is the human-readable name of the content.
	Name string `json:"name"`
	// Type is the content kind: image, video, url, or html.
	Type string `json:"type"`
	// Source is the download URL (signed URL for media, plain URL for web content).
	Source string `json:"source"`
	// DurationS is how long the item should be displayed, in seconds.
	DurationS int `json:"duration_s"`
	// Checksum is the SHA-256 hex digest of the media file.
	Checksum string `json:"checksum"`
	// SizeBytes is the file size for download progress tracking.
	SizeBytes int64 `json:"size_bytes"`
	// AllowPopups controls whether the embedded browser permits popups.
	AllowPopups bool `json:"allow_popups"`
}

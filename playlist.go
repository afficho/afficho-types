package types

// PlaylistSync is sent from cloud to device as a full playlist replacement.
// The device discards its current playlist and adopts this one.
type PlaylistSync struct {
	// PlaylistID is the unique identifier for the playlist.
	PlaylistID string `json:"playlist_id"`
	// Name is the human-readable playlist name.
	Name string `json:"name"`
	// Items is the ordered list of content references in this playlist.
	Items []PlaylistSyncItem `json:"items"`
}

// PlaylistSyncItem is a single entry in a PlaylistSync, referencing a content
// item by ID with display ordering and duration.
type PlaylistSyncItem struct {
	// ContentID references a ContentSyncItem.ID.
	ContentID string `json:"content_id"`
	// Position is the display order (0-indexed).
	Position int `json:"position"`
	// DurationS overrides the content item's default display duration.
	DurationS int `json:"duration_s"`
}

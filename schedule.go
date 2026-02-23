package types

// ScheduleSync is sent from cloud to device as part of a schedule manifest.
// Each entry maps a time window to a playlist.
type ScheduleSync struct {
	// ID is the unique identifier for the schedule.
	ID string `json:"id"`
	// PlaylistID references the playlist to activate during this window.
	PlaylistID string `json:"playlist_id"`
	// CronExpr is a time-window expression in "HH:MM-HH:MM days" format,
	// where days is one of: everyday, weekdays, weekends.
	CronExpr string `json:"cron_expr"`
	// Priority determines which schedule wins when multiple overlap.
	// Higher values take precedence.
	Priority int `json:"priority"`
}

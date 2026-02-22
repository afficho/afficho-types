package types

// DeviceRegistration is sent from device to cloud on first connect.
// It identifies the device and its capabilities.
type DeviceRegistration struct {
	// DeviceID is a unique, stable identifier for the device.
	DeviceID string `json:"device_id"`
	// Hostname is the device's network hostname.
	Hostname string `json:"hostname"`
	// Arch is the CPU architecture (amd64, arm64, armv7, armv6).
	Arch string `json:"arch"`
	// OSVersion describes the operating system running on the device.
	OSVersion string `json:"os_version"`
	// AppVersion is the semantic version of the afficho-client binary.
	AppVersion string `json:"app_version"`
	// LocalIP is the device's LAN IP address.
	LocalIP string `json:"local_ip"`
}

// Heartbeat is sent periodically from device to cloud to report status.
type Heartbeat struct {
	// DeviceID identifies which device is reporting.
	DeviceID string `json:"device_id"`
	// CurrentItemID is the content item currently being displayed.
	CurrentItemID string `json:"current_item_id,omitempty"`
	// PlaylistID is the active playlist, if any.
	PlaylistID string `json:"playlist_id,omitempty"`
	// UptimeS is the device uptime in seconds.
	UptimeS int64 `json:"uptime_s"`
	// CPUTempC is the CPU temperature in degrees Celsius.
	CPUTempC float64 `json:"cpu_temp_c,omitempty"`
	// MemUsedPct is the percentage of memory in use (0–100).
	MemUsedPct float64 `json:"mem_used_pct"`
	// DiskUsedPct is the percentage of disk in use (0–100).
	DiskUsedPct float64 `json:"disk_used_pct"`
	// StorageUsed is the number of bytes used by afficho content storage.
	StorageUsed int64 `json:"storage_used_bytes"`
	// ScreenOn indicates whether the display is currently powered on.
	ScreenOn bool `json:"screen_on"`
	// Timestamp is the time of the heartbeat in RFC 3339 format.
	Timestamp string `json:"timestamp"`
}

// DeviceCommand is sent from cloud to device to trigger an action.
type DeviceCommand struct {
	// Command is the action to perform (reload, reboot, update, screenshot).
	Command string `json:"command"`
	// Params carries optional command-specific parameters.
	Params map[string]any `json:"params,omitempty"`
}

# Afficho Types — Implementation TODOs

Shared Go types imported by both `afficho-client` and `afficho-cloud`.
Contains **only** wire-format types — no business logic, no dependencies beyond stdlib.

**Legend:** `[ ]` pending · `[~]` in progress · `[x]` done

---

## Phase 1 — Core Wire Types

- [x] `WSMessage` — WebSocket envelope: `{ "type": "...", "payload": {} }`
  (replaces `api.Message` currently in the client)
- [x] `WSMessage.Type` constants: `current`, `reload`, `alert`, `clear_alert`,
  `ticket`, `settings`, `heartbeat_ack`, `sync_content`, `sync_playlist`
- [x] `DeviceRegistration` — device → cloud on first connect
  ```go
  type DeviceRegistration struct {
      DeviceID    string `json:"device_id"`
      Hostname    string `json:"hostname"`
      Arch        string `json:"arch"`        // amd64, arm64, armv7, armv6
      OSVersion   string `json:"os_version"`
      AppVersion  string `json:"app_version"`
      LocalIP     string `json:"local_ip"`
  }
  ```
- [x] `Heartbeat` — device → cloud, periodic
  ```go
  type Heartbeat struct {
      DeviceID      string  `json:"device_id"`
      CurrentItemID string  `json:"current_item_id,omitempty"`
      PlaylistID    string  `json:"playlist_id,omitempty"`
      UptimeS       int64   `json:"uptime_s"`
      CPUTempC      float64 `json:"cpu_temp_c,omitempty"`
      MemUsedPct    float64 `json:"mem_used_pct"`
      DiskUsedPct   float64 `json:"disk_used_pct"`
      StorageUsed   int64   `json:"storage_used_bytes"`
      ScreenOn      bool    `json:"screen_on"`
      Timestamp     string  `json:"timestamp"` // RFC 3339
  }
  ```
- [x] `ContentSyncItem` — cloud → device, content to download
  ```go
  type ContentSyncItem struct {
      ID          string `json:"id"`
      Name        string `json:"name"`
      Type        string `json:"type"`         // image, video, url, html
      Source      string `json:"source"`        // signed URL (media) or plain URL
      DurationS   int    `json:"duration_s"`
      Checksum    string `json:"checksum"`      // SHA-256 of media file
      SizeBytes   int64  `json:"size_bytes"`
      AllowPopups bool   `json:"allow_popups"`
  }
  ```
- [x] `PlaylistSync` — cloud → device, full playlist replacement
  ```go
  type PlaylistSync struct {
      PlaylistID string             `json:"playlist_id"`
      Name       string             `json:"name"`
      Items      []PlaylistSyncItem `json:"items"`
  }
  type PlaylistSyncItem struct {
      ContentID string `json:"content_id"`
      Position  int    `json:"position"`
      DurationS int    `json:"duration_s"`
  }
  ```
- [x] `AlertMessage` — cloud → device
  ```go
  type AlertMessage struct {
      Text    string `json:"text"`
      TTLS    int    `json:"ttl_s,omitempty"` // 0 = until manually cleared
      Urgency string `json:"urgency"`         // info, warning, critical
  }
  ```
- [x] `DeviceCommand` — cloud → device (reload, reboot, update, screenshot)
  ```go
  type DeviceCommand struct {
      Command string         `json:"command"` // reload, reboot, update, screenshot
      Params  map[string]any `json:"params,omitempty"`
  }
  ```

---

## Phase 2 — Client Migration

After the types are published, refactor the client:

- [ ] Replace `api.Message` in `afficho-client` with `types.WSMessage`
- [ ] Import shared types for future cloud sync code
- [ ] Keep the module minimal — no business logic, no dependencies beyond stdlib

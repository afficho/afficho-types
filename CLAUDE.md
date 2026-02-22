# Afficho Types — Claude Context

## What this is

Tiny Go module (`github.com/afficho/afficho-types`) containing **only** the
wire-format types and constants shared between `afficho-client` (edge daemon)
and `afficho-cloud` (SaaS backend).

See `TODOS.md` for all planned work.

## Rules

- **No business logic.** No database types. No HTTP types.
- Only the shapes of data that cross the wire between client and cloud.
- **Dependencies:** Standard library only (no third-party deps).
- **Versioning:** Semantic versioning. Both repos pin a specific version.
  Breaking changes require a coordinated version bump.

## Project layout

```
message.go      WSMessage struct + type constants
device.go       DeviceRegistration, Heartbeat, DeviceCommand
content.go      ContentSyncItem
playlist.go     PlaylistSync, PlaylistSyncItem
alert.go        AlertMessage
```

## Conventions

- Keep every file focused on one domain concept.
- All JSON tags must use `snake_case`.
- Document every exported type and field.
- Run `go vet ./...` before committing.

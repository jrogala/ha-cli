# ha-cli

CLI for Home Assistant smart home control.

## Install

Download a binary from the [latest release](https://github.com/jrogala/ha-cli/releases/latest), or install with Go:

```bash
go install github.com/jrogala/ha-cli@latest
```

## Setup

Set env vars or run setup:

```bash
export HA_URL=http://homeassistant.local:8123
export HA_TOKEN=your-long-lived-access-token
```

Or use the interactive setup:

```bash
ha setup --token YOUR_TOKEN
```

## Commands

| Command | Description |
|---|---|
| `setup` | Configure URL and token interactively |
| `config` | Show current Home Assistant configuration |
| `entities` | List entities, optionally filtered by domain |
| `state` | Get current state of an entity |
| `on` | Turn on an entity |
| `off` | Turn off an entity |
| `toggle` | Toggle an entity |
| `call` | Call any service with arbitrary data |
| `services` | List available services |

## Examples

```bash
$ ha entities --domain light
ENTITY_ID          STATE  NAME
light.kitchen      on     Kitchen Light
light.bedroom      off    Bedroom Light
light.living_room  on     Living Room

$ ha call light.turn_on --entity light.living_room --data '{"brightness_pct": 50}'
Called light.turn_on on light.living_room

$ ha toggle switch.garage_door
Toggled switch.garage_door

$ ha state sensor.outdoor_temperature
Entity:  sensor.outdoor_temperature
Name:    Outdoor Temperature
State:   21.5
Updated: 2026-03-24T10:00:00Z
```

## JSON Output

All commands support `--json` for machine-readable output.

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
ENTITY_ID                                                   STATE        NAME
light.ikea_of_sweden_tradfri_bulb_e27_cws_806lm_light_2    unavailable  Lumiere bureau
light.ikea_of_sweden_tradfri_bulb_e27_ww_806lm_light       unavailable  Salle de bain
light.ikea_of_sweden_tradfri_bulb_e27_ww_globe_806lm_light unavailable  Chambre

$ ha state sensor.flip_5_battery_level
Entity:  sensor.flip_5_battery_level
Name:    flip 5 Battery level
State:   72
Updated: 2026-03-18T18:06:00.382339+00:00
  unit_of_measurement: %
  device_class: battery

$ ha config
Name:     Home
Version:  2024.11.2
Timezone: Europe/Paris
```

## JSON Output

All commands support `--json` for machine-readable output.

# ha-cli

CLI for Home Assistant smart home control.

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
# List all light entities
ha entities --domain light

# Turn on a light at 50% brightness
ha call light.turn_on --entity light.living_room --data '{"brightness_pct": 50}'

# Toggle a switch
ha toggle switch.garage_door

# Check state of a sensor
ha state sensor.outdoor_temperature
```

## JSON Output

All commands support `--json` for machine-readable output.

# Jetson Fan Controller

**English** | [한국어](README-Korean.md)

> Lightweight, Easy to use, Advanced

Automatically control the fan speed of Jetson (Jetson Nano) by monitoring the temperature of module.

It uses a parabolic curve to control the speed.

## Installation

It doesn't require any languages (like python, go etc...), so you don't need to install other dependency. Just clone this repository and install it.

It requires `root` permission to installation.

```bash
git clone https://github.com/asheswook/jetson-fan-ctl
cd jetson-fan-ctl
./install.sh
```

## Settings

If you want to edit configuration, edit `/etc/jetson-fan-ctl.conf` file.

Its default value is automatically set as below, but you can change it if you want. (Requires restart service)

```json
{
  "FAN_OFF_TEMP": 30,
  "FAN_MAX_TEMP": 70,
  "INTERVAL": 2,
  "MAX_CLOCK_SPEED": true
}
```

```bash
sudo systemctl restart jetson-fan-ctl
```

- `FAN_OFF_TEMP`: Fan will be turned off if the temperature is lower than this value.
- `FAN_MAX_TEMP`: Fan will be turned on with maximum speed if the temperature is higher than this value.
- `INTERVAL`: Interval of checking temperature (in seconds).
- `MAX_CLOCK_SPEED`: If this value is `true`, the clock speed and power will be set as maximum (execute jetson_clocks)

## Uninstallation

There is an uninstall script in the build folder of this repo.

```bash
./uninstall.sh
```

# Jetson Fan Controller

[English](README.md) | **한국어**

> 가볍고, 사용하기 쉽고, 고도화된 젯슨 팬 제어기

Jetson (Jetson Nano)의 팬 속도를 모듈의 온도를 모니터링하여 자동으로 제어합니다.

팬 속도 제어에는 포물선 함수 (2차함수)를 사용합니다.

## Installation

Python, Go 등의 언어 설치가 필요하지 않으므로 다른 종속성을 설치할 필요가 없습니다. 이 레포를 클론하고 설치하면 됩니다.

설치에는 `root` 권한이 필요합니다.

```bash
git clone https://github.com/asheswook/jetson-fan-ctl
cd jetson-fan-ctl
./install.sh
```

## Settings

구성을 편집하려면 /etc/jetson-fan-ctl.conf 파일을 편집하면 됩니다.

기본 값은 다음과 같이 자동으로 설정되지만 원하는 경우 변경할 수 있습니다. (서비스 재시작 필요)

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

- `FAN_OFF_TEMP`: 온도가 이 값보다 낮으면 팬이 꺼집니다.
- `FAN_MAX_TEMP`: 온도가 이 값보다 높으면 팬이 최대 속도로 켜집니다.
- `INTERVAL`: 온도를 확인하고 팬 속도를 갱신하는 간격 (초 단위).
- `MAX_CLOCK_SPEED`: 이 값이 true인 경우 CPU 클럭 속도와 전력 공급이 최대로 설정됩니다 (jetson_clocks 실행)

## Uninstallation

이 레포의 build 폴더에 제거 스크립트가 있습니다. 실행하면 제거됩니다.

```bash
./uninstall.sh
```

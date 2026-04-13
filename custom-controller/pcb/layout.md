# Layout / Components

## BOM

| Name | Quantity | Link |
| ---- | -------- | ---- |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |
|      |          |      |

## Outline

USB-PD Controller (AP33772S)
- Request 28V via USB PD 3.1 EPR
- Connected to the Pico via I2C
- VOUT connected to a hardware on/off switch
- VOUT splits into two rails
  - One for the flip-power rail (main-VCC)
  - One for the pico and ICs on the display board (sub-VCC)

main-VCC
- N-channel MOSFET, enables power rail, via Pico
- Maybe an electronic fuse?
- Configurable step-down converter (13-24V), via I2C
  - Maybe with power / voltage sensing

sub-VCC (TPS62933?)
- Fixed step-down converter to 5V
- Must be capable to power Pico and ICs on the display!

Specials
- USB test points from the Pico need to be connected to the USB-C port
  - From the bottom of the PCB (requires holes) 
- Hardware reset button for the Pico

## USB-C Controller

- **AP33772S**
- USB-C Port connects to
  - CC1, CC2
- Ref implementation
  - https://github.com/CentyLab/AP33772S-CentyLab
  - https://github.com/CentyLab/AP33772S-Cpp
  - (https://github.com/CentyLab/PicoPD)
  - https://github.com/CentyLab/RotoPD
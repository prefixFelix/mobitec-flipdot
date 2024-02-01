<h1 align="center">Mobitec Flipdot Display</h1>

## Features

## Overview

## MobiPy

### Setup

## MobiController



## Controller

![controller-front](img/controller-front.jpg)

![controller-back](img/controller-back.jpg)

### General



### Ports

The board has a 10 and a 26 pin port.  
The 10-pin port is used to connect the HMI to the controller board. The HMI is operated by the bus driver, who controls what the flipdot displays show. Data between the HMI and the controller board is transmitted via [RS-485](https://en.wikipedia.org/wiki/RS-485). Power is as well provide via the HMI. The original cable *KABELUTSTICK* with article number *CE311007.1* is used for the connection. The female 10 pin plug of the cable should be of type XY and can be found here. Please note that although the original cable has two different wires soldered to the two GND pins of the plug, only pin 3 is connected to GND on the other end of the plug.  Pin number 5 therefore has no useful functionality. The male connector at the other end of the cable is a proprietary one with unusual notches.

![10pin](img/10pin.png)

| Pin  | Function       |
| ---- | -------------- |
| 1    | Data -         |
| 2    | Data +         |
| 3    | GND            |
| 4    | VCC 12-24V     |
| 5    | GND (Not used) |
| 6-10 | Not used       |

![plug](img/cable-plug.png) 



```
Output connector:
	- 26 pin (13 pin x2)
	Pinout (left to right, notch top)
		Top row
			- 11 VCC (24)
			- 12 GND (25)
		Bottom row
			- 24 VCC
			- 25 GND
```



### Interfaces

A button, a rotary selector and a status LED are provided on the board. If the button is pressed once, a test program starts, which outputs various information on the flipdot display. A button, a rotary switch and a status LED are provided on the board. If the button is pressed once, a test program starts, which outputs various information on the flipdot display. The rotary selector can be used to set the address of the board. It has a hexadecimal scale from 0x0 to 0xE. The blink codes of the status LED can be obtained from the following table:

| Sequence                | Function                    |
| ----------------------- | --------------------------- |
| Fast blinking           | idle, waiting for a command |
| Powered on, but LED off | Executing a command         |
| ...                     | ...                         |



## Related projects

- **[flipdot-mobitec (Anton Christensen)](https://github.com/anton-christensen/flipdot-mobitec)**
  - Custom controller based on an ESP32
- **[flipdot](https://github.com/openspaceaarhus/flipdot)**
  - Custom controller based on an ATmega
  - [Protocol description?](https://groups.google.com/g/openspaceaarhus/c/YMDPcS3pnHA) 
  - [Photos / videos](https://www.vagrearg.org/content/dotflipctl)
- **[mobitec-rs485 (duffrohde)](https://github.com/duffrohde/mobitec-rs485)**
  - Basic C RS-485 API
- **[pymobitec-flipdot (Bjarne)](https://github.com/bjarnekvae/pymobitec-flipdot)**
  - Simple Python RS-485 API
- **[mqtt-flipdot-driver (Chalmers Robotics)](https://github.com/ChalmersRobotics/mqtt-flipdot-driver)**
  - Simple Python RS-485 API
- **[maskin-flipdot (Nosen92)](https://github.com/Nosen92/maskin-flipdot)**
  - More advanced Python RS-485 API
- [**elektronikforumet**](https://elektronikforumet.com/forum/viewtopic.php?t=65264)
  - Protocol information
- [**buselektro**](https://www.busselektro.no/tips-og-funksjonsbeskrivelser/mobitec-rs485/)
  - Protocol information
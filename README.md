
<div align="center">
<img alt="Mobitec logo" src="docs/img/logo.png" heigth=200 width=300/>
<h1>Mobitec Flipdot Display</h1>
</div>

> [!WARNING]
>
> This repository is still **work in progress**!

## Features

- Golang / Python library to control the mobitec flipdot displays via serial with the original controller (*completed*).
- Replacement of the control board with an Raspberry Pi Pico / ESP32 to control the display directly (*under development*).
- Documentation of the original [control board](docs/controller.md) (*completed*).

## Serial libraries (original controller)

Controlling the flipdot display via the original controller is simple, effective and accessible, because very little additional hardware is required (namely just a RS-485 adapter, like [this one](https://www.berrybase.de/en/usb-rs485-konverter)).  You can find further instructions on how to connect everything [here](docs/controller.md#Connecting-to-the-Board). The Mobitec proprietary protocol is used for communication via the serial interface. Mobitec's own protocol is now used via the serial interface (very detailed [protocol description](https://github.com/Nosen92/maskin-flipdot/blob/main/protocol.md)). Basically, the controller can accept and display *text* in different fonts/sizes, predefined *symbols* and freely designable *bitmaps*. A list of fonts / symbols for our variations of boards can be found in [fonts.md](docs/fonts.md).  
**Capabilities / Limitations:**  
The controller is slow. Processing a command and updating the display takes several seconds. Animations can therefore not be displayed (< 1 FPS). However, this limitation has nothing to do with the flipdot display but with the controller.

### Golang

### Python

> [!TIP]
> A complete list of all functions and their example use can be found in [example.py](serial-lib/py/MobiPy/example.py).

The python libary offers the following featrues:

- Text with all available fonts / sizes.
- Symbols
- Bitmap: basic operations (fill, invert shift); basic geometric shapes (dot, line, rectangle)

## Replacement controller board

### Parts / Order

### Flash

### Usage

## Related projects

- **[flipdot-mobitec (Anton Christensen)](https://github.com/anton-christensen/flipdot-mobitec)**
    - Custom controller based on an ESP32
- **[flipdot (Open Space Aarhus)](https://github.com/openspaceaarhus/flipdot)**
    - Custom controller based on an ATmega
    - [Protocol description?](https://groups.google.com/g/openspaceaarhus/c/YMDPcS3pnHA)
    - [Photos / videos](https://www.vagrearg.org/content/dotflipctl)
- **[flipdot-games (Anton Berneving)](https://github.com/antbern/flipdot-games)**
    - Custom controller board based on an Raspberry Pi Pico
    - Firmware written in Rust
- **[mobitec-rs485 (duffrohde)](https://github.com/duffrohde/mobitec-rs485)**
    - Basic C RS-485 API
- **[pymobitec-flipdot (Bjarne)](https://github.com/bjarnekvae/pymobitec-flipdot)**
    - Simple Python RS-485 API
- **[mqtt-flipdot-driver (Chalmers Robotics)](https://github.com/ChalmersRobotics/mqtt-flipdot-driver)**
    - Simple Python RS-485 API
- **[maskin-flipdot (Nosen92)](https://github.com/Nosen92/maskin-flipdot)**
    - More advanced Python RS-485 API
    - Detailed serial protocol [documentation](https://github.com/Nosen92/maskin-flipdot/blob/main/protocol.md)
- [**elektronikforumet**](https://elektronikforumet.com/forum/viewtopic.php?t=65264)
    - Protocol information
- [**buselektro**](https://www.busselektro.no/tips-og-funksjonsbeskrivelser/mobitec-rs485/)
    - Protocol information
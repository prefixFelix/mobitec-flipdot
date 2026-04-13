<div align="center">
<img alt="Mobitec logo" src="docs/img/logo.png" heigth=200 width=300/>
<h1>Mobitec Flipdot Display</h1>
</div>


> [!WARNING]
>
> This repository is still **work in progress**!

## Features

- Golang / Python library to control the mobitec flipdot displays via serial with the original controller (*completed*).
- Custom PCB that is based on the Raspberry Pi Pico (W) to replace the original controller (*under development*).
- Documentation of the original [control board](docs/controller.md) (*completed*).

## Serial libraries (original controller)

Controlling the flipdot display via the original controller is simple, effective and accessible, because very little additional hardware is required (namely just a RS-485 adapter, like [this one](https://www.berrybase.de/en/usb-rs485-konverter)). You can find further instructions on how to connect everything [here](docs/controller.md#Connecting-to-the-Board). Mobitec's own protocol is used for communication via the serial interface. The exact structure and flow of the protocol is documented [here](https://github.com/Nosen92/maskin-flipdot/blob/main/protocol.md). Basically, the controller can accept and display *text* in different fonts/sizes, predefined *symbols* and freely designable *bitmaps*. An incomplete list of fonts / symbols for our variations of boards can be found in [fonts.md](docs/fonts.md).  

> [!IMPORTANT]
>
> The controller is slow. Processing a command and updating the display takes several seconds. Animations can therefore not be displayed (< 1 FPS). However, this limitation has nothing to do with the flipdot display but with the controller.

### Golang

The golang library offers the following features:

- Text with all available fonts / sizes.
- Custom fonts via BDF files.
- Matrix/bitmap: basic operations (fill, negate, shift); row/column manipulation; path drawing.

All available features and how to use them are outlined in the [main.go](serial-lib/go/mobitec/examples/matrix-manipulation/main.go) file. Make sure that Go 1.20+ is installed. Dependencies are fetched automatically via `go mod`.

### Python

The python library offers the following features:

- Text with all available fonts / sizes.
- Symbols
- Bitmap: basic operations (fill, invert shift); basic geometric shapes (dot, line, rectangle)

All available features and how to use them are outlined in the [example.py]() file. Make sure that `serial` and `numpy` are installed via pip!

## Replacement controller board

> [!WARNING]
>
> Under development! Updates in the `custom-controller` dir.

Our PCB replaces the original Mobitec controller and directly drives the register ICs on the flipdot display. This enables significantly faster flip times than with the original controller board. The display state data is transmitted as JSON to the Pico via WebSocket / REST API. As with the old controller, this can either be a bitmap representation of the display or just text. Latency-sensitive content such as animations or scrolling text is processed directly on the Pi. Other projects have already designed their own PCB or used a Pico as a controller (see table).

- Single USB-C connector for data and power
- Data connection via WiFi
- Variable voltage 12-24V
- Raspberry Pi Pico W
- TinyGo
- *Requirement: Charger that supports USB PD 3.1 ERP (28V)*

### Parts

Added later

### Flash

> [!NOTE]
>
> The current TinyGo implementation is functional but with limited features.

### Usage

Added later

## Related projects

| Project Name                                                 | Author            | Serial Procotol | Direct Display Driver | Notes                                                        |
| ------------------------------------------------------------ | ----------------- | --------------- | --------------------- | ------------------------------------------------------------ |
| [flipdot-mobitec](https://github.com/anton-christensen/flipdot-mobitec) | Anton Christensen |                 | X                     | Custom controller based on an ESP32                          |
| [flipdot](https://github.com/openspaceaarhus/flipdot)        | Open Space Aarhus |                 | X                     | Custom controller based on an ATmega + PCB ([Photos / videos](https://www.vagrearg.org/content/dotflipctl), [Additional infos](https://groups.google.com/g/openspaceaarhus/c/YMDPcS3pnHA)) |
| [flipdot-games](https://github.com/antbern/flipdot-games)    | Anton Berneving   |                 | X                     | Custom controller board based on an Raspberry Pi Pico (Rust) |
| [mobitec-rs485](https://github.com/duffrohde/mobitec-rs485)  | duffrohde         | X               |                       | Basic C RS-485 API                                           |
| [pymobitec-flipdot](https://github.com/bjarnekvae/pymobitec-flipdot) | Bjarne            | X               |                       | Simple Python RS-485 API                                     |
| [mqtt-flipdot-driver](https://github.com/ChalmersRobotics/mqtt-flipdot-driver) | Chalmers Robotics | X               |                       | Simple Python RS-485 API                                     |
| [maskin-flipdot](https://github.com/Nosen92/maskin-flipdot)  | Nosen92           | X               |                       | Advanced Python RS-485 API + detailed serial protocol [documentation](https://github.com/Nosen92/maskin-flipdot/blob/main/protocol.md) |
| [FlipDot_ESP_MQTT](https://github.com/PaulSalz/FlipDot_ESP_MQTT) | Paul Salz         | X               |                       | ESP32 MQTT API                                               |
| [elektronikforumet](https://elektronikforumet.com/forum/viewtopic.php?t=65264) | -                 |                 |                       | Protocol information                                         |
| [buselektro](https://www.busselektro.no/tips-og-funksjonsbeskrivelser/mobitec-rs485/) | -                 |                 |                       | Protocol information                                         |
| [Radow](https://radow.org/flip-dot-en.php)                   | Rainer Radow      |                 |                       | Useful knowledge about flipdots                              |
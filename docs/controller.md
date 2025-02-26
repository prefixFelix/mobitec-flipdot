# Mobitec Controller Board

We encountered multiple different controllers. They seem to be functionally the same, even though they are quite different in their appearances. The boards are equal in that they have a 26 pin port that is connected to the flipdot display and a 10 pin port, that would normally be connected to an HMI (Human Machine Interface), with which the Bus Driver would control what is shown on the display under normal operation. They also have a rotary switch, with which the boards address is selected. The boards were located on the left side of the flipdot-enclosure (where the cable come out of it), fixed under the flipdot display with some screws.

![controller-1](img/controller-1.png)
![controller-1-back](img/controller-1-back.png)
![controller-2](img/controller-2.png)
![controller-2-back](img/controller-2-back.png)

One of the boards (*normal board*) has a red button, with which a test-program can be started. It also seems to have a swappable EEPROM (?) and a single status LED. We think that this is a normal production board, which was delivered to the customers. The other board (*development board*) uses other ICs which look more modern. It also has no button and two status LEDs. Since there are additional PINs on this board, which look like debug PINs to us, we assume that this board was possibly used in the development. But it could also just be a newer revision. Both boards behave similarly and have the same execution / flip times.

## Status LEDs

**Normal Board (One LED)**

| Sequence      | Function                    |
|---------------|-----------------------------|
| Fast blinking | Idle, waiting for a command |
| LED off       | Executing a command         |

**Development-Board (Two LEDs)**

| Sequence (Two LEDs)                          | Function                    |
|----------------------------------------------|-----------------------------|
| LED1: on, LED2: blinking                     | Idle, waiting for a command |
| LED1: off, LED2: blinking                    | Error with last message     |
| LED1: on, LED2: not blinking (static on/off) | Executing a command         |


## Connecting to the Board

The board can be controlled via a serial interface ([RS-485](https://en.wikipedia.org/wiki/RS-485)). To control the board with your PC you can simply buy a cheap USB-to-RS-485 converter. Connect the two data terminals of the converter to the corresponding data pins on the controller board (see table / picture). In addition, connect a power supply with an output voltage of 12-14V to the voltage and ground pins of the controller. It is recommended to use a laboratory power supply for testing. The 10 pin plug on the PCB seems to be approximately of the type `JST-XH-9S1P`. Ordinary jumper cables also fit on the pins.

| Pin  | Function                           |
| ---- | ---------------------------------- |
| 1    | Data -                             |
| 2    | Data +                             |
| 3    | GND                                |
| 4    | VCC 12-24V                         |
| 5    | GND (Not used, shorted with pin 3) |
| 6-10 | Not used                           |

![10pin_dia](img/10pin_dia.png)

![10_pin](img/10pin.png)

Our flipdot display was sold to us with the original connection cables. These have the 10 pin connector at one end and a proprietary 4 pin connector at the other. The pinout is shown in the graphic below (please note the cut-outs for orientation).

![plug](img/cable-plug.png)


## Flipdot Display Port

There is a 26-pin connector on the controller board that connects the controller to the flipdot display. Through this connector, the ICs on the flipdot display are controlled, which apply the "flip voltage" with the correct polarity and for the correct row/column on the display. The ICs are controlled with a logic voltage of 3-5V. The display is also supplied with power through the connector (12-24V (dots) + 5V (ICs)). The functionality of flipdot displays is already described in these projects: 

- [Open Space Aarhus](https://github.com/openspaceaarhus/flipdot/blob/master/flipper/master_setup.pdf)
- [Anton Christensen](https://github.com/anton-christensen/flipdot-mobitec/blob/master/mobitecSign/output.pdf) / [Anton Christensen](https://github.com/anton-christensen/flipdot-mobitec/blob/master/controllerDesign/controller.pdf)
- [Anton Berneving](https://github.com/antbern/flipdot-games/blob/main/pico-hardware/schematic_rev1.pdf)

The pinout of the connector can be seen in the image below (please note the lateral notch for orientation).
![26pin](img/26pin.png)

## Capabilities and Limitations

One drawback of the controller is its speed: On power on, it can take 5-15 seconds for the board to be ready for commands. In addition, the processing is very slow aswell. No matter how much data is actually sent to the board, it takes 1-2 seconds for it to process. If a full bitmap matrix is sent, all dots are energized, regardless of whether they changed or not. This can be sped up, by managing the dots state in software and only sending small matrices to the board, but even when only one column (1x16) is set at a time, the board struggles to reach 1 fps, executing 2-3 packets in quick succession and then lagging for over a full second. This means that the controller is not suitable for animations of any kind. This [seems to be](https://www.youtube.com/watch?v=opCHlJ_8fGk) a controller problem, not a flipdot one.

## Notes

- The devboard repeats the command every minute, even tho no new command was send.
- There are also LEDs, which are placed on the bottom of the display enclosure. These illuminate the dots from below at an angle. These LEDs can only be switched on or off and are operated with 24V via a simple 2-pin connector.
- Since basically every project we found had a different list of fonts, we assume that the fonts may not be the same for every board (alternatively, people just made mistakes). The fonts and their implementation by the contoller is quite strange, with most of the fonts having some oddities and drawbacks. 
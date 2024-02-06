 # MobiController

## Pinout

```
Lanes (top to bittom):
26 Not connected to somthing, except the plug
23
24
21
22
19 Not connected to somthing, except the plug
5V
GND
24V
---------------------------------------
      Main plug

    |  A  | 
  3 | --- | VCC | 4
  1 | GND | GND | 2
        Bottom
----------------------------------------
TC74 HC02A?		-> Quad 2-Input NOR Gate
	-> 7402
OUT |     | VCC |
 IN |   9 |
 IN |  12 |
OUT |
 IN |  17 |  y? | OUT
 IN |  17 |  x? | IN
    | GND |  x? | IN
	
74V HC139		-> Dual 2-to-4 Decoder/Demultiplexer
	-> 74LS139

MM74HC4514WM	-> 4-to-16 line decoder !5V!
ULN2003A 		-> High-Current Darlington Transistor Arrays !5V->24V!
TD62783AF 		-> HIGH−VOLTAGE SOURCE DRIVER !5V->24V!
```



```
Example:
- Switching dot x=4, y=7 form black to white
	- Panel = 0
	- 

```



## Anton Christensen C-Code

```c
// Board used: esp32doit-devkit-v1

// Current display state; Colors: WHITE = 0, BLACK = 1, TOGGLE 2, DONOTHING 3
unsigned char screenState[PANEL_WIDTH][PANEL_HEIGHT] PROGMEM; // PROGMEM Daten im flash speichern

// Display data which will be displayed in the future -> overwrites screenState
unsigned char screenBuffer[PANEL_WIDTH][PANEL_HEIGHT] PROGMEM;

// Generate lookup table for decimal to binary mapping 
unsigned char index_to_bitpattern_map[28];
int FlipScreen::index_to_bitpattern(int index) {
	return 1+index + (index/7);
}
// fill fast index-map
for(int i = 0; i < 28; i++) {
	index_to_bitpattern_map[i] = index_to_bitpattern(i);
}
// --> [1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20, 21, 22, 23, 25, 26, 27, 28, 29, 30, 31]

const unsigned char panel_triggers[4] = {13,14,15,16};
const unsigned char row_addr_pins[5] =  {25,26,27,32,33};
const unsigned char col_addr_pins[5] =  {18,19,21,22,23};
const unsigned char color_pin = 17;
//const unsigned char backlight_pin = 24;


void FlipScreen::_setDot(unsigned int x, unsigned int y, unsigned char color, bool longForcing) {
	// Ignore the dot if it is already set to the right color
    if(x < 0 || y < 0 || x >= PANEL_WIDTH || y >= PANEL_HEIGHT || color > 1 || this->screenState[x][y] == color) return;
    this->screenState[x][y] = color;

    int panel = x/28; // Number of the flipdot-panel on which the dot is to be flipped   
    x %= 28; // x-coordinate relative to the current flipdot-panel
	
    // Convert the integer value of the coordinates into a binary value:
	// The binary value indicates the address lanes which must be high or low -> 5bit long
    for(int i = 0; i < 5; i++) {
    	_digitalWrite(this->row_addr_pins[i], index_to_bitpattern_map[y]&(1<<i));
    	_digitalWrite(this->col_addr_pins[i], index_to_bitpattern_map[x]&(1<<i));
    }

    _digitalWrite(this->color_pin, color); // High = BLACK, Low = WHITE

    delayMicroseconds(10);
    _digitalWrite(this->panel_triggers[panel], HIGH);
    // delayMicroseconds(color == BLACK ? 300 : 300); 
	
    // Different state changes have different on-off times
    delayMicroseconds(longForcing ? 1000 : color == BLACK ? this->WhiteToBlackMicroseconds : this->blackToWhiteMicroseconds); 
    _digitalWrite(this->panel_triggers[panel], LOW);
	delayMicroseconds(10);
}
```

## Open Space Aarhus C-Code

```C
/* Enable all outputs we need */
DDRD = 0xfe;
DDRB = 0x1f;
DDRC = 0x3f;

static const unsigned char x_index[] = {1, 4, 5, 2, 3,
										6, 7, 9, 12, 13,
										10, 11, 14, 15, 17,
										20, 21, 18, 19, 22,
										0,0,0,0,0};

// Lookup table for column addresses?
static const unsigned char y_index[] = {
	0x10,0x08,0x18,0x04,0x14,0x0c,0x1c,0x12,0x0a,0x1a,0x06,0x16,0x0e,0x1e,0x11,0x09,
	0x19,0x05,0x15,0x0d,0x1d,0x13,0x0b,0x1b,0x07,0x17,0x0f,0x1f,0x10,0x08,0x18,0x04,
	0x14,0x0c,0x1c,0x12,0x0a,0x1a,0x06,0x16,0x0e,0x1e,0x11,0x09,0x19,0x05,0x15,0x0d,
	0x1d,0x13,0x0b,0x1b,0x07,0x17,0x0f,0x1f,0x10,0x08,0x18,0x04,0x14,0x0c,0x1c,0x12,
	0x0a,0x1a,0x06,0x16,0x0e,0x1e,0x11,0x09,0x19,0x05,0x15,0x0d,0x1d,0x13,0x0b,0x1b,
	0x07,0x17,0x0f,0x1f,0x10,0x08,0x18,0x04,0x14,0x0c,0x1c,0x12,0x0a,0x1a,0x06,0x16,
	0x0e,0x1e,0x11,0x09,0x19,0x05,0x15,0x0d,0x1d,0x13,0x0b,0x1b,0x07,0x17,0x0f,0x1f,
	0x10,0x08,0x18,0x04,0x14,0x0c,0x1c,0x12};

// Lookup table for panel addresses?
static const unsigned char y_board[] = {
	0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,
	0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x80,0x80,0x80,0x80,
	0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,
	0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x80,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,
	0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,
	0x20,0x20,0x20,0x20,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,
	0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,0x08,
	0x00,0x00,0x00,0x00,0x00,0x00,0x00,0x00};

pixel(uart_buf[0], uart_buf[1], uart_buf[0] & 0x80);
static void pixel (unsigned char x, unsigned char y, unsigned char set)
{
    // PORTB, PORTC, PORTD collection of GPIO pins?
    PORTB = x_index[x & 0x1f]; // 0x1f = 31
	
    // 
	if (set) {
		/* Select Y Board and B1.24 */
		PORTC = y_index[y];
		PORTD = y_board[y] | 0x10;
    //
	} else {
		/* Select Y Board, BX.23 and B2.24 */
		PORTC = y_index[y] | 0x20;
		PORTD = y_board[y] | 0x4;
	}

	_delay_us(320);

	PORTD = 0;
}
```


 

```

|  A  | 
| --- | VCC |
| GND | GND |

TC74 HC02A?		-> Quad 2-Input NOR Gate
	-> 7402
74V HC139		-> Dual 2-to-4 Decoder/Demultiplexer
	-> 74LS139

MM74HC4514WM	-> 4-to-16 line decoder !5V!
ULN2003A 		-> High-Current Darlington Transistor Arrays !5V->24V!
TD62783AF 		-> HIGH−VOLTAGE SOURCE DRIVER !5V->24V!
```

```c
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


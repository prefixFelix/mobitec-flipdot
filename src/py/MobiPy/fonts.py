class Font:
    """
    Basic font objects.
    Attributes:
        height (byte): Height of the font. Used for position glitch fix.
        code (byte): Font code used by the sign. Between 0x60 - 0x80.
    """

    def __init__(self, height, code):
        self.height = height
        self.code = code

# Font height != actual letter height!!! Don't ask me why
F20_NUMBERS_UPPER       = Font(19, 0x6B)    # Only uppercase and numbers
F20_16_NUMBERS_UPPER    = Font(19, 0x6A)    # Only uppercase and numbers / numbers 20px letters: 16px

F19                     = Font(23, 0x63)

F16                     = Font(15, 0x78)    # Upper/lowercase different height
F16_MEDIUM              = Font(19, 0x68)    # Gg not in line
F16_NUMBERS             = Font(15, 0x74)    # Only numbers
F16_NUMBERS2            = Font(15, 0x6E)    # Only numbers

F15                     = Font(15, 0x76)    # Normal
F15_MEDIUM              = Font(15, 0x71)    # Upper/lowercase letter not in line

F14_NUMBERS             = Font(13, 0x6F)    # Only numbers

F13                     = Font(15, 0x61)    # Normal width
F13_MEDIUM              = Font(15, 0x69)    # Medium width
F13_SMALL               = Font(15, 0x73)    # Smaller than medium width
F13_NUMBERS             = Font(15, 0x79)    # Only numbers and A

F12_NUMBERS             = Font(11, 0x6C)    # Only numbers

F9                      = Font(11, 0x62)    # Normal width
F9_MEDIUM_UPPER_NUMBERS = Font(8, 0x70)     # Mono
F9_SMALL_UPPER_NUMBERS  = Font(8, 0x75)     # Mono

F7                      = Font(8, 0x64)     # Normal width
F7_MEDIUM               = Font(8, 0x65)     # Medium width
F7_SMALL_UPPER          = Font(6, 0x6D)     # Small width

F6                      = Font(6, 0x66)
F5                      = Font(4, 0x72)


class Symbol:
    """
    Basic symbol objects.
    Attributes:
    """
    def __init__(self, value):
        self.value = value
        self.font = Font(0, 0x67)

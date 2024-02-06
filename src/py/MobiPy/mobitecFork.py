"""

--- Mobitec Fork ---
https://github.com/Nosen92/maskin-flipdot/blob/main/mobitec.py

"""

import serial
import font
import symbol
import bitmap

SPECIAL_CHARS = {
    "Å": 0x5d,
    "å": 0x7d,
    "Ä": 0x5b,
    "ä": 0x7b,
    "Ö": 0x5c,
    "ö": 0x7c
}

class MobitecDisplay:
    def __init__(self, port, address, width, height):
        """"Mobitect Flipdot-display object."""
        self.port = port
        self.address = address
        self.width = width
        self.height = height

        self.text_buffer = []
        self.bitmap_buffer = []

    def display(self):
        """Sends contents of the buffer to the display."""
        packet = self._create_packet()
        with serial.Serial(self.port, 4800, timeout=1) as ser:
            ser.write(packet)
        self.text_buffer = []
        self.bitmap_buffer = []

    def _create_packet(self):
        """Serializes all data and generates a complete Mobitec packet."""
        packet = bytearray()
        packet.append(0xFF) # Start byte
        packet.extend(self._get_packet_header())

        for text in self.text_buffer:
            packet.extend(self._serialize_text(text))
        for bm in self.bitmap_buffer:
            packet.extend(self._serialize_bitmap(bm))

        packet.extend(self._calculate_check_sum(packet))
        packet.append(0xFF) # Stop byte

        return packet

    def _get_packet_header(self):
        """Generates mobitec protocol packet header."""
        # Only the address and 0xA2 are required!
        return bytearray([self.address, 0xA2, 0xD0, self.width, 0xD1, self.height])

    def _calculate_check_sum(self, packet):
        """Algorithm is: add up all bytes (except start byte). The least significant byte is the check sum.
        Special cases for 0xfe and 0xff."""
        packet_sum = 0
        for byte in packet[1:]:
            packet_sum = packet_sum + byte
        packet_sum = packet_sum & 0xFF

        check_sum = bytearray()
        if packet_sum == 0xFE:
            check_sum.append(0xFE)
            check_sum.append(0x00)
        elif packet_sum == 0xFF:
            check_sum.append(0xFE)
            check_sum.append(0x01)
        else:
            check_sum.append(packet_sum)
        return check_sum

    def _serialize_bitmap(self, bitmap):
        """Serializes bitmap to mobitec protocol."""
        data = bytearray()
        mobitec_subcolumn_matrix = bitmap.convert_to_sub_column()
        for band in range(len(mobitec_subcolumn_matrix)):
            data_header = self._get_data_header(0x77, bitmap.x_pos, bitmap.y_pos + band * 5 + 4)
            data.extend(data_header)
            for subcolumn in range(bitmap.width):
                subcolumn_code = self.addBits(mobitec_subcolumn_matrix[band][subcolumn])
                data.append(subcolumn_code)
        return data

    def addBits(self, bits):
        ret = 32
        for i in range(len(bits)):
            ret += bits[i]*2**i
        return ret

    def _serialize_text(self, text):
        """Serializes text object to mobitec protocol.
        Accounts for deviations from ASCII codes."""
        horizontal_offset = text.pos_x
        vertical_offset = text.pos_y + text.font.height
        data = self._get_data_header(text.font.code, horizontal_offset, vertical_offset)

        for char in text.string:
            if char in SPECIAL_CHARS:
                char = SPECIAL_CHARS[char]
            else:
                char = ord(char)
            data.append(char)
        return data

    def _get_data_header(self, font, horizontal_offset, vertical_offset):
        """Generates mobitec protocol data section header."""
        return bytearray([0xD2, horizontal_offset, 0xD3, vertical_offset, 0xD4, font])

    def set_text(self, string, font, x_pos=0, y_pos=0):
        """Adds text to the text buffer."""
        text = _BasicText(string, font, x_pos, y_pos)
        self.text_buffer.append(text)

    def set_symbol(self, symbol, x_pos=0, y_pos=0):
        """Text Wrapper"""
        text = _BasicText(symbol.value, symbol.font, x_pos, y_pos)
        self.text_buffer.append(text)

    def set_bitmap(self, bm, x_pos=0, y_pos=0):
        """Adds bitmap to the bitmap buffer."""
        bm.x_pos = x_pos
        bm.y_pos = y_pos
        self.bitmap_buffer.append(bm)

class _BasicText:
    """
    Basic text objects. Gets queued in the buffer.
    Attributes:
        string (string): Text to be written.
        font (Font): Font to write the text width.
        pos_x (byte): Horizontal offset from left side.
        pos_y (byte): Vertical offset from upper side.
    """
    def __init__(self, string, font, pos_x, pos_y):
        self.string = string
        self.font = font
        self.pos_x = pos_x
        self.pos_y = pos_y


if __name__ == '__main__':
    """ DEV TEST """

    flipdot = MobitecDisplay('/dev/ttyUSB0', address=0x0b, width=28, height=16)

    circle = [
        [0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0],
        [0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0],
        [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
        [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
        [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1],
        [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
        [0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
        [0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0],
        [0, 0, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0]
    ]

    test = [[1, 1, 1, 1],
            [1, 0, 0, 1],
            [1, 0, 0, 1],
            [1, 1, 1, 1]]

    bm = bitmap.Bitmap(28, 16)

    bm.fill(1)
    bm.invert()
    bm.line_horizontal(1,1, 10, 2, 1)
    bm.line_vertical(1, 1, 10, 2, 1)
    bm.line_diagonal_right(2, 2, 10, 2, 1)
    bm.dot(14, 14, 1)
    bm.rectangle(18, 3, 4, 4, 1)
    bm.shift_left(11, keep=True)

    flipdot.set_bitmap(bm)
    flipdot.set_symbol(symbol.SOCCER1)
    flipdot.display()

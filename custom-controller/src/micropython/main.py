from machine import Pin
import time

class LEDMatrix:
    # Constants for the display
    ROW_DRIVER_ROWS = 16
    COL_DRIVER_COLUMNS = 28

    def __init__(self, rows=16, cols=28):
        self.ROWS = rows
        self.COLS = cols

        # Initialize the pins
        self.row_address = [
            Pin(2, Pin.OUT),
            Pin(3, Pin.OUT),
            Pin(4, Pin.OUT),
            Pin(5, Pin.OUT),
            Pin(6, Pin.OUT)
        ]
        self.row_high_en = Pin(7, Pin.OUT)
        self.row_low_en = Pin(8, Pin.OUT)
        self.col_address = [
            Pin(9, Pin.OUT),
            Pin(10, Pin.OUT),
            Pin(11, Pin.OUT),
            Pin(12, Pin.OUT),
            Pin(13, Pin.OUT)
        ]
        self.col_high_low = Pin(14, Pin.OUT)
        self.col_select = [Pin(16, Pin.OUT), Pin(17, Pin.OUT)]

        # Initialize display buffers
        self.buffer_active = [[False for _ in range(cols)] for _ in range(rows)]
        self.buffer_shadow = [[False for _ in range(cols)] for _ in range(rows)]

        # Row address lookup table
        self.row_lookup = [1, 4, 5, 2, 3, 6, 7, 9, 12, 13, 10, 11, 14, 15, 17, 20]

        # Column address lookup table
        self.col_lookup = [
            1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 18,
            19, 20, 21, 22, 23, 25, 26, 27, 28, 29, 30, 31
        ]

    def set_row_address(self, row):
        """Set the row address pins based on lookup table"""
        address = self.row_lookup[row]
        for i in range(5):
            self.row_address[i].value((address >> i) & 1)

    def set_row_direction(self, high_drive):
        """Set row direction (high or low drive)"""
        if high_drive:
            self.row_low_en.value(0)
            self.row_high_en.value(1)
        else:
            self.row_high_en.value(0)
            self.row_low_en.value(1)

    def set_column_address(self, column):
        """Set the column address pins based on lookup table"""
        address = self.col_lookup[column]
        for i in range(5):
            self.col_address[i].value((address >> i) & 1)

    def set_column_direction(self, high_drive):
        """Set column direction (high or low drive)"""
        self.col_high_low.value(not high_drive)

    def set_column_enabled(self, column=None):
        """Enable or disable column drivers"""
        for col in self.col_select:
            col.value(0)
        if column is not None:
            self.col_select[column].value(1)

    def drive_pixel(self, row, col, high_drive, energize_time_us=190):
        """Drive a single pixel"""
        self.set_row_address(row)
        self.set_column_address(col % self.COL_DRIVER_COLUMNS)

        self.set_row_direction(high_drive)
        self.set_column_direction(not high_drive)

        self.set_column_enabled(col // self.COL_DRIVER_COLUMNS)
        time.sleep_us(energize_time_us)
        self.set_column_enabled(None)

    def clear(self):
        """Clear the display buffer"""
        for row in range(self.ROWS):
            for col in range(self.COLS):
                self.buffer_active[row][col] = False

    def fill(self, value):
        """Fill the display buffer with a value"""
        for row in range(self.ROWS):
            for col in range(self.COLS):
                self.buffer_active[row][col] = value

    def set_pixel(self, row, col, value):
        """Set a pixel in the display buffer"""
        self.buffer_active[row][col] = value

    def refresh(self, force_refresh=False):
        """Refresh the display, updating changed pixels"""
        for row in range(self.ROWS):
            for col in range(self.COLS):
                if force_refresh or self.buffer_active[row][col] != self.buffer_shadow[row][col]:
                    self.drive_pixel(
                        row,
                        col,
                        not self.buffer_active[row][col]  # Invert because True = LOW drive
                    )
                    self.buffer_shadow[row][col] = self.buffer_active[row][col]
                    time.sleep_us(10)

# Example usage
def main():
    # Initialize the display
    display = LEDMatrix()
    led = Pin("LED", Pin.OUT)

    # Test pattern
    led.value(1)

    # Clear and deep refresh
    display.clear()
    display.refresh(force_refresh=True)
    time.sleep_ms(2000)

    # Blink pattern
    for _ in range(3):
        display.fill(True)
        display.refresh()
        time.sleep_ms(50)

        display.fill(False)
        display.refresh()
        time.sleep_ms(50)

    led.value(0)

    while True:
        pass

if __name__ == "__main__":
    main()

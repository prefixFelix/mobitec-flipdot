/// 2025 (C) Anton Berneving: https://github.com/antbern/flipdot-games
use core::convert::Infallible;

use cortex_m::delay::Delay;
use embedded_hal::digital::OutputPin;

/// The number of rows connected to each row and column driver chip.
const ROW_DRIVER_ROWS: u8 = 16;
const COL_DRIVER_COLUMNS: u8 = 28; // each chip has 24 outputs (7 segments * 4 digits)
pub struct Pins<'a> {
    pub row_address: [&'a mut dyn OutputPin<Error = Infallible>; 5],
    pub row_high_en: &'a mut dyn OutputPin<Error = Infallible>,
    pub row_low_en: &'a mut dyn OutputPin<Error = Infallible>,
    pub col_address: [&'a mut dyn OutputPin<Error = Infallible>; 5],
    pub col_high_low: &'a mut dyn OutputPin<Error = Infallible>,
    pub col_select: &'a mut [&'a mut dyn OutputPin<Error = Infallible>],
}

#[derive(Clone, Copy)]
pub enum DriveDirection {
    High,
    Low,
}
impl DriveDirection {
    fn opposite(&self) -> Self {
        match self {
            Self::High => Self::Low,
            Self::Low => Self::High,
        }
    }
}

impl Pins<'_> {
    fn set_row_address(&mut self, row: u8) {
        debug_assert!(row < 16); // the lookup table only supports 16 rows for now, although more are supported in hardware

        // define a lookup table that gives a bus address for each row index
        let lookup: [_; ROW_DRIVER_ROWS as usize] =
            [1, 4, 5, 2, 3, 6, 7, 9, 12, 13, 10, 11, 14, 15, 17, 20];

        let address = lookup[row as usize];

        for i in 0..5 {
            if (address & (1 << i)) == 0 {
                self.row_address[i].set_low().unwrap();
            } else {
                self.row_address[i].set_high().unwrap();
            }
        }
    }

    fn set_row_direction(&mut self, direction: DriveDirection) {
        match direction {
            DriveDirection::High => {
                self.row_low_en.set_low().unwrap();
                self.row_high_en.set_high().unwrap();
            }
            DriveDirection::Low => {
                self.row_high_en.set_low().unwrap();
                self.row_low_en.set_high().unwrap();
            }
        }
    }

    fn set_column_address(&mut self, column: u8) {
        debug_assert!(column < COL_DRIVER_COLUMNS); // each column driver chip can only do maximum 28 columns

        // the address space is not continous since it is based on 4 digits * 7 segments so this lookup table
        // basically skips each multiple of 7 + 1
        let lookup: [_; COL_DRIVER_COLUMNS as usize] = [
            1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 17, 18, 19, 20, 21, 22, 23, 25, 26, 27,
            28, 29, 30, 31,
        ];

        let address = lookup[column as usize];

        for i in 0..5 {
            if (address & (1 << i)) == 0 {
                self.col_address[i].set_low().unwrap();
            } else {
                self.col_address[i].set_high().unwrap();
            }
        }
    }

    fn set_column_direction(&mut self, direction: DriveDirection) {
        match direction {
            DriveDirection::High => {
                self.col_high_low.set_low().unwrap();
            }
            DriveDirection::Low => {
                self.col_high_low.set_high().unwrap();
            }
        }
    }

    fn set_column_enabled(&mut self, column: Option<u8>) {
        for c in self.col_select.iter_mut() {
            c.set_low().unwrap();
        }

        if let Some(column) = column {
            self.col_select[column as usize].set_high().unwrap();
        }
    }

    pub fn drive_pixel(
        &mut self,
        row: usize,
        col: usize,
        direction: DriveDirection,
        delay: &mut Delay,
        energize_time_us: u32,
    ) {
        // calculate the column and row driver and offsets

        // rows cannot be selected (never more than ROW_DRIVER_ROWS rows available)
        // debug_assert!(row < ROW_DRIVER_ROWS as usize);
        self.set_row_address(row as u8);

        // for columns, the address is the remainder
        self.set_column_address(col as u8 % COL_DRIVER_COLUMNS);

        self.set_row_direction(direction);
        self.set_column_direction(direction.opposite());

        // make sure time is not too high (above 25 ms) before allowing current flow
        debug_assert!(energize_time_us < 25 * 1000);

        // enabling the column driver actually enables the current flow
        self.set_column_enabled(Some(col as u8 / COL_DRIVER_COLUMNS));

        delay.delay_us(energize_time_us);

        // turn off the current flow by disabling all column drivers
        self.set_column_enabled(None);
    }
}

pub struct Display<'a, const ROWS: usize, const COLS: usize> {
    pins: Pins<'a>,
    buffer_shadow: [[bool; COLS]; ROWS],
    buffer_active: [[bool; COLS]; ROWS],
}

impl<const ROWS: usize, const COLS: usize> common::display::PixelDisplay
    for Display<'_, ROWS, COLS>
{
    const ROWS: usize = ROWS;
    const COLUMNS: usize = COLS;

    fn set_pixel(&mut self, row: usize, col: usize, value: common::display::Pixel) {
        self.set_pixel(
            row,
            col,
            match value {
                common::display::Pixel::On => true,
                common::display::Pixel::Off => false,
            },
        )
    }
}

impl<'a, const ROWS: usize, const COLS: usize> Display<'a, ROWS, COLS> {
    pub fn new(pins: Pins<'a>) -> Self {
        assert!(ROWS <= ROW_DRIVER_ROWS as usize);
        assert!(COLS <= COL_DRIVER_COLUMNS as usize * pins.col_select.len());

        Display {
            pins,
            buffer_shadow: [[false; COLS]; ROWS],
            buffer_active: [[false; COLS]; ROWS],
        }
    }

    pub fn clear(&mut self) {
        self.fill(false);
    }

    pub fn fill(&mut self, value: bool) {
        for row in 0..ROWS {
            for col in 0..COLS {
                self.buffer_active[row][col] = value;
            }
        }
    }

    pub fn set_pixel(&mut self, row: usize, col: usize, value: bool) {
        self.buffer_active[row][col] = value;
    }

    pub fn refresh(&mut self, delay: &mut Delay, force_refresh: bool) {
        for row in 0..ROWS {
            for col in 0..COLS {
                if force_refresh || self.buffer_active[row][col] != self.buffer_shadow[row][col] {
                    // need to update the display pixel
                    self.pins.drive_pixel(
                        row,
                        col,
                        if self.buffer_active[row][col] {
                            DriveDirection::Low
                        } else {
                            DriveDirection::High
                        },
                        delay,
                        190,
                    );

                    self.buffer_shadow[row][col] = self.buffer_active[row][col];

                    delay.delay_us(10);
                }
            }
        }
    }
}

import numpy as np


class Bitmap:
    """
    Basic bitmap objects. Gets queued in the image buffer.
    Attributes:
        width (byte): Bitmap width.
        height (byte): Bitmap height.
        bitmap (list of lists): Bitmap. Adressed like this: bitmap[y][x]
    """
    def __init__(self, width, height):
        self.width = width
        self.height = height
        self.bitmap = np.zeros((height, width), dtype=bool)
        self.x_pos = 0
        self.y_pos = 0

    def convert_to_sub_column(self):
        """Converts a regular bitmap to subcolumn matrix."""
        sub_column = []
        bm_t = np.transpose(self.bitmap)
        for i in range(0, self.height, 5):
            end_index = min(i + 5, self.height)
            sub_column.append(bm_t[:, i:end_index])
        return sub_column


    def fill(self, color):
        self.bitmap = np.full_like(self.bitmap, color)

    def invert(self):
        self.bitmap = np.logical_not(self.bitmap).astype(dtype=bool)

    def shift_left(self, n, keep=False):
        shifted_bm = np.roll(self.bitmap, -n, axis=1)
        if keep:
            shifted_bm[:, -n:] = self.bitmap[:, :n]
        else:
            shifted_bm[:, -n:] = 0
        self.bitmap = shifted_bm

    def shift_right(self, n, keep=False):
        shifted_bm = np.roll(self.bitmap, n, axis=1)
        if keep:
            shifted_bm[:, n:] = self.bitmap[:, :-n]
        else:
            shifted_bm[:, n:] = 0
        self.bitmap = shifted_bm

    def line_horizontal(self, x_pos, y_pos, length, width, color):
        self.bitmap[y_pos:y_pos + width, x_pos:x_pos + length] = color

    def line_vertical(self, x_pos, y_pos, length, width, color):
        self.bitmap[y_pos:y_pos + length, x_pos:x_pos + width] = color

    def line_diagonal_right(self, x_pos, y_pos, length, width, color):
        for i in range(length):
            for j in range(width):
                self.bitmap[(y_pos + j) + i, x_pos + i] = color

    def line_diagonal_left(self, x_pos, y_pos, length, width, color):
        for i in range(length):
            for j in range(width):
                self.bitmap[(y_pos + j) + i, x_pos - i] = color

    # Not working
    def rectangle(self, x_pos, y_pos, height, width, color):
        self.bitmap[y_pos + 1:y_pos + height - 1, x_pos + 1:x_pos + width - 1] = color

    def circle(self, x_pos, y_pos, radius, color):
        print('todo')

    def dot(self, x_pos, y_pos, color):
        self.bitmap[x_pos, y_pos] = color

    def random(self):
        print('todo')
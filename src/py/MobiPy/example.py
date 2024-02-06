import mobitec
import font
import bitmap
import symbol

""" DEV TEST """

flipdot = mobitec.MobitecDisplay('/dev/ttyUSB0', address=0x0b, width=28, height=16)

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
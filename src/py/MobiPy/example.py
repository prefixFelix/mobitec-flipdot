import time
import numpy as np
import mobitec
import font
import bitmap
import symbol

flipdot = mobitec.MobitecDisplay('/dev/ttyUSB0', address=0x0b, width=28, height=16)

""" Text """
flipdot.text('1337', font.F5, 5, 5)
time.sleep(1)
flipdot.text('test', font.F9)
time.sleep(1)

""" Symbol """
flipdot.symbol(symbol.SOCCER1)
time.sleep(1)

""" Bitmap """
test_sape = [[1, 0, 0, 1],
             [1, 0, 0, 1],
             [1, 0, 0, 1],
             [1, 1, 1, 1]]
bm = bitmap.Bitmap(4, 4)
bm.bitmap = np.array(test_sape)
flipdot.bitmap(bm, 4, 4)
time.sleep(1)

bm = bitmap.Bitmap(28, 16)
flipdot.bitmap(bm)
time.sleep(1)
bm.invert()
flipdot.bitmap(bm)
time.sleep(1)

bm.fill(0)
bm.line_horizontal(1,1, 10, 2, 1)
bm.line_vertical(1, 1, 10, 2, 1)
flipdot.bitmap(bm)
time.sleep(1)

bm.line_diagonal_right(2, 2, 10, 2, 1)
bm.dot(14, 14, 1)
flipdot.bitmap(bm)
time.sleep(1)

for i in range(28):
    bm.shift_left(1, keep=True)
    flipdot.bitmap(bm)
    time.sleep(1)

""" Buffer """
flipdot.buffer_text('69', font.F5, 0, 5)
test_sape = [[1, 1, 0, 0],
             [0, 1, 0, 1],
             [1, 0, 1, 1],
             [0, 1, 1, 0]]
bm = bitmap.Bitmap(4, 4)
flipdot.buffer_bitmap(bm, 14, 5)
flipdot.display()
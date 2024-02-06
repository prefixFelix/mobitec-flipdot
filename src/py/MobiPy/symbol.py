import font

class Symbol:
    """
    Basic symbol objects.
    Attributes:
    """
    def __init__(self, value):
        self.value = value
        self.font = font.Font(0, 0x67)


CHURCH = Symbol('1')
SOCCER1 = Symbol('2')
SOCCER2 = Symbol('3')
# todo add more
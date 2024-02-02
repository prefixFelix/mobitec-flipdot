"""
Special symbols:
    text_5px = 0x72  # Large letters only
    text_6px = 0x66
    text_7px = 0x65
    text_7px_bold = 0x64
    text_9px = 0x75
    text_9px_bold = 0x70
    text_9px_bolder = 0x62
    text_13px = 0x73
    text_13px_bold = 0x69
    text_13px_bolder = 0x61
    text_13px_boldest = 0x79
    numbers_14px = 0x00
    text_15px = 0x71
    text_16px = 0x68
    text_16px_bold = 0x78
    text_16px_bolder = 0x74
    symbols = 0x67
    bitmap = 0x77

    F61 = (0x61, 5)
    F62 = (0x62, 5)
    F63 = (0x63, 5)
    F64 = (0x64, 5)
    F65 = (0x65, 5)
    F66 = (0x66, 5)
    F67 = (0x67, 5)
    F68 = (0x68, 5)
    F69 = (0x69, 5)
    F70 = (0x70, 5)
    F71 = (0x71, 5)
    F72 = (0x72, 5)
    F73 = (0x73, 5)
    F74 = (0x74, 5)
    F75 = (0x75, 5)
    F76 = (0x76, 5)
    SMALL_F = (0x64, 5)
    BITMAP = (0x77, 5)

    # F<HEIGHT>_<Fat>/<Thin>

    F13_F = (0x61, 13) # on 0x70 too
    F9_F = (0x62, 9)
    F19_F = (0x63, 19)
    F7_F = (0x64, 7)
    F7 = (0x65, 7)
    F6 = (0x66, 7)
    SYMBOL = (0x67, 16)
    F16_T = (0x68, 16)
    F13_T = (0x69, 13)
    #F13_F = (0x70, 13)
    F15_T = (0x71, 15)
    F5 =(0x72, 5)
    F13_TT = (0x73, 13)

    # 74: Only one character... and only A seems to work
    F13 = (0x75, 13) # same as F13_T but wider...
"""
fonts = {
        # name, height, code
        "7px": Font("7px", 7, 0x60),
        "7px_wide": Font("7px_wide", 7, 0x62),
        "12px": Font("12px", 12, 0x63),
        "13px": Font("13px", 13, 0x64),
        "13px_wide": Font("13px_wide", 13, 0x65),
        "13px_wider": Font("13px_wider", 13, 0x69),
        "16px_numbers": Font("16px_numbers", 16, 0x68),
        "16px_numbers_wide": Font("16px_numbers_wide", 16, 0x6a),
        "pixel_subcolumns": Font("pixel_subcolumns", 5, 0x77)
    }
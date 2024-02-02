package fonts

const (
	Font_7px               = "7px"
	Font_7px_wide          = "7px_wide"
	Font_12px              = "12px"
	Font_13px              = "13px"
	Font_13px_wide         = "13px_wide"
	Font_13px_wider        = "13px_wider"
	Font_16px_numbers      = "16px_numbers"
	Font_16px_numbers_wide = "16px_numbers_wide"
	Font_pixel_subcolumns  = "pixel_subcolumns"
)

var (
	FONTS = Fonts{
		Font_7px:               {7, 0x60},
		Font_7px_wide:          {7, 0x62},
		Font_12px:              {12, 0x63},
		Font_13px:              {13, 0x64},
		Font_13px_wide:         {13, 0x65},
		Font_13px_wider:        {13, 0x69},
		Font_16px_numbers:      {16, 0x68},
		Font_16px_numbers_wide: {16, 0x6a},
		Font_pixel_subcolumns:  {5, 0x77},
	}
	FONTDEFAULT = Font{5, 0x77}
	CHARMAP     = map[rune]byte{
		'Ä': 0x5b, 'ä': 0x7b,
		'Ö': 0x5c, 'ö': 0x7c,
	}
)

type Font struct {
	Height int
	Code   byte
}

type Fonts map[string]Font

func (fs Fonts) Get(name string) Font {
	if f, ok := fs[name]; ok {
		return f
	}
	return FONTDEFAULT
}

package cfg

const (
	WIDTH  = 28 * 4
	HEIGHT = 16
	PORT   = "COM4"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

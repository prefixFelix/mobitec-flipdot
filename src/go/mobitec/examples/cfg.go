package cfg

import "time"

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

func Wait() {
	time.Sleep(750 * WIDTH / 28 * time.Millisecond)
}

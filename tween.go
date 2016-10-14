package newyeartree

import (
	"log"
	"time"
)

func TweenTo(o *Pixel, to Pixel, d time.Duration) {
	log.Println("stop tween")

	b := time.Now().UnixNano()

	from := o.Clone()
	log.Println("to.Blue", to.Blue)

	for {
		ratio := float64((time.Now().UnixNano() - b)) / float64(d)

		if ratio >= 1 {
			return
		}

		o.X = int(int32(from.X) + int32(round(ratio*float64(to.X-from.X))))
		o.X = int(int32(from.Y) + int32(round(ratio*float64(to.Y-from.Y))))
		log.Println(o.X, o.Y)

		o.Color.Red = uint32(int32(from.Red) + int32(ratio*float64(int32(to.Red)-int32(from.Red))))
		o.Color.Green = uint32(int32(from.Green) + int32(ratio*float64(int32(to.Green)-int32(from.Green))))
		o.Color.Blue = uint32(int32(from.Blue) + int32(ratio*float64(int32(to.Blue)-int32(from.Blue))))

		time.Sleep(10 * time.Millisecond)
	}

}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

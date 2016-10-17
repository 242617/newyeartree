package newyeartree

import "time"

func TweenTo(o *Pixel, to Pixel, d time.Duration) {
	start := time.Now().UnixNano()
	from := o.Clone()

	for {
		ratio := float64((time.Now().UnixNano() - start)) / float64(d)

		if ratio >= 1 {
			o.X, o.Y = to.X, to.Y
			o.Red, o.Green, o.Blue = to.Red, to.Green, to.Blue
			return
		}

		o.X = int(int32(from.X) + int32(round(ratio*float64(to.X-from.X))))
		o.Y = int(int32(from.Y) + int32(round(ratio*float64(to.Y-from.Y))))
		o.Color.Red = uint32(int32(from.Red) + int32(ratio*float64(int32(to.Red)-int32(from.Red))))
		o.Color.Green = uint32(int32(from.Green) + int32(ratio*float64(int32(to.Green)-int32(from.Green))))
		o.Color.Blue = uint32(int32(from.Blue) + int32(ratio*float64(int32(to.Blue)-int32(from.Blue))))

		time.Sleep(20 * time.Millisecond)
	}

}

func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

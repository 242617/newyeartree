
Example

	package main

	import (
		"flag"
		"log"
		"time"

		"github.com/stylerucom/newyeartree"
	)

	func main() {
		log.SetFlags(log.Lshortfile)

		colors := map[string]newyeartree.Color{
			"black":  newyeartree.Color{0, 0, 0},
			"white":  newyeartree.Color{255, 255, 255},
			"red":    newyeartree.Color{255, 0, 0},
			"green":  newyeartree.Color{0, 255, 0},
			"orange": newyeartree.Color{255, 69, 0},
			"yellow": newyeartree.Color{255, 255, 0},
			"cyan":   newyeartree.Color{0, 255, 255},
		}

		port := flag.String("port", "COM4", "Path to serial port")
		width := flag.Int("width", 2, "LED width")
		height := flag.Int("height", 2, "LED height")
		brightness := flag.Uint64("brightness", 128, "Tree brightness")
		framerate := flag.Int("framerate", 60, "Animation framerate")

		t := newyeartree.Tree{
			Framerate: *framerate,
			Height:    *height,
			Port:      *port,
			Width:     *width,
		}
		t.Start()
		t.SetBrightness(uint32(*brightness))

		red := newyeartree.Pixel{X: 0, Y: 0, Color: colors["red"]}
		green := newyeartree.Pixel{X: 1, Y: 1, Color: colors["green"]}
		t.AddPixel(&red)

		newyeartree.TweenTo(&red, green, 10*time.Second)
	}

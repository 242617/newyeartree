Example

	package main

	import (
		"flag"
		"log"
		"time"

		. "github.com/stylerucom/newyeartree"
	)

	func main() {
		log.SetFlags(log.Lshortfile)

		colors := map[string]Color{
			"black":  Color{0, 0, 0},
			"white":  Color{255, 255, 255},
			"red":    Color{255, 0, 0},
			"green":  Color{0, 255, 0},
			"orange": Color{255, 69, 0},
			"yellow": Color{255, 255, 0},
			"cyan":   Color{0, 255, 255},
		}

		port := flag.String("port", "COM4", "Path to serial port")
		width := flag.Int("width", 2, "LED width")
		height := flag.Int("height", 2, "LED height")
		brightness := flag.Uint64("brightness", 128, "Tree brightness")
		framerate := flag.Int("framerate", 60, "Animation framerate")

		s := NewStrip(Options{
			Framerate: *framerate,
			Height:    *height,
			Port:      *port,
			Width:     *width,
		})
		s.Start()
		s.SetBrightness(uint32(*brightness))

		mtx := [][]Color{
			[]Color{colors["red"], colors["orange"]},
			[]Color{colors["green"], colors["cyan"]},
		}

		for y := 0; y < *height; y++ {
			for x := 0; x < *width; x++ {
				p := Pixel{X: x, Y: y}
				s.AddPixel(&p)
				go func(p *Pixel, c Color) {
					for {
						TweenTo(p, Pixel{X: p.X, Y: p.Y, Color: c}, 2*time.Second)
						TweenTo(p, Pixel{X: p.X, Y: p.Y, Color: colors["black"]}, 2*time.Second)
					}
				}(&p, mtx[y][x])
			}
		}

		for {
		}

	}
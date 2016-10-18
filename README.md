Example. Blink with all 4 LEDs.

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
		brightness := flag.Uint64("brightness", 128, "Brightness")
		framerate := flag.Int("framerate", 60, "Animation framerate")
		flag.Parse()
		log.Printf("Options: {port: %s, dimensions: [%d:%d], brightness: %d, framerate: %d}", *port, *width, *height, *brightness, *framerate)

		s := NewStrip(Options{
			Framerate: *framerate,
			Height:    *height,
			Port:      *port,
			Width:     *width,
		})
		s.Start()

		for y := 0; y < *height; y++ {
			for x := 0; x < *width; x++ {
				p := NewPixel(x, y, colors["red"])
				s.AddPixel(&p)
			}
		}

		go func() {
			for {
				s.Brightness = 0
				time.Sleep(200 * time.Millisecond)
				s.Brightness = uint32(*brightness)
				time.Sleep(200 * time.Millisecond)
			}
		}()

		select {}

	}

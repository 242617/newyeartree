package newyeartree

import (
	"log"
	"sync"
	"time"

	"github.com/tarm/serial"
)

type Tree struct {
	Framerate  int
	Height     int
	Port       string
	Width      int
	brightness uint32
	pixels     []*Pixel
	delay      time.Duration
	serialPort *serial.Port
	sync.Mutex
}

func (t *Tree) Start() {
	log.Println("Start")
	var err error

	t.delay = time.Duration(1000000000 / t.Framerate)

	c := &serial.Config{Name: t.Port, Baud: 115200}
	t.serialPort, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	go func() {
		for {
			t.run()
			time.Sleep(t.delay)
		}
	}()
}

func (t *Tree) AddPixel(o *Pixel) {
	t.pixels = append(t.pixels, o)
}

func (t *Tree) run() {

	data := make([]Color, t.Width*t.Height)

	for y := 0; y < t.Height; y++ {
		for x := 0; x < t.Width; x++ {
			i := t.Width*y + x
			for _, v := range t.pixels {
				data[i].Mix(v.Get(x, y))
			}
		}
	}

	for k, v := range data {
		t.writeData([]byte{byte(k), byte(v.Red), byte(v.Green), byte(v.Blue)})
	}
}

func (t Tree) setColor(n int, c Color, b uint32) {
	r := c.Clone()
	r.SetBrightness(b)
}

func (t Tree) writeData(d []byte) {
	// t.Lock()
	// defer t.Unlock()

	_, err := t.serialPort.Write(d)
	if err != nil {
		log.Fatal(err)
	}

	/*buf := make([]byte, 256)
	_, err = t.serialPort.Read(buf)
	if err != nil {
		log.Fatal(err)
	}*/
}

func (t *Tree) SetBrightness(b uint32) {
	t.brightness = b
}

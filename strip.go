package newyeartree

import (
	"log"
	"sync"
	"time"

	"github.com/tarm/serial"
)

func NewStrip(o Options) (s Strip) {
	s.framerate = o.Framerate
	s.height = o.Height
	s.port = o.Port
	s.width = o.Width
	return
}

type Options struct {
	Framerate int
	Height    int
	Port      string
	Width     int
}
type Strip struct {
	brightness uint32
	delay      time.Duration
	framerate  int
	height     int
	pixels     []*Pixel
	port       string
	serialPort *serial.Port
	width      int
	sync.Mutex
}

func (s *Strip) Start() {
	log.Println("Start")
	var err error

	s.delay = time.Duration(1000000000 / s.framerate)

	c := &serial.Config{Name: s.port, Baud: 115200}
	s.serialPort, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	go func() {
		for {
			s.run()
			time.Sleep(s.delay)
		}
	}()
}

func (s *Strip) AddPixel(o *Pixel) {
	s.pixels = append(s.pixels, o)
}

func (s *Strip) run() {

	data := make([]Color, s.width*s.height)

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			i := s.width*y + x
			for _, v := range s.pixels {
				data[i].Mix(v.Get(x, y))
			}
		}
	}

	for k, v := range data {
		s.writeData([]byte{byte(k), byte(v.Red), byte(v.Green), byte(v.Blue)})
	}
}

func (s Strip) setColor(n int, c Color, b uint32) {
	r := c.Clone()
	r.SetBrightness(b)
}

func (s Strip) writeData(d []byte) {
	s.Lock()
	defer s.Unlock()

	_, err := s.serialPort.Write(d)
	if err != nil {
		log.Fatal(err)
	}

	/*buf := make([]byte, 256)
	_, err = s.serialPors.Read(buf)
	if err != nil {
		log.Fatal(err)
	}*/
}

func (t *Strip) SetBrightness(b uint32) {
	t.brightness = b
}

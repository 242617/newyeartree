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

type Strip struct {
	brightness uint32
	delay      time.Duration
	framerate  int
	height     int
	pixels     map[*Pixel]Pixel
	port       string
	serialPort *serial.Port
	width      int
	sync.Mutex
}

func (s *Strip) Start() {

	s.delay = time.Duration(1000000000 / s.framerate)

	log.Println("Connecting...")
	c := &serial.Config{
		Name:        s.port,
		Baud:        115200,
		ReadTimeout: time.Second * 5,
	}
	var err error
	s.serialPort, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	log.Println("Start")
	s.pixels = map[*Pixel]Pixel{}

	go func() {
		for {
			s.run()
			time.Sleep(s.delay)
		}
	}()
}

func (s *Strip) AddPixel(o *Pixel) {
	s.pixels[o] = *o
}

func (s *Strip) RemovePixel(o *Pixel) {
	delete(s.pixels, o)
}

func (s *Strip) SetBrightness(b uint32) {
	s.brightness = b
}

func (s *Strip) run() {
	data := make([]Color, s.width*s.height)

	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			i := s.width*y + x
			for k, _ := range s.pixels {
				data[i].Mix((*k).Get(x, y))
			}
		}
	}

	for k, v := range data {
		v.Red = v.Red * s.brightness >> 8
		v.Green = v.Green * s.brightness >> 8
		v.Blue = v.Blue * s.brightness >> 8
		s.writeData([]byte{byte(k), byte(v.Red), byte(v.Green), byte(v.Blue)})
	}
}

func (s *Strip) writeData(d []byte) {
	s.Lock()
	defer s.Unlock()

	_, err := s.serialPort.Write(d)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 256)
	_, err = s.serialPort.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
}

// Terminal based mandelbrot set explorer.
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/encoding"
)

var style = tcell.StyleDefault
var quit chan struct{}

func main() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	encoding.Register()

	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorWhite).
		Background(tcell.ColorBlack))
	s.EnableMouse()
	s.Clear()

	quit = make(chan struct{})
	go pollEvents(s)

	s.Show()

	go func() {
		for {
			drawScreen(s)
			time.Sleep(time.Millisecond * 10)
		}
	}()

	<-quit
	s.Fini()
}

type viewport struct {
	x0, x1, y0, y1 float64
}

func NewViewport(x0, x1, y0, y1 float64) viewport {
	vp := viewport{}
	vp.x0 = x0
	vp.x1 = x1
	vp.y0 = y0
	vp.y1 = y1

	return vp
}

var vp = NewViewport(-2.0, 1.0, -1.0, 1.0)

func zoom(s tcell.Screen, direction, x, y int) {
	//	w, h := s.Size()

	factorx := (vp.x0 - vp.x1) / 10.0
	factory := (vp.y0 - vp.y1) / 10.0

	if direction == 1 {
		vp.x0 -= factorx
		vp.x1 += factorx
		vp.y0 -= factory
		vp.y1 += factory
	} else {
		vp.x0 += factorx
		vp.x1 -= factorx
		vp.y0 += factory
		vp.y1 -= factory
	}
}

func drawScreen(s tcell.Screen) {
	w, h := s.Size()

	if w == 0 || h == 0 {
		return
	}

	st := tcell.StyleDefault
	const gl = '▄'
	//	const gl = ' '

	for x := 0; x < w; x++ {
		for y := 0; y < h*2; y += 1 {
			r := mapnum(x, 0, w, vp.x0, vp.x1)
			i := mapnum(y, 0, h*2, vp.y0, vp.y1)
			i2 := mapnum(y+1, 0, h*2, vp.y0, vp.y1)

			converges, iter := calc_spot(r, i)
			if converges {
				st = st.Background(tcell.Color(232))
			} else {
				st = st.Background(asColor(iter))
			}

			converges, iter = calc_spot(r, i2)
			if converges {
				st = st.Foreground(tcell.Color(232))
			} else {
				st = st.Foreground(asColor(iter))
			}

			s.SetCell(x, y/2, st, gl)
		}
	}

	s.Show()
}

func calc_spot(ca, cb float64) (bool, int) {
	var a, b float64 = 0, 0
	max := 1000
	var i int
	for i = 0; i < max; i++ {
		as, bs := a*a, b*b
		if as+bs > 16 {
			return false, i
		}
		a, b = as-bs+ca, 2*a*b+cb
	}
	return true, i
}

func asColor(n int) tcell.Color {
	color := tcell.Color((16 + n) % 229)
	return color
}

func mapnum(x, in_min, in_max int, out_min, out_max float64) float64 {
	return (float64(x)-float64(in_min))*(out_max-out_min)/(float64(in_max)-float64(in_min)) + out_min
}

func pollEvents(s tcell.Screen) {
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				close(quit)
				return
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'z':
					zoom(s, 1, 1, 1)
				case 'x':
					zoom(s, 0, 1, 1)
				case 'q':
					close(quit)
					return
				}
				//s.Sync()
			case tcell.KeyUp:
				step := (vp.y0 - vp.y1) / 10
				vp.y0 += step
				vp.y1 += step
			case tcell.KeyDown:
				step := (vp.y0 - vp.y1) / 10
				vp.y0 -= step
				vp.y1 -= step
			case tcell.KeyLeft:
				step := (vp.x0 - vp.x1) / 10
				vp.x0 += step
				vp.x1 += step
			case tcell.KeyRight:
				step := (vp.x0 - vp.x1) / 10
				vp.x0 -= step
				vp.x1 -= step
			}
		case *tcell.EventMouse:
			x, y := ev.Position()
			button := ev.Buttons()
			/*if button&tcell.WheelUp != 0 {
				bstr += " WheelUp"
			}*/
			// Only buttons, not wheel events
			button &= tcell.ButtonMask(0xff)
			switch ev.Buttons() {
			case tcell.Button1:
				zoom(s, 1, x, y)
			case tcell.Button2:
				zoom(s, 0, x, y)
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}

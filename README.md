# frac`tty`

An interactive Mandelbrot set explorer for your teletype terminal.
Or more likely pseudo-TTY.

## Why

1. Terminals are cool, and we need to teach them new tricks.
2. Fractals are beautiful and I'd like to understand them better.
3. Practice.

## Thanks & Inspiration

* [Benoit Mandelbrot][mb] (20 November 1924 – 14 October 2010)
* [gdamore/tcell][tcell] for the excellent terminal library
* [12-minute Mandelbrot: fractals on a 50 year old IBM 1401 mainframe][ken_shirriff]

[tcell]: https://github.com/gdamore/tcell
[mb]: https://en.wikipedia.org/wiki/Benoit_Mandelbrot
[ken_shirriff]: http://www.righto.com/2015/03/12-minute-mandelbrot-fractals-on-50.html

## Usage

Use the `z` and `x` keys to zoom in and out.
Use arrow keys for navigation and `q` to quit.
The mouse wheel zooms, but panning around using click-and-drag
is currently broken.

## Gallery

![m1](img/mandelb1.png?raw=true "m1")
![m2](img/mandelb2.png?raw=true "m2")
![m3](img/mandelb3.png?raw=true "m3")
![m4](img/mandelb4.png?raw=true "m4")

## FAQ

### It looks kind of low-resolution and crappy

* Thanks!  That's not a question though.  The resolution is limited  
to the number of columns on the terminal and twice the number of
rows.  Vertical resolution is doubled by using the half-block
character "▄" (U+2584), so an old-school 80 column, 24 row terminal
window is effectively a measely 80x48 pixels.  You can increase
the resolution by shrinking the font in the terminal window.

### Why is it slow weird and broken

* I only barely know what I'm doing!

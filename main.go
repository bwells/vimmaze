package main

import (
	"fmt"

	"github.com/bwells/vimmaze/maze"
	"github.com/nsf/termbox-go"
)

const fg = termbox.ColorBlack
const bg = termbox.ColorWhite

func handleKey(m maze.Maze, x, y int, ev termbox.Event) (int, int, error) {
	switch ev.Ch {
	case 'h':
		if m.IsValidMove(x, y, maze.LEFT) {
			return x - 1, y, nil
		} else {
			return -1, -1, fmt.Errorf("Inalid Move")
		}

	case 'j':
		if m.IsValidMove(x, y, maze.DOWN) {
			return x, y + 1, nil
		} else {
			return -1, -1, fmt.Errorf("Inalid Move")
		}

	case 'k':
		if m.IsValidMove(x, y, maze.UP) {
			return x, y - 1, nil
		} else {
			return -1, -1, fmt.Errorf("Inalid Move")
		}

	case 'l':
		if m.IsValidMove(x, y, maze.RIGHT) {
			return x + 1, y, nil
		} else {
			return -1, -1, fmt.Errorf("Inalid Move")
		}
	}
	return -1, -1, fmt.Errorf("Key not bound")
}

func drawMaze(m maze.Maze) {

	var c rune

	for x := 0; x < m.Width*2+1; x++ {
		termbox.SetCell(x, 0, '_', fg, bg)
	}

	for y := 0; y < m.Height; y++ {

		termbox.SetCell(0, y+1, '|', fg, bg)

		for x := 0; x < m.Width; x++ {
			cell := m.GetAt(x, y)

			if cell.Bottom {
				c = ' '
			} else {
				c = '_'
			}
			termbox.SetCell(x*2+1, y+1, c, fg, bg)

			if cell.Right {
				c = ' '
			} else {
				c = '|'
			}
			termbox.SetCell(x*2+2, y+1, c, fg, bg)
		}

	}
	termbox.Flush()
}

func drawLocation(m maze.Maze, x, y int) {
	// TODO: Using cursor as * hides a wall below
	// the current position
	termbox.SetCell(x*2+1, y+1, '*', fg, bg)
}

func clearLocation(m maze.Maze, x, y int) {
	var c rune
	cell := m.GetAt(x, y)
	if !cell.Bottom {
		c = '_'
	} else {
		c = ' '
	}
	termbox.SetCell(x*2+1, y+1, c, fg, bg)
}

func main() {

	err := termbox.Init()
	if err != nil {
		panic(err)

	}
	defer termbox.Close()

	m := maze.NewMaze(80, 40)

	x, y := 0, 0

	drawMaze(m)
	drawLocation(m, x, y)
	termbox.Flush()

mainLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				break mainLoop
			}
			new_x, new_y, err := handleKey(m, x, y, ev)
			if err != nil {
				continue
			}
			clearLocation(m, x, y)
			drawLocation(m, new_x, new_y)
			termbox.Flush()
			x, y = new_x, new_y
		case termbox.EventError:
			panic(ev.Err)
		}
	}

}

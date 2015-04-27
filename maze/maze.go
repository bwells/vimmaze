package maze

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

type Direction uint8

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

type Cell struct {
	Top    bool
	Right  bool
	Bottom bool
	Left   bool
}

type Maze struct {
	Width  int
	Height int
	cells  []Cell
}

func NewMaze(width, height int) Maze {
	m := Maze{Width: width, Height: height}
	m.cells = make([]Cell, width*height)

	m.generate()

	return m
}

func (m Maze) generate() {
	// We treat each cell as a disjoint set
	// we keep removing a random wall until
	// we are left with only one set

	// populate the matrix with each cell as a unique set
	sets := make([]int, len(m.cells))
	for i := 0; i < len(sets); i++ {
		sets[i] = i
	}

	rand.Seed(time.Now().UnixNano())

	distinct_sets := len(sets)
	for distinct_sets != 1 {
		// get two random adjoining cell locations
		index1 := rand.Intn(len(sets))
		index2 := m.getRandomNeighbor(index1)

		// if the cells are already in the same set
		// then start over
		if sets[index1] == sets[index2] {
			continue
		}

		// they aren't already in the same set
		// so remove the wall between them
		// and update the sets to mark that they are
		m.removeWall(index1, index2)
		new_set := sets[index1]
		old_set := sets[index2]
		for i := 0; i < len(sets); i++ {
			if sets[i] == old_set {
				sets[i] = new_set
			}
		}

		distinct_sets -= 1
	}
}

func (m Maze) getRandomNeighbor(index int) int {
	x, y := m.getXY(index)
	original_x, original_y := x, y
	rand.Seed(time.Now().UnixNano())
	for {
		x, y = original_x, original_y
		if rand.Intn(2) == 1 {
			if rand.Intn(2) == 1 {
				x += 1
			} else {
				x -= 1
			}
		} else {
			if rand.Intn(2) == 1 {
				y += 1
			} else {
				y -= 1
			}
		}
		if x >= 0 && x < m.Width && y >= 0 && y < m.Height {
			return m.getIndex(x, y)
		}

	}
}

func (m Maze) removeWall(index1, index2 int) {
	// swap indexes to simplify states to check as necessary
	if index2 < index1 {
		index1, index2 = index2, index1
	}

	// if this is a left-right neighbor
	if index2-index1 == 1 {
		m.cells[index1].Right = true
		m.cells[index2].Left = true
		// else top-bottom neighbor
	} else {
		m.cells[index1].Bottom = true
		m.cells[index2].Top = true
	}
}

func (m Maze) getXY(index int) (int, int) {
	y := int(math.Floor(float64(index) / float64(m.Width)))
	x := index - m.Width*y
	return x, y
}

func (m Maze) getIndex(x, y int) int {
	return m.Width*y + x
}

// 0,  1,  2,  3,  4
// 5,  6,  7,  8,  9
// 10, 11, 12, 13, 14
// 15, 16, 17, 18, 19
// 20, 21, 22, 23, 24
// 25, 26, 27, 28, 29
// 30, 31, 32, 33, 34
// 35, 36, 37, 38, 39
// 40, 41, 42, 43, 44
// 45, 46, 47, 48, 49
//
// 3,2 = 13 -> 5 * 2 + 3 == 13

// Getting neighbor for 46 -> (4, 7)
// Reject neighbor (5, 8)
// Reject neighbor (4, 9)
// Reject neighbor (5, 8)
// Reject neighbor (4, 9)

func (m Maze) GetAt(x, y int) Cell {
	return m.cells[m.getIndex(x, y)]
}

func (m Maze) setTopAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].Top = !wall
}

func (m Maze) setRightAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].Right = !wall
}

func (m Maze) setBottomAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].Bottom = !wall
}

func (m Maze) setLeftAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].Left = !wall
}

func (m Maze) String() string {
	var buffer bytes.Buffer

	for i := 0; i < m.Width*2+1; i++ {
		buffer.WriteString("_")
	}

	buffer.WriteString("\n")

	for y := 0; y < m.Height; y++ {

		buffer.WriteString("|")

		for x := 0; x < m.Width; x++ {
			index := m.getIndex(x, y)
			if m.cells[index].Bottom {
				buffer.WriteString(" ")
			} else {
				buffer.WriteString("_")
			}
			if m.cells[index].Right {
				buffer.WriteString(" ")
			} else {
				buffer.WriteString("|")
			}

		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

func (m Maze) IsValidMove(x, y int, dir Direction) bool {

	var new_x, new_y int

	switch dir {
	case LEFT:
		new_x, new_y = x-1, y
	case RIGHT:
		new_x, new_y = x+1, y
	case UP:
		new_x, new_y = x, y-1
	case DOWN:
		new_x, new_y = x, y+1
	}

	// not valid if it will move out of bounds
	if new_x < 0 || new_y < 0 || new_x > m.Width || new_y > m.Height {
		return false
	}

	cell := m.GetAt(x, y)
	// cell direction value == false means there is a wall there
	switch dir {
	case LEFT:
		if !cell.Left {
			return false
		}
	case RIGHT:
		if !cell.Right {
			return false
		}
	case UP:
		if !cell.Top {
			return false
		}
	case DOWN:
		if !cell.Bottom {
			return false
		}
	}

	return true
}

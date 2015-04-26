package maze

import (
	"bytes"
	"math"
	"math/rand"
	"time"
)

type cell struct {
	top    bool
	right  bool
	bottom bool
	left   bool
}

type Maze struct {
	width  int
	height int
	cells  []cell
}

func NewMaze(width, height int) Maze {
	m := Maze{width: width, height: height}
	m.cells = make([]cell, width*height)

	m.generate()

	return m
}

func (m Maze) generate() {
	// We treat each cell as a disjoint set
	// we keep removing a random wall until
	// we are left with only one set

	// each array value in Maze.conent represents a wall
	//     we need +1 for the outer walls, - fence posts
	// each array value in sets represents a cell

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
		if x >= 0 && x < m.width && y >= 0 && y < m.height {
			return m.getIndex(x, y)
		}

	}
}

func (m Maze) removeWall(index1, index2 int) {
	if index2 < index1 {
		index1, index2 = index2, index1
	}

	// case: left and right
	//       2 and 3:
	// case: top and bottom
	//       2 and 7:

	// if this is a left-right neighbor
	if index2-index1 == 1 {
		m.cells[index1].right = true
		m.cells[index2].left = true
	} else {
		m.cells[index1].bottom = true
		m.cells[index2].top = true
	}
}

func (m Maze) getXY(index int) (int, int) {
	y := int(math.Floor(float64(index) / float64(m.width)))
	x := index - m.width*y
	return x, y
}

func (m Maze) getIndex(x, y int) int {
	return m.width*y + x
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

func (m Maze) getAt(x, y int) cell {
	return m.cells[m.getIndex(x, y)]
}

func (m Maze) setTopAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].top = !wall
}

func (m Maze) setRightAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].right = !wall
}

func (m Maze) setBottomAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].bottom = !wall
}

func (m Maze) setLeftAt(x, y int, wall bool) {
	m.cells[m.getIndex(x, y)].left = !wall
}

func (m Maze) String() string {
	var buffer bytes.Buffer

	for i := 0; i < m.width*2+1; i++ {
		buffer.WriteString("_")
	}

	buffer.WriteString("\n")

	for y := 0; y < m.height; y++ {

		buffer.WriteString("|")

		for x := 0; x < m.width; x++ {
			index := m.getIndex(x, y)
			if m.cells[index].bottom {
				buffer.WriteString(" ")
			} else {
				buffer.WriteString("_")
			}
			if m.cells[index].right {
				buffer.WriteString(" ")
			} else {
				buffer.WriteString("|")
			}

		}
		buffer.WriteString("\n")
	}

	return buffer.String()
}

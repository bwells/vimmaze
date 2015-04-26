package maze

import (
	"testing"
)

func TestNewMaze(t *testing.T) {
	m := NewMaze(5, 10)
	if m.width != 5 {
		t.Errorf("Width expected 5, got %d", m.width)
	}
	if m.height != 10 {
		t.Errorf("Height expected 10, got %d", m.width)
	}
	if len(m.cells) != 50 {
		t.Errorf("Content slice size expected 50, got %d", len(m.cells))
	}
}

// func TestGetAt(t *testing.T) {
// 	m := NewMaze(5, 3)
// 	fillMaze(m)
//
// 	expected := 0
// 	msg := "Expected %d at (%d, %d), got %d"
// 	for y := 0; y < m.height; y++ {
// 		for x := 0; x < m.width; x++ {
// 			got := m.getAt(x, y)
// 			if got != expected {
// 				t.Errorf(msg, expected, x, y, got)
// 			}
// 			expected += 1
// 		}
// 	}
// }

// func TestSetAt(t *testing.T) {
// 	m := NewMaze(5, 3)
//
// 	expected := 0
// 	msg := "Expected %d at (%d, %d), got %d"
// 	for y := 0; y < m.height; y++ {
// 		for x := 0; x < m.width; x++ {
// 			m.setAt(x, y, expected)
// 			got := m.content[expected]
// 			if got != expected {
// 				t.Errorf(msg, expected, x, y, got)
// 			}
// 			expected += 1
// 		}
// 	}
// }

func TestGetXY(t *testing.T) {
	m := NewMaze(5, 3)

	x, y := m.getXY(7)
	if x != 2 || y != 1 {
		t.Errorf("Expected (%d, %d) at %s, got (%d, %d)", 2, 1, 7, x, y)
	}
}

func TestGetIndex(t *testing.T) {
	m := NewMaze(5, 10)
	got := m.getIndex(4, 8)
	if got != 44 {
		t.Errorf("Expected 44, got %d", got)
	}
}

// func fillMaze(m Maze) {
// 	for i := 0; i < len(m.cells); i++ {
// 		m.content[i] = i
// 	}
// }

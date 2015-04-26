package main

import (
	"fmt"

	"github.com/bwells/vimmaze/maze"
)

func main() {
	m := maze.NewMaze(80, 40)
	fmt.Println(m)
}

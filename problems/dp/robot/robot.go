package robot

import "fmt"

// Given an M x N matrix and an end position, what's the minimum amount of turns a robot can make startine from (0,0) or top left
// The matrix is of integer type, 0 represents valid element, -1 represents a wall
func Robot(m [][]int, end []int) int {
	move(m, 0, 0, 0, Up)
	fmt.Println(m)
	return m[end[0]][end[1]]
}

const (
	Up = iota
	Right
	Down
	Left
)

var dirs = []int{Up, Right, Down, Left}

// dir 0 top, 1 right, down 2, left 3
func move(m [][]int, i, j, val, dir int) {
	m[i][j] = val

	for _, nextDir := range dirs {
		toAdd := 0
		if dir != nextDir {
			toAdd = 1
		}

		switch nextDir {
		case Up:
			if i-1 >= 0 && (m[i-1][j] == 0 || val+toAdd < m[i-1][j]) {
				move(m, i-1, j, val+toAdd, Up)
			}
		case Down:
			if i+1 < len(m) && (m[i+1][j] == 0 || val+toAdd < m[i+1][j]) {
				move(m, i+1, j, val+toAdd, Down)
			}
		case Left:
			if j-1 >= 0 && (m[i][j-1] == 0 || val+toAdd < m[i][j-1]) {
				move(m, i, j-1, val+toAdd, Left)
			}
		case Right:
			if j+1 < len(m[0]) && (m[i][j+1] == 0 || val+toAdd < m[i][j+1]) {
				move(m, i, j+1, val+toAdd, Right)
			}
		}
	}

}

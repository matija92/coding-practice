package robot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobot(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, -1, -1, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, -1, -1, -1, 0},
		{0, 0, 0, 0, 0, 0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	end := []int{6, 8}

	minTurns := Robot(matrix, end)
	assert.Equal(t, 3, minTurns)

	// Add obstacle to optimal path
	matrix = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, -1, -1, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, -1, -1, -1, -1},
		{0, 0, 0, 0, 0, 0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	minTurns = Robot(matrix, end)
	assert.Equal(t, 4, minTurns)
}

package containersort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCanBeSorted_OneSwap is the first test case from problem description
func TestCanBeSorted_OneSwap(t *testing.T) {
	containers := [][]int{
		{1, 2},
		{2, 1},
	}

	assert.True(t, CanBeSorted(containers))
}

// TestCanBeSorted_OneSwap is the second test case from problem description
func TestCanBeSorted_Unsortable(t *testing.T) {
	containers := [][]int{
		{10, 20, 30},
		{1, 1, 1},
		{0, 0, 1},
	}
	assert.False(t, CanBeSorted(containers))
}

func TestCanBeSorted_AlreadySorted(t *testing.T) {
	containers := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	assert.True(t, CanBeSorted(containers))
}

func TestCanBeSorted_OneContainer(t *testing.T) {
	containers := [][]int{
		{1},
	}
	assert.True(t, CanBeSorted(containers))
}

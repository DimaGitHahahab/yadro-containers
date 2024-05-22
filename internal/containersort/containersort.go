package containersort

import (
	"slices"
	"sort"
)

func CanBeSorted(containers [][]int) bool {
	ballsPerContainer := make([]int, len(containers))
	colorFrequency := make([]int, len(containers))

	for i := range containers {
		var containerCapacity int
		for color, amount := range containers[i] {
			colorFrequency[color] += amount
			containerCapacity += amount
		}
		ballsPerContainer[i] += containerCapacity
	}

	sort.Ints(ballsPerContainer)
	sort.Ints(colorFrequency)

	return slices.Equal(ballsPerContainer, colorFrequency)
}

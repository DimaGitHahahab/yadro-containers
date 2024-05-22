package containersort

import (
	"slices"
	"sort"
)

// CanBeSorted checks if the balls in the containers can be sorted the way each container contains balls of only one color
func CanBeSorted(containers [][]int) bool {
	ballsPerContainer := make([]int, len(containers))
	colorFrequency := make([]int, len(containers))

	// Calculate the total number of balls in each container and the overall frequency of each color
	for i := range containers {
		var containerCapacity int
		for color, amount := range containers[i] {
			colorFrequency[color] += amount
			containerCapacity += amount
		}
		ballsPerContainer[i] += containerCapacity
	}

	// Sort to attempt mapping each container's capacity to the number of identical balls
	sort.Ints(ballsPerContainer)
	sort.Ints(colorFrequency)

	return slices.Equal(ballsPerContainer, colorFrequency)
}

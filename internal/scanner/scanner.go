package scanner

import (
	"fmt"
	"os"
)

const (
	maxContainersNum          = 100
	maxSameColorsPerContainer = 1000000000
)

func ScanContainersNum(src *os.File) (int, error) {
	var num int
	_, err := fmt.Fscan(src, &num)
	if err != nil {
		return 0, err
	}
	if num < 1 || num > maxContainersNum {
		return 0, fmt.Errorf("invalid number of containers: %d, expected between 1 and 100", num)
	}

	return num, nil
}

func ScanContainers(src *os.File, containersNum int) ([][]int, error) {
	containers := make([][]int, containersNum)
	for i := 0; i < len(containers); i++ {
		containers[i] = make([]int, containersNum)
		for j := 0; j < len(containers[i]); j++ {
			if _, err := fmt.Fscan(src, &containers[i][j]); err != nil {
				return nil, err
			}
			if containers[i][j] < 0 || containers[i][j] > maxSameColorsPerContainer {
				return nil, fmt.Errorf(
					"invalid number of balls: %d, expected between 0 and 1000000000",
					containers[i][j],
				)
			}
		}
	}

	return containers, nil
}

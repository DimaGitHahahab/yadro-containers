package main

import (
	"fmt"
	"log"
	"os"

	"yadro-containers/internal/containersort"
	"yadro-containers/internal/scanner"
)

var inputSource = os.Stdin

func main() {
	n, err := scanner.ScanContainersNum(inputSource)
	if err != nil {
		log.Fatal("Failed to scan n: ", err)
	}

	containers, err := scanner.ScanContainers(inputSource, n)
	if err != nil {
		log.Fatal("Failed to scan containers: ", err)
	}

	if containersort.CanBeSorted(containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

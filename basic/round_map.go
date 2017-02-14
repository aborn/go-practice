package basic

import (
	"fmt"
)

var roundMap map[int]int

const RoundMapMaxSize int = 2 * 1000

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func Get(key int) int {
	if _, isPresent := roundMap[key]; isPresent {
		return roundMap[key]
	}
	return -1
}

func Put(key int, value int) {
	currentIndex := key % RoundMapMaxSize
	roundMap[currentIndex] = value
	for i := 0; i < RoundMapMaxSize; i++ {
		if abs(i-currentIndex) >= RoundMapMaxSize/2 {
			delete(roundMap, i)
		}
	}
}

func init() {
	roundMap = make(map[int]int, RoundMapMaxSize)
	fmt.Println("init map size = ", len(roundMap))
}

package functions

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	maxspeed := 0
	for _, pile := range piles {
		if pile > maxspeed {
			maxspeed = pile
		}
	}

	minspeed := 0
	speed := maxspeed
	for minspeed < maxspeed {
		speed = (minspeed + maxspeed) / 2
		time := 0
		for _, items := range piles {
			time += int(math.Ceil(float64(items) / float64(speed)))
		}
		if time > h {
			minspeed = speed
		} else {
			maxspeed = speed
		}
		if maxspeed-minspeed == 1 {
			speed = maxspeed
			break
		}
	}
	return speed
}

func timetaken(piles []int, speed int) int {
	time := 0
	for _, items := range piles {
		time += int(math.Ceil(float64(items) / float64(speed)))
	}
	return time
}

func TestminEatingSpeed() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}

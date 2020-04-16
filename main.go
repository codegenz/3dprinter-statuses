package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now())
	var requiredStatuses int
	var printersCount int = 1
	//var result int = 1

	fmt.Scan(&requiredStatuses)

	lastifDouble := requiredStatuses
	if requiredStatuses < 4 {
		fmt.Println(requiredStatuses)
		return
	}
	if requiredStatuses == 4 {
		fmt.Println(requiredStatuses - 1)
		return
	}
	ifdouble := ifDouble(lastifDouble, printersCount, requiredStatuses, 1) + 1
	fmt.Println(ifdouble)

}

func ifDouble(lastDouble, printersCount, requiredStatuses, days int) int {

	var r int

	printersCount = printersCount * 2
	res := requiredStatuses/printersCount + days
	if requiredStatuses == printersCount {
		return res - 1
	}
	if res < lastDouble {
		days++
		r = ifDouble(res, printersCount, requiredStatuses, days)
	} else {
		r = res
		return r
	}
	return r

}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
}

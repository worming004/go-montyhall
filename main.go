package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ch := make(chan bool)
	var result []bool

	run1000Times := func(f func(chan bool), subChan chan bool) {
		for i := 0; i < 1000; i++ {
			f(subChan)
		}
		close(subChan)
	}

	go run1000Times(runMontyHall, ch)

	for elem := range ch {
		result = append(result, elem)
	}

	countTrueF := func(a []bool) int {
		result := 0
		for _, i := range a {
			if i == true {
				result++
			}
		}
		return result
	}

	countFalseF := func(a []bool) int {
		result := 0
		for _, i := range a {
			if i == false {
				result++
			}
		}
		return result
	}

	countTrue := countTrueF(result)
	countFalse := countFalseF(result)

	fmt.Println("nbr of win: ", countTrue, "nbr of fail:", countFalse)
}

func runMontyHall(ch chan bool) {

	monty := initMontyHall()

	selectedDoorIndex := monty.doors.selectRandom()
	monty.playerSelectDoor(selectedDoorIndex)

	monty.openFirstDoor()

	unselectedDoor := monty.doors.applyFilterfilter(monty.filterUnselectedDoor())
	unselectedClosedDoor := unselectedDoor.applyFilterfilter(isClose)
	selectedDoor := unselectedClosedDoor.selectRandom()
	monty.playerChooseSecondDoor(selectedDoor)
	monty.openAllDoor()

	// monty.writeResult()

	ch <- monty.isWin.value
}

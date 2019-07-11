package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type state int

const (
	start state = iota
	playerChooseFirstDoor
	openFirstDoor
	playerChooseSecondDoor
	openAll
)

type maybeBool struct {
	value     bool
	isDefined bool
}

func defaultMaybeBool() *maybeBool {
	return &maybeBool{false, false}
}

type montyHall struct {
	doors                 doors
	state                 state
	selectedDoor          *door
	isWin                 maybeBool
	firstSelectedDoor     *door
	secondSelectedDoor    *door
	doorOpenedByPresenter *door
}

func initMontyHall() montyHall {
	winningDoorPosition := rand.Intn(3)
	var doors []*door
	for i := 0; i < 3; i++ {
		if i == winningDoorPosition {
			doors = append(doors, &door{i, car, false})
		} else {
			doors = append(doors, &door{i, goat, false})
		}
	}
	return montyHall{doors, start, nil, *defaultMaybeBool(), nil, nil, nil}
}

func (m *montyHall) playerSelectDoor(door *door) error {
	if m.state != start {
		return errors.New("invalid operation")
	}

	m.selectedDoor = door
	m.state = playerChooseFirstDoor
	m.firstSelectedDoor = door

	return nil
}

func (m *montyHall) openFirstDoor() error {
	if m.state != playerChooseFirstDoor {
		return errors.New("invalid operation")
	}
	unselectedAndGoatDoorsF := func(ds []*door) doors {
		var result doors
		for _, door := range ds {
			if door.content != car && door != m.selectedDoor {
				result = append(result, door)
			}
		}

		return result
	}

	var unselectedAndGoatDoors = unselectedAndGoatDoorsF(m.doors)
	doorSelectedByPresender := unselectedAndGoatDoors.selectRandom()
	doorSelectedByPresender.isOpen = true

	m.doorOpenedByPresenter = doorSelectedByPresender
	m.state = openFirstDoor
	return nil
}

func (m *montyHall) playerChooseSecondDoor(d *door) error {
	m.selectedDoor = d

	m.secondSelectedDoor = d
	m.state = playerChooseSecondDoor
	return nil
}

func (m *montyHall) openAllDoor() (bool, error) {
	if m.state != playerChooseSecondDoor {
		return false, errors.New("invalid operation")
	}

	if m.selectedDoor.content == car {
		m.isWin = maybeBool{true, true}
	} else {
		m.isWin = maybeBool{false, true}
	}

	m.state = openAll
	return m.isWin.value, nil
}

func (m montyHall) isOver() bool {
	return m.state == openAll
}

func (m *montyHall) filterUnselectedDoor() func(door *door) bool {
	return func(door *door) bool {
		return door != m.selectedDoor
	}
}

func (m montyHall) writeResult() {
	fmt.Printf("win: %v", m.isWin.value)
}

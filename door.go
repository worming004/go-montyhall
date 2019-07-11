package main

import (
	"math/rand"
)

type content string

type door struct {
	id      int
	content content
	isOpen  bool
}

const (
	car  content = "car"
	goat content = "goat"
)

type doors []*door

func (ds doors) applyFilterfilter(filter func(door *door) bool) doors {
	var result doors
	for _, door := range ds {
		if filter(door) {
			result = append(result, door)
		}
	}

	return result
}

func (door door) isClose() bool {
	return door.isOpen == false
}
func isClose(door *door) bool {
	return door.isClose()
}

func (ds *doors) selectRandom() *door {
	dss := (*ds)
	length := len(dss)
	selected := rand.Intn(length)
	selectedDoor := dss[selected]
	return selectedDoor
}

func (ds *doors) first() *door {
	dss := (*ds)
	return dss[0]
}

func (door *door) filterIsClose() func() bool {
	return func() bool {
		return true
	}
}

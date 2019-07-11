package main

import (
	"testing"
)

func TestMontyHallShouldHave3Doors(t *testing.T) {
	m := initMontyHall()
	if len(m.doors) != 3 {
		t.Errorf("monty hall should have 3 doors. Result: %d", len(m.doors))
	}
}

func TestMontyHallShouldHave1DoorWithCar(t *testing.T) {
	m := initMontyHall()
	var winningDoors []*door
	for _, door := range m.doors {
		if door.content == car {
			winningDoors = append(winningDoors, door)
		}
	}
	if len(winningDoors) != 1 {
		t.Errorf("month hall should have 1 winning door. Result: %d", len(winningDoors))
	}
}

func TestMontyHallShouldHave2DoorWithGoat(t *testing.T) {
	montyHall := initMontyHall()
	var loosingDoors []*door
	for _, door := range montyHall.doors {
		if door.content == goat {
			loosingDoors = append(loosingDoors, door)
		}
	}
	if len(loosingDoors) != 2 {
		t.Errorf("month hall should have 2 loosing doors. Result: %d", len(loosingDoors))
	}
}

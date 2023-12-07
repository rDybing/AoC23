package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	d := dataT{
		time: []int{
			7,
			15,
			30,
		},
		distance: []int{
			9,
			40,
			200,
		},
	}
	tt := struct {
		name string
		data dataT
		sumA int
	}{
		name: "Test1",
		data: d,
		sumA: 288,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("%s\n", tt.name)
		var r raceT
		r.part1(d)
		if tt.sumA != r.sum {
			t.Fatalf("\n%s Expected: %d - got: %d\n", tt.name, tt.sumA, r.sum)
		}
	})
}

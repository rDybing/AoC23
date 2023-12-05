package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	d := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	tt := struct {
		name string
		data []string
		sumA int
	}{
		name: "Test",
		data: d,
		sumA: 4361,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("%s\n", tt.name)
		var res resultT
		res.getLimits(d)
		res.part1(tt.data)
		if tt.sumA != res.sumP1 {
			t.Fatalf("\n%s Expected: %d - got: %d\n", tt.name, tt.sumA, res.sumP1)
		}
	})
}

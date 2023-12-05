package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	d := []string{
		"",
	}
	tt := struct {
		name string
		data []string
		sumQ int
		sumA int
	}{
		name: "Test",
		data: d,
		sumA: 4361,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("%s\n", tt.name)
		for i, v := range tt.data {
			// test stuff
		}
		if tt.sumA != tt.sumQ {
			t.Fatalf("\n%s Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

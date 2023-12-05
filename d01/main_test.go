package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	d := dataT{
		cv: []string{
			"1abc2",
			"pqr3stu8vwx",
			"a1b2c3d4e5f",
			"treb7uchet",
		},
	}
	tt := struct {
		name string
		sumQ int
		sumA int
	}{
		name: "Test1",
		sumA: 142,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("Test 1\n")
		tt.sumQ, _ = d.sumIt()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\nSum1 %s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

func TestPart2(t *testing.T) {
	d := dataT{
		cv: []string{
			"two1nine",
			"eightwothree",
			"abcone2threexyz",
			"xtwone3four",
			"4nineeightseven2",
			"zoneight234",
			"7pqrstsixteen",
			"one2one",
		},
	}
	tt := struct {
		name string
		sumQ int
		sumA int
	}{
		name: "Test2",
		sumA: 292,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("Test 2\n")
		_, tt.sumQ = d.sumIt()
		if tt.sumA != tt.sumQ {
			t.Fatalf("\nSum2 %s: Expected: %d - got: %d\n", tt.name, tt.sumA, tt.sumQ)
		}
	})
}

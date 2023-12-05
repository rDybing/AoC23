package main

import (
	"fmt"
	"testing"
)

func TestParseData(t *testing.T) {
	d := dataT{
		game: []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
		},
	}
	tt := struct {
		name string
		sumQ []rgbT
		sumA []rgbT
	}{
		name: "ParseData",
		sumA: fillAnswerParse(),
	}
	var game gameT
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("Test %v\n", tt.name)
		d.setMax()
		game.parseGames(d)
		tt.sumQ = game.cubes
		for i, v := range tt.sumQ {
			if v != tt.sumA[i] {
				t.Fatalf("\n%s round %d: Expected: %+v - got: %+v\n", tt.name, i, tt.sumA[i], v)
			}
		}
	})
}

func fillAnswerParse() []rgbT {
	var out []rgbT
	// round 1
	temp := rgbT{
		red:   4,
		green: 2,
		blue:  6,
	}
	out = append(out, temp)
	// round 2
	temp = rgbT{
		red:   1,
		green: 3,
		blue:  4,
	}
	out = append(out, temp)
	// round 3
	temp = rgbT{
		red:   20,
		green: 13,
		blue:  6,
	}
	out = append(out, temp)
	// round 4
	temp = rgbT{
		red:   14,
		green: 3,
		blue:  15,
	}
	out = append(out, temp)
	// round 5
	temp = rgbT{
		red:   6,
		green: 3,
		blue:  2,
	}
	out = append(out, temp)
	return out
}

func TestPart1(t *testing.T) {
	d := dataT{
		game: []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			"",
		},
	}
	tt := struct {
		name string
		sumQ int
		sumA int
	}{
		name: "Test1",
		sumA: 8,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("Test 1\n")
		d.setMax()
		game.parseGames(d)
		game.part1(d.maxRed, d.maxGreen, d.maxBlue)
		if tt.sumA != game.sum1 {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, game.sum1)
		}
	})
}

func TestPart2(t *testing.T) {
	d := dataT{
		game: []string{
			"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			"",
		},
	}
	tt := struct {
		name string
		sumQ int
		sumA int
	}{
		name: "Test2",
		sumA: 2286,
	}
	t.Run(tt.name, func(t *testing.T) {
		fmt.Printf("Test 1\n")
		d.setMax()
		game.parseGames(d)
		game.part2()
		if tt.sumA != game.sum2 {
			t.Fatalf("\n%s: Expected: %d - got: %d\n", tt.name, tt.sumA, game.sum2)
		}
	})
}

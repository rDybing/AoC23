package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	game     []string
	maxRed   int
	maxGreen int
	maxBlue  int
}

type gameT struct {
	cubes []rgbT
	sum1  int
	sum2  int
}

type rgbT struct {
	red   int
	green int
	blue  int
}

var game gameT

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	fmt.Println("parsing data")
	game.parseGames(d)
	game.part1(d.maxRed, d.maxGreen, d.maxBlue)
	game.part2()
	fmt.Printf("Part 1: The sum of possible games are %d\n", game.sum1)
	fmt.Printf("Part 2: The sum of the powers of cubes in all games are %d\n", game.sum2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (g *gameT) part1(mRed, mGreen, mBlue int) {
	for i, v := range game.cubes {
		if v.red <= mRed && v.green <= mGreen && v.blue <= mBlue {
			sum := i + 1
			game.sum1 += sum
		} else {
		}
	}
}

func (g *gameT) part2() {
	for _, v := range g.cubes {
		power := v.red * v.green * v.blue
		g.sum2 += power
	}
}

func importData() (dataT, error) {
	var d dataT
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	d.game = append(d.game, arr...)
	fmt.Printf("loaded %d data-points\n", len(d.game)-1)
	d.setMax()
	return d, nil
}

func (d *dataT) setMax() {
	d.maxRed = 12
	d.maxGreen = 13
	d.maxBlue = 14
}

func (g *gameT) parseGames(d dataT) {
	for _, v := range d.game {
		if v != "" {
			// get rid of "Game X: " text
			prepRound := strings.Split(v, ": ")[1]
			// get array of bunches of cubes in a round
			bunchStr := strings.Split(prepRound, "; ")
			var roundRGB []rgbT
			// for each bunch, get cubes
			for _, bunch := range bunchStr {
				cubeStr := strings.Split(bunch, ", ")
				var tRGB rgbT
				// get individual RGB cubes
				for _, cubes := range cubeStr {
					cube := strings.Split(cubes, " ")
					switch cube[1] {
					case "red":
						tRGB.red, _ = strconv.Atoi(cube[0])
					case "green":
						tRGB.green, _ = strconv.Atoi(cube[0])
					case "blue":
						tRGB.blue, _ = strconv.Atoi(cube[0])
					}
				}
				bunchRGB := rgbT{
					red:   tRGB.red,
					green: tRGB.green,
					blue:  tRGB.blue,
				}
				roundRGB = append(roundRGB, bunchRGB)
			}
			// get highest number of each color in a round
			var roundSumRGB rgbT
			for _, round := range roundRGB {
				if round.red > roundSumRGB.red {
					roundSumRGB.red = round.red
				}
				if round.green > roundSumRGB.green {
					roundSumRGB.green = round.green
				}
				if round.blue > roundSumRGB.blue {
					roundSumRGB.blue = round.blue
				}
			}
			g.cubes = append(g.cubes, roundSumRGB)
		}
	}
}

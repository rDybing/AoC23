package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type resultT struct {
	hLimit int
	vLimit int
	sumP1  int
	sumP2  int
}

type posIndexT struct {
	temp  int
	start int
	stop  int
}

type partT struct {
	isPart  bool
	symbols string
	number  int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	var result resultT
	result.getLimits(d)
	result.part1(d)
	fmt.Printf("Part 1: sum of engine parts are %d\n", result.sumP1)
	//fmt.Printf("Part 2: Following new strategy in part 2, score is %d\n", score2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (r *resultT) getLimits(d []string) {
	r.hLimit = len(d[0]) - 1
	r.vLimit = len(d) - 1
}

func (r *resultT) part1(d []string) {
	var pi posIndexT
	var part partT
	var partList []partT
	for vIndex, lineTxt := range d {
		var fullNum []int
		pi.temp = -1
		var checkAdjacent bool
		for hIndex, char := range lineTxt {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				fullNum = append(fullNum, digit)
				if pi.temp == -1 {
					pi.temp = hIndex
				}
			} else {
				if len(fullNum) != 0 {
					part.number = getNumberFromDigits(fullNum)
					pi.stop = hIndex
					pi.start = pi.temp
					// reset
					pi.temp = -1
					fullNum = nil
					checkAdjacent = true
				}
				// check adjacent
				if checkAdjacent {
					switch vIndex {
					// check if first line, if so send current and next line
					case 0:
						part.symbols = pi.checkLineAdjSymbol(r.hLimit, lineTxt, true)
						part.symbols += pi.checkLineAdjSymbol(r.hLimit, d[vIndex+1], false)
					// check if last line, if so send current and previous line
					case r.vLimit:
						part.symbols = pi.checkLineAdjSymbol(r.hLimit, lineTxt, true)
						part.symbols += pi.checkLineAdjSymbol(r.hLimit, d[vIndex-1], false)
					// else send current, previous and next line
					default:
						part.symbols = pi.checkLineAdjSymbol(r.hLimit, lineTxt, true)
						part.symbols += pi.checkLineAdjSymbol(r.hLimit, d[vIndex+1], false)
						part.symbols += pi.checkLineAdjSymbol(r.hLimit, d[vIndex-1], false)
					}
					// set isPart if symbols are filled
					if part.symbols != "" {
						part.isPart = true
					} else {
						part.isPart = false
					}
					// transfer to partslist
					partList = append(partList, part)
					// reset
					checkAdjacent = false
				}
			}
		}
	}
	fmt.Printf("%+v\n", partList)
	r.sumP1 = part1Sum(partList)
}

func part1Sum(p []partT) int {
	var out int
	for _, v := range p {
		if v.isPart {
			out += v.number
		}
	}
	return out
}

func getNumberFromDigits(in []int) int {
	var numStr string
	for _, v := range in {
		numStr += fmt.Sprintf("%d", v)
	}
	out, _ := strconv.Atoi(numStr)
	return out
}

func (p posIndexT) checkLineAdjSymbol(hLimit int, line string, sameLine bool) string {
	var out string
	start := p.start - 1
	stop := p.stop + 1
	// check if start of line
	if p.start == 0 {
		start = 0
	}
	// check if end of line
	if p.stop == len(line)-1 {
		stop = p.stop
	}
	for i := start; i < stop; i++ {
		if rune(line[i]) != '.' && !isDigit(rune(line[i])) {
			out += string(line[i])
		}
	}
	return out
}

func isDigit(r rune) bool {
	var out bool
	if _, err := strconv.Atoi(string(r)); err == nil {
		out = true
	}
	return out
}

func importData() ([]string, error) {
	var d []string
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	d = append(d, arr...)
	fmt.Printf("loaded %d data-points\n", len(d))
	return d, nil
}

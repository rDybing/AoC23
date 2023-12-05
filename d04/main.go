package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type cardT struct {
	winNum     []int
	scratchNum []int
}

type resultT struct {
	sum1     int
	sum2     int
	cardsSum []int
}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	var res resultT
	res.part1(d)
	fmt.Printf("Part 1: sum of winning cards are: %d\n", res.sum1)
	res.part2(d)
	fmt.Printf("Part 2: sum of total cards are %d\n", res.sum2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (r *resultT) part1(d []string) {
	for _, v := range d {
		var card cardT
		var cardWin int
		card.parseData(v)
		for _, cs := range card.scratchNum {
			for _, cw := range card.winNum {
				if cs == cw {
					if cardWin == 0 {
						cardWin = 1
					} else {
						cardWin *= 2
					}
				}
			}
		}
		r.sum1 += cardWin
	}
}

func (r *resultT) part2(d []string) {
	r.prepPart2(len(d))
	for i, v := range d {
		var card cardT
		var sumNewCards int
		card.parseData(v)
		for j := 0; j < r.cardsSum[i]; j++ {
			var newCards int
			for _, cs := range card.scratchNum {
				for _, cw := range card.winNum {
					if cs == cw {
						newCards++
					}
				}
			}
			sumNewCards += newCards
			r.addCards(i+1, newCards)
		}
	}
	r.sumCardsPart2()
}

func (r *resultT) sumCardsPart2() {
	for _, v := range r.cardsSum {
		r.sum2 += v
	}
}

func (r *resultT) prepPart2(arrLength int) {
	for i := 0; i < arrLength; i++ {
		r.cardsSum = append(r.cardsSum, 1)
	}
}

func (r *resultT) addCards(index int, newCards int) {
	for i := index; i < index+newCards; i++ {
		r.cardsSum[i]++
	}
}

func (c *cardT) parseData(data string) {
	card := strings.Split(data, ":")[1]
	temp := strings.Split(card, "|")
	winStr := strings.Split(temp[0], " ")
	for _, v := range winStr {
		num, _ := strconv.Atoi(v)
		if num != 0 {
			c.winNum = append(c.winNum, num)
		}
	}
	scratchStr := strings.Split(temp[1], " ")
	for _, v := range scratchStr {
		num, _ := strconv.Atoi(v)
		if num != 0 {
			c.scratchNum = append(c.scratchNum, num)
		}
	}
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

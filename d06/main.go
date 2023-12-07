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
	time     []int
	distance []int
}

type raceT struct {
	result []int
	sum1   int
	sum2   int
}

var d dataT

func main() {
	err := d.importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	var r raceT
	r.part1(d)
	fmt.Printf("Part 1: sum of race wins are %d\n", r.sum1)
	r.part2(d)
	fmt.Printf("Part 2: sum of long race is %d\n", r.sum2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (r *raceT) part1(d dataT) {
	for i, time := range d.time {
		var round []int
		dist := d.distance[i]
		for ms := 1; ms < time; ms++ {
			timeToMove := time - ms
			distResult := timeToMove * ms
			fmt.Printf("Time: %d - Distance: %d - MS: %d - TTM: %d - DR: %d\n", time, dist, ms, timeToMove, distResult)
			round = append(round, distResult)
		}
		var resTemp int
		for _, result := range round {
			if result > dist {
				resTemp++
			}
		}
		fmt.Printf("temp result round %d: %d winning rounds\n", i, resTemp)
		r.result = append(r.result, resTemp)
	}
	for _, sumTemp := range r.result {
		if r.sum1 == 0 {
			r.sum1 = sumTemp
		} else {
			r.sum1 *= sumTemp
		}
	}
	fmt.Printf("result: %+v\n", r)
}

func (r *raceT) part2(d dataT) {
	time := joinNumbers(d.time)
	dist := joinNumbers(d.distance)
	var round []int
	for ms := 1; ms < time; ms++ {
		timeToMove := time - ms
		distResult := timeToMove * ms
		//fmt.Printf("Time: %d - Distance: %d - MS: %d - TTM: %d - DR: %d\n", time, dist, ms, timeToMove, distResult)
		round = append(round, distResult)
	}
	for _, result := range round {
		if result > dist {
			r.sum2++
		}
	}
	fmt.Printf("result: %+v\n", r)
}

func joinNumbers(in []int) int {
	var outStr string
	for _, v := range in {
		outStr += strconv.Itoa(v)
	}
	out, _ := strconv.Atoi(outStr)
	return out
}

func (d *dataT) importData() error {
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return err
	}
	arr := strings.Split(string(f), "\n")
	timeArr := strings.Fields(arr[0])
	distArr := strings.Fields(arr[1])
	for i, v := range timeArr {
		if i > 0 {
			timeNum, _ := strconv.Atoi(v)
			distNum, _ := strconv.Atoi(distArr[i])
			d.time = append(d.time, timeNum)
			d.distance = append(d.distance, distNum)
		}
	}
	fmt.Printf("data import: %+v\n", d)
	return nil
}

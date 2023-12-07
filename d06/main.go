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
	sum    int
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
	fmt.Printf("Part 1: sum of race wins are %d\n", r.sum)
	//fmt.Printf("Part 2: Following new strategy in part 2, score is %d\n", score2)
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
		if r.sum == 0 {
			r.sum = sumTemp
		} else {
			r.sum *= sumTemp
		}
	}
	fmt.Printf("result: %+v\n", r)
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

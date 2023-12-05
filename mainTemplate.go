package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()

	sum1 := part1(d)
	fmt.Printf("Part 1: sum of engine parts are %d\n", sum1)
	//fmt.Printf("Part 2: Following new strategy in part 2, score is %d\n", score2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func part1(d []string) int {
	for _, v := range d {
		fmt.Println(v)
	}
	return 0
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

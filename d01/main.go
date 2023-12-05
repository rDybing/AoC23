package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type dataT struct {
	cv []string
}

type resultT struct {
	number int
	index  int
	line   int
}

var strNum = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	d, err := importData()
	if err != nil {
		log.Panicf("Failed loading data file: %v\n", err)
	}
	start := time.Now()
	sum1, sum2 := d.sumIt()
	fmt.Printf("Part 1: Calibration sum is %d \n", sum1)
	fmt.Printf("Part 2: Calibration including numbers as strings is %d\n", sum2)
	done := time.Now()
	diff := done.Sub(start)
	fmt.Printf("Execution time: %d µSeconds\n", diff.Microseconds())
}

func (d dataT) sumIt() (int, int) {
	var sum1, sum2 int
	for i, v := range d.cv {
		if v != "" {
			res := getNumbers(i, v)
			if res != nil {
				sum1 += (res[0].number * 10) + res[1].number
			}
			res = getStrNum(i, v, res)
			sum2 += (res[0].number*10 + res[1].number)
		}
	}
	return sum1, sum2
}

func getNumbers(line int, in string) []resultT {
	fmt.Println("getNumbers")
	var res []resultT
	for i, v := range in {
		if x, err := strconv.Atoi(string(v)); err == nil {
			resTemp := resultT{
				number: x,
				index:  i,
				line:   line,
			}
			res = append(res, resTemp)
		}
	}
	if len(res) != 2 && res != nil {
		res = checkLen(res)
	}
	return res
}

func getStrNum(line int, in string, res []resultT) []resultT {
	fmt.Printf("getStrNum: %s\n", in)
	for i, v := range strNum {
		if pos := strings.Index(in, v); pos != -1 {
			tempRes := resultT{
				number: i + 1,
				index:  pos,
				line:   line,
			}
			fmt.Printf("found str %s as num %d at index %d\n", v, tempRes.number, tempRes.index)
			res = append(res, tempRes)
			if lastPos := strings.LastIndex(in, v); lastPos != pos {
				tempRes := resultT{
					number: i + 1,
					index:  lastPos,
					line:   line,
				}
				fmt.Printf("found str %s as num %d at index %d\n", v, tempRes.number, tempRes.index)
				res = append(res, tempRes)
			}

		}
	}
	if len(res) != 2 {
		res = sortOnIndex(res)
		res = checkLen(res)
	}
	return res
}

func sortOnIndex(res []resultT) []resultT {
	sort.Slice(res, func(i, j int) bool {
		return res[i].index < res[j].index
	})
	return res
}

func checkLen(in []resultT) []resultT {
	fmt.Println("checkLen")
	var res []resultT
	fmt.Printf("in : %+v\n", in)
	if len(in) == 1 {
		in = append(in, in[0])
		res = in
	}
	if len(in) > 2 {
		res = append(res, in[0])
		res = append(res, in[len(in)-1])
	}
	fmt.Printf("out: %+v\n", res)
	return res
}

func importData() (dataT, error) {
	var d dataT
	f, err := os.ReadFile("./data.txt")
	if err != nil {
		return d, err
	}
	arr := strings.Split(string(f), "\n")
	d.cv = append(d.cv, arr...)
	fmt.Printf("loaded %d data-points\n", len(d.cv))
	return d, nil
}

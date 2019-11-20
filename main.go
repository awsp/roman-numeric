package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)


func main() {
	var notations = map[int]string{1000: "M", 500: "D", 100: "C", 50: "L", 10: "X", 5: "V", 1: "I",}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	r := Calc(n, GetKeys(notations))
	s := Convert(r, notations)
	fmt.Printf("%d => %s\n", n, s)
}

func GetKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func (i, j int) bool {
		return keys[i] > keys[j]
	})

	return keys
}

func Convert(arr []int, symbols map[int]string) string {
	var result = ""
	for _, v := range arr {
		if s, ok := symbols[v]; ok != false {
			result += s
		}
	}

	return result
}

func NLength(n int) int {
	return len(strconv.Itoa(n))
}

func Calc(n int, arr []int) []int {
	if n <= 0 {
		return nil
	}

	nn := n
	var result []int

	for {
		if nn <= 0 {
			break
		}

		for _, v := range arr {
			l := NLength(nn)
			p := int(math.Pow(10, float64(l-1)))

			if v <= nn {
				result = append(result, v)
				nn = nn - v
				break
			}
			if (nn + p) >= v {
				result = append(result, p)
				result = append(result, v)
				nn = nn - (v - p)
				break
			}
		}
	}

	return result
}

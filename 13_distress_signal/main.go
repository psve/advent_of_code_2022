package main

import (
	"fmt"
	"helper"
	"reflect"
	"sort"
	"strconv"
)

func parseString(in string, index int) ([]any, int) {
	out := make([]any, 0)
	x := ""
	for index < len(in) {
		switch in[index] {
		case '[':
			elem, i := parseString(in, index+1)
			out = append(out, elem)
			index = i
		case ']':
			if x != "" {
				elem, _ := strconv.Atoi(x)
				out = append(out, elem)
			}
			return out, index
		case ',':
			if x != "" {
				elem, _ := strconv.Atoi(x)
				out = append(out, elem)
				x = ""
			}
		default:
			x += string(in[index])
		}
		index++
	}
	return out, index
}

func compare(left, right []any) int {
	for i, l := range left {
		if i == len(right) {
			return -1
		}

		r := right[i]
		lType := reflect.TypeOf(l).Kind()
		rType := reflect.TypeOf(r).Kind()
		switch {
		case lType == reflect.Int && rType == reflect.Int:
			if l.(int) < r.(int) {
				return 1
			}
			if l.(int) > r.(int) {
				return -1
			}
		case lType == reflect.Slice && rType == reflect.Int:
			if res := compare(l.([]any), []any{r}); res != 0 {
				return res
			}
		case lType == reflect.Int && rType == reflect.Slice:
			if res := compare([]any{l}, r.([]any)); res != 0 {
				return res
			}
		case lType == reflect.Slice && rType == reflect.Slice:
			if res := compare(l.([]any), r.([]any)); res != 0 {
				return res
			}
		}
	}

	if len(left) < len(right) {
		return 1
	}

	return 0
}

func checkOrder(path string) int {
	var left, right []any
	pair, sum := 1, 0
	helper.ForEachLine(path, func(line string) error {
		if line == "" {
			return nil
		}
		if left == nil {
			left, _ = parseString(line, 1)
			return nil
		}
		right, _ = parseString(line, 1)

		if compare(left, right) == 1 {
			sum += pair
		}
		pair++
		left, right = nil, nil
		return nil
	})

	return sum
}

type packets []any

// Implement sort.Interface for packets
func (p packets) Len() int {
	return len(p)
}

func (p packets) Less(i, j int) bool {
	if compare(p[i].([]any), p[j].([]any)) == 1 {
		return true
	}
	return false
}

func (p packets) Swap(i, j int) {
	t := p[i]
	p[i] = p[j]
	p[j] = t
}

func orderPackets(path string) int {
	d1, d2 := []any{[]any{2}}, []any{[]any{6}}
	packets := packets{d1, d2}
	helper.ForEachLine(path, func(line string) error {
		if line == "" {
			return nil
		}
		packet, _ := parseString(line, 1)
		packets = append(packets, packet)
		return nil
	})

	sort.Sort(packets)
	res := 1
	for i, p := range packets {
		if reflect.DeepEqual(p, d1) || reflect.DeepEqual(p, d2) {
			res *= i + 1
		}
	}
	return res
}

func main() {
	fmt.Println(checkOrder("./input"))
	fmt.Println("---------")
	fmt.Println(orderPackets("./input"))
}

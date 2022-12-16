package main

import (
	"reflect"
	"testing"
)

func TestParseString(t *testing.T) {
	cases := []struct {
		in  string
		out []any
	}{
		{in: "[1,1,3,1,1]", out: []any{1, 1, 3, 1, 1}},
		{in: "[1,1,5,1,1]", out: []any{1, 1, 5, 1, 1}},
		{in: "[[1],[2,3,4]]", out: []any{[]any{1}, []any{2, 3, 4}}},
		{in: "[[1],4]", out: []any{[]any{1}, 4}},
		{in: "[9]", out: []any{9}},
		{in: "[[8,7,6]]", out: []any{[]any{8, 7, 6}}},
		{in: "[[4,4],4,4]", out: []any{[]any{4, 4}, 4, 4}},
		{in: "[[4,4],4,4,4]", out: []any{[]any{4, 4}, 4, 4, 4}},
		{in: "[7,7,7,7]", out: []any{7, 7, 7, 7}},
		{in: "[7,7,7]", out: []any{7, 7, 7}},
		{in: "[]", out: []any{}},
		{in: "[3]", out: []any{3}},
		{in: "[[[]]]", out: []any{[]any{[]any{}}}},
		{in: "[[]]", out: []any{[]any{}}},
		{in: "[1,[2,[3,[4,[5,6,7]]]],8,9]", out: []any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9}},
		{in: "[1,[2,[3,[4,[5,6,0]]]],8,9]", out: []any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9}},
	}

	for _, c := range cases {
		res, _ := parseString(c.in, 1)
		if !reflect.DeepEqual(res, c.out) {
			t.Errorf("wrong result: %+v != %+v", res, c.out)
		}
	}
}

func TestCompare(t *testing.T) {
	cases := []struct {
		out         int
		left, right []any
	}{
		{out: 1, left: []any{1, 1, 3, 1, 1}, right: []any{1, 1, 5, 1, 1}},
		{out: 1, left: []any{[]any{1}, []any{2, 3, 4}}, right: []any{[]any{1}, 4}},
		{out: -1, left: []any{9}, right: []any{[]any{8, 7, 6}}},
		{out: 1, left: []any{[]any{4, 4}, 4, 4}, right: []any{[]any{4, 4}, 4, 4, 4}},
		{out: -1, left: []any{7, 7, 7, 7}, right: []any{7, 7, 7}},
		{out: 1, left: []any{}, right: []any{3}},
		{out: -1, left: []any{[]any{[]any{}}}, right: []any{[]any{}}},
		{out: -1, left: []any{1, []any{2, []any{3, []any{4, []any{5, 6, 7}}}}, 8, 9}, right: []any{1, []any{2, []any{3, []any{4, []any{5, 6, 0}}}}, 8, 9}},
	}

	for _, c := range cases {
		res := compare(c.left, c.right)
		if res != c.out {
			t.Errorf("wrong result %d", res)
		}
	}
}

func TestCheckOrder(t *testing.T) {
	res := checkOrder("./input_test")
	if res != 13 {
		t.Errorf("wrong result %d", res)
	}
}

func TestOrderPackets(t *testing.T) {
	res := orderPackets("./input_test")
	if res != 140 {
		t.Errorf("wrong result %d", res)
	}
}

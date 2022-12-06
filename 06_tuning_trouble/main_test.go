package main

import "testing"

func TestFindMarker(t *testing.T) {
	cases := []struct {
		input  string
		result int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, c := range cases {
		if res := findMarker([]rune(c.input)); c.result != res {
			t.Errorf("wrong result (%d != %d)", c.result, res)
		}
	}
}

func TestFindMessage(t *testing.T) {
	cases := []struct {
		input  string
		result int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, c := range cases {
		if res := findMessage([]rune(c.input)); c.result != res {
			t.Errorf("wrong result (%d != %d)", c.result, res)
		}
	}
}

package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func manhattanDistance(x1, y1, x2, y2 int64) int64 {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	return (x2 - x1) + (y2 - y1)
}

type sensor struct {
	x, y             int64
	beaconX, beaconY int64
	distance         int64
}

func parse(path string) []sensor {
	sensors := make([]sensor, 0)
	helper.ForEachLine(path, func(line string) error {
		s, b, _ := strings.Cut(line, ":")
		sx, sy, _ := strings.Cut(strings.TrimPrefix(s, "Sensor at x="), ",")
		sy = strings.TrimPrefix(sy, " y=")
		bx, by, _ := strings.Cut(strings.TrimPrefix(b, " closest beacon is at x="), ",")
		by = strings.TrimPrefix(by, " y=")

		sensorX, _ := strconv.ParseInt(sx, 10, 64)
		sensorY, _ := strconv.ParseInt(sy, 10, 64)
		beaconX, _ := strconv.ParseInt(bx, 10, 64)
		beaconY, _ := strconv.ParseInt(by, 10, 64)
		distance := manhattanDistance(sensorX, sensorY, beaconX, beaconY)

		sensors = append(sensors, sensor{
			x:        sensorX,
			y:        sensorY,
			beaconX:  beaconX,
			beaconY:  beaconY,
			distance: distance,
		})
		return nil
	})
	return sensors
}

func checkRow(sensors []sensor, row int64) map[int64]any {
	rowExclusion := make(map[int64]any)

	for _, s := range sensors {
		// Calculate the width of the area the sensor covers in the target row. This is equal to twice
		// the distance from the row to the furthest point64 the sensor covers.
		rowDist := manhattanDistance(0, s.y, 0, row)
		halfZoneWidth := s.distance - rowDist
		for i := s.x - halfZoneWidth; i <= s.x+halfZoneWidth; i++ {
			rowExclusion[i] = true
		}

		// Remove any beacon that happens to be in the row
		if s.beaconY == row {
			delete(rowExclusion, s.beaconX)
		}
	}

	return rowExclusion
}

func findBeacon(sensors []sensor, minX, maxX int64) int64 {
	for y := int64(0); y <= maxX; y++ {
		for x := int64(0); x <= maxX; x++ {
			covered := false
			for _, s := range sensors {
				if manhattanDistance(x, y, s.x, s.y) <= s.distance {
					covered = true
					// We are inside the range of a sensor. We can skip ahead to the earliest point that is
					// outside its range.
					rowDist := manhattanDistance(0, s.y, 0, y)
					halfZoneWidth := s.distance - rowDist
					x = min(s.x+halfZoneWidth, maxX)
					break
				}
			}
			if !covered {
				return x*4000000 + y
			}
		}
	}
	return 0
}

func main() {
	sensors := parse("./input")
	fmt.Println(len(checkRow(sensors, 2000000)))
	fmt.Println("---------")
	fmt.Println(findBeacon(sensors, 0, 4000000))
}

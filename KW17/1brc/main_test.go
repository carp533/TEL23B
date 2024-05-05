package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ExampleCreateDataFile() {
	//Output:
	//
	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	out, err := os.OpenFile("measurements1m.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	for i := 0; i < 1_000_000; i++ {
		j := i % len(lines)
		fmt.Fprint(out, lines[j], "\n")
	}
}
func ExampleBaseline() {
	//Output:
	//
	inputPath := "measurements1m.txt"
	type stats struct {
		min, max, sum float64
		count         int64
	}

	f, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stationStats := make(map[string]stats)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		station, tempStr, hasSemi := strings.Cut(line, ";")
		if !hasSemi {
			continue
		}

		temp, err := strconv.ParseFloat(tempStr, 64)
		if err != nil {
			log.Fatal(err)
		}

		s, ok := stationStats[station]
		if !ok {
			s.min = temp
			s.max = temp
			s.sum = temp
			s.count = 1
		} else {
			s.min = min(s.min, temp)
			s.max = max(s.max, temp)
			s.sum += temp
			s.count++
		}
		stationStats[station] = s
	}

	stations := make([]string, 0, len(stationStats))
	for station := range stationStats {
		stations = append(stations, station)
	}
	sort.Strings(stations)

	for _, station := range stations {
		s := stationStats[station]
		mean := s.sum / float64(s.count)
		fmt.Printf("%s=%.1f/%.1f/%.1f\n", station, s.min, mean, s.max)
	}
}
func min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}
func max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

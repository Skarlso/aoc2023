package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type mapRange struct {
	srcFrom, srcTo int
	dstFrom, dstTo int
}

type mapMap struct {
	ranges []mapRange
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	// maps
	// 50 98 2
	// 52 50 48
	// dest - source -> 2 ( this is the offset )
	// range is 48 os dest - source + range?
	// Seed number 79 corresponds to soil number 81.
	// Which is 79 + 2 -> because it's in the range of the second line in which the dest - source offset is 2!
	// Nice

	var (
		seedToSoil   mapMap
		soilToFert   mapMap
		fertToWater  mapMap
		waterToLight mapMap
		lightToTemp  mapMap
		tempToHum    mapMap
		humToLoc     mapMap
	)

	seeds := make([]int, 0)
	split := strings.Split(string(content), "\n")

	// parsing
	var (
		soil  bool
		fert  bool
		water bool
		light bool
		temp  bool
		hum   bool
		loc   bool
	)
	for _, l := range split {
		if l == "" {
			continue
		}

		if strings.Contains(l, "seeds") {
			sds := strings.Split(l, ": ")
			numbers := sds[1]
			nums := strings.Split(numbers, " ")
			for _, n := range nums {
				i, _ := strconv.Atoi(n)
				seeds = append(seeds, i)
			}

			continue
		}

		// this is ugly as fuck but who cares.
		if strings.Contains(l, "seed-to-soil map:") {
			soil = true
			continue
		}
		if strings.Contains(l, "soil-to-fertilizer map:") {
			soil = false
			fert = true
			continue
		}
		if strings.Contains(l, "fertilizer-to-water map:") {
			fert = false
			water = true
			continue
		}
		if strings.Contains(l, "water-to-light map:") {
			water = false
			light = true
			continue
		}
		if strings.Contains(l, "light-to-temperature map:") {
			light = false
			temp = true
			continue
		}
		if strings.Contains(l, "temperature-to-humidity map:") {
			temp = false
			hum = true
			continue
		}
		if strings.Contains(l, "humidity-to-location map:") {
			hum = false
			loc = true
			continue
		}

		switch {
		case soil:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			seedToSoil.ranges = append(seedToSoil.ranges, mm)
		case fert:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			soilToFert.ranges = append(soilToFert.ranges, mm)
		case water:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			fertToWater.ranges = append(fertToWater.ranges, mm)
		case light:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			waterToLight.ranges = append(waterToLight.ranges, mm)
		case temp:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			lightToTemp.ranges = append(lightToTemp.ranges, mm)
		case hum:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			tempToHum.ranges = append(tempToHum.ranges, mm)
		case loc:
			var (
				fromSource, fromDestination, rng int
			)
			fmt.Sscanf(l, "%d %d %d", &fromDestination, &fromSource, &rng)
			mm := mapRange{}
			mm.dstFrom = fromDestination
			mm.dstTo = fromDestination + rng - 1 // including the starting point hence -1
			mm.srcFrom = fromSource
			mm.srcTo = fromSource + rng - 1 // including the starting point hence -1

			humToLoc.ranges = append(humToLoc.ranges, mm)
		}
	}
	// parse through the seeds and start mapping them.

	// TODO: Are there multiple matches?
	var locations []int
	for _, seed := range seeds {
		fmt.Println("Seed: ", seed)
		soilSrc := seed

		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range seedToSoil.ranges {
			if seed >= ss.srcFrom && seed <= ss.srcTo {
				// fmt.Println(ss)
				// fmt.Println("abs: ", abs(ss.dstFrom-ss.srcFrom))
				soilSrc = seed + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("Soil: ", soilSrc)

		fertSrc := soilSrc
		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range soilToFert.ranges {
			if fertSrc >= ss.srcFrom && fertSrc <= ss.srcTo {
				fertSrc = soilSrc + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("Fert: ", fertSrc)

		waterSrc := fertSrc
		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range fertToWater.ranges {
			if waterSrc >= ss.srcFrom && waterSrc <= ss.srcTo {
				waterSrc = fertSrc + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("Water: ", waterSrc)

		lightSrc := waterSrc
		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range waterToLight.ranges {
			if waterSrc >= ss.srcFrom && waterSrc <= ss.srcTo {
				lightSrc = waterSrc + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("Light: ", lightSrc)

		tempSrc := lightSrc
		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range lightToTemp.ranges {
			if lightSrc >= ss.srcFrom && lightSrc <= ss.srcTo {
				tempSrc = lightSrc + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("temp: ", tempSrc)

		humSrc := tempSrc
		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range tempToHum.ranges {
			if tempSrc >= ss.srcFrom && tempSrc <= ss.srcTo {
				humSrc = tempSrc + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("Hum: ", humSrc)

		loc := humSrc
		// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
		for _, ss := range humToLoc.ranges {
			if humSrc >= ss.srcFrom && humSrc <= ss.srcTo {
				loc = humSrc + (ss.dstFrom - ss.srcFrom) // seed + offset
				// assume there is only one match for now I guess... and hopefully we don't have to fucking try all of them.
				break
			}
		}

		fmt.Println("Loc: ", loc)

		locations = append(locations, loc)
	}

	sort.Ints(locations)

	fmt.Println("locations: ", locations)
}

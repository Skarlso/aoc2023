package main

import (
	"fmt"
	"os"
	"strings"
)

type signal struct {
	p    pulse
	from string
	to   string
}

type pulse string

var low pulse = "low"
var high pulse = "high"

// this could be an embedded field, but I don't care right now
// this guy now needs syncing.... does it though...?
var modules = map[string]*module{}

type orchestrator struct {
	queue []signal
}

func (o *orchestrator) Process() {
	var s signal
	for len(o.queue) > 0 {
		s, o.queue = o.queue[0], o.queue[1:]

		if s.p == high {
			highCount++
		} else {
			lowCount++
		}

		if s.to == "tr" && s.p == low {
			fmt.Println("tr got a low call")
			fmt.Println(lowCount, highCount, lowCount*highCount)
			fmt.Println("count: ", number)
			os.Exit(0)
		}

		if _, ok := modules[s.to]; ok {
			modules[s.to].Notify(s)
		}
	}
}

func (o *orchestrator) Receive(s signal) {
	o.queue = append(o.queue, s)
}

// either I make implementations for each of these with its own thing... or handle everything together for now.
type module struct {
	Type         byte
	Name         string
	Connections  []string // consider making this a map
	Orchestrator *orchestrator

	// receiver chan pulse
	status   bool             // on / off
	remember map[string]pulse // the last signal we remember for each of our connection
}

// Somehow I need a global mutex kind-a thing so the notifies aren't running over each other.

var (
	lowCount  int
	highCount int
)

// Notify is called if a pulse has been planted in the pulses slice.
func (m *module) Notify(s signal) {
	if len(m.Connections) == 0 {
		return
	}

	if m.Type == '&' {
		m.remember[s.from] = s.p // where the pulse came from
		// fmt.Printf("remember %s: %v\n", m.Name, m.remember)

		allHigh := true
		for _, v := range m.remember {
			if v == low {
				allHigh = false

				break
			}
		}

		if allHigh {
			// nh, dr, xm, tr
			// if m.Name == "nh" || m.Name == "dr" || m.Name == "xm" || m.Name == "tr" {
			// fmt.Println("number, name: ", number, m.Name)
			// fmt.Println("------------------")
			// time.Sleep(500 * time.Millisecond)
			// }
			m.sendPulse(low)
		} else {
			m.sendPulse(high)
		}
	}

	if m.Type == '%' {
		if s.p == high {
			return
		}

		if m.status {
			m.sendPulse(low)
		} else {
			m.sendPulse(high)
		}

		// flip
		m.status = !m.status
		// fmt.Printf("flip status of %s to %v\n", m.Name, m.status)
	}
}

func (m *module) sendPulse(p pulse) {
	for _, c := range m.Connections {
		// if _, ok := modules[c]; ok {
		// fmt.Printf("sending pulse %s from %s to %s\n", p, m.Name, c)
		m.Orchestrator.Receive(signal{from: m.Name, to: c, p: p})
	}
}

var number int

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	conjunctions := map[string]struct{}{}
	orc := &orchestrator{}

	for _, line := range split {
		lineSplit := strings.Split(line, " -> ")
		m := module{
			Orchestrator: orc,
		}
		m.Type = lineSplit[0][0]
		switch lineSplit[0][0] {
		case '%':
			m.Name = string(lineSplit[0][1:])
		case '&':
			m.Name = string(lineSplit[0][1:])
			m.remember = make(map[string]pulse)
			conjunctions[m.Name] = struct{}{}
		case 'b':
			m.Name = "broadcaster"
		}

		connections := strings.Split(lineSplit[1], ", ")
		m.Connections = connections
		for _, c := range m.Connections {
			if _, ok := modules[c]; !ok {
				modules[c] = &module{Name: c}
			}
		}

		modules[m.Name] = &m
	}

	// gather conjunctures

	// set up inputs for conjectures.
	for k, v := range modules {
		for _, c := range v.Connections {
			if _, ok := conjunctions[c]; ok {
				modules[c].remember[k] = low
			}
		}
	}

	// fmt.Println(modules["con"], modules["inv"])

	// fmt.Println(modules)
	// obviously have to divide this after we find the repeating frequency
	// need to find the
	// analyze the input and figure out when does xs get a low pulse.
	// in doing that, we'll come closer to understanding when the cycle needs to be repeated and how many times.
	// calculate using part 1 solution how many steps it takes for those conjunctions to loop
	// and the LCM of those 4 numbers will be the step number it takes to get to the end that is rx.
	// Run until xm, nh, dr, tr output a low... And the number of steps LCM of those four is the solution.
	// xm: 3761
	// nh: 3889
	// dr: 3797
	// tr: 3739
	// LCM(3761, 3889, 3797, 3739) = 207652583562007 -> STAR!!
	for {
		number++
		lowCount++ // add the button push
		broadcaster := modules["broadcaster"]
		broadcaster.sendPulse(low)
		orc.Process()

		// fmt.Println("----------")
		// fmt.Println(lowCount, highCount, lowCount*highCount)
		// lowCount = 0
		// highCount = 0
		// time.Sleep(1 * time.Second)
	}
}

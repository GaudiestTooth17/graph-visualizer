package main

import (
	"image/color"
	"os"
	"strconv"
	"strings"
)

// DynamicNetwork holds a graph and a list of sets of node colors at time steps
type DynamicNetwork struct {
	graph        Graph
	stepToColors []StepColors
}

// StepColors is a map of node to state at step in time
type StepColors map[int]color.RGBA

func readDynamicNetwork(file *os.File) DynamicNetwork {
	graph, fileReader := readAdjacencyList(file)

	// read list of changes
	stepToColors := make([]StepColors, 1)
	for {
		line, err := fileReader.ReadString('\n')
		check(err)
		// exit when you read end
		if strings.Compare(line, "end\n") == 0 {
			break
		}
		nodeToColor := make(StepColors)
		// read in states until hitting a blank line
		for len(line) > 1 {
			node, state := lineToInts(line)
			nodeToColor[node] = stateToColor(state)
			line, err = fileReader.ReadString('\n')
			check(err)
		}
		stepToColors = append(stepToColors, nodeToColor)
	}

	return DynamicNetwork{graph: graph, stepToColors: stepToColors}
}

func lineToInts(line string) (int, int) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		panic("Error: " + line)
	}
	node, err := strconv.Atoi(fields[0])
	check(err)

	state, err := strconv.Atoi(fields[1])
	check(err)

	return node, state
}

func stateToColor(state int) color.RGBA {
	switch state {
	case stateS:
		return colorS
	case stateE:
		return colorE
	case stateI:
		return colorI
	case stateR:
		return colorR
	}
	panic("unknown state " + strconv.Itoa(state))
}

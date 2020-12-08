package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Coordinate represents a 2D coordinate and uses float64 to be compatible with pixel
type Coordinate struct {
	x float64
	y float64
}

// Edge contains the IDs of two nodes that are connect to each other
type Edge struct {
	nodeA int
	nodeB int
}

// Graph contains an adjacency list and a map telling where each node is located on a plane
type Graph struct {
	idToCoordinate map[int]Coordinate
	edges          []Edge
}

func readAdjacencyList(fileName string) Graph {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)

	edges := make([]Edge, 0)
	// populate edge list
	for {
		line, err := fileReader.ReadString('\n')
		// stop upon reading an blank line
		if len(line) == 1 {
			break
		} else if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		edge := lineToEdge(line)
		edges = append(edges, edge)
	}

	idToCoordinate := make(map[int]Coordinate)
	// populate idToCoordinate
	for {
		line, err := fileReader.ReadString('\n')
		// stop at EOF
		if err != nil && errors.Is(err, io.EOF) {
			break
		} else if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		nodeID, coordinate := lineToCoordinate(line)
		idToCoordinate[nodeID] = coordinate
	}

	graph := Graph{idToCoordinate: idToCoordinate, edges: edges}
	return graph
}

func lineToEdge(line string) Edge {
	coordinates := strings.Fields(line)
	if len(coordinates) != 2 {
		panic("Error: " + line)
	}
	a, err := strconv.Atoi(coordinates[0])
	check(err)

	b, err := strconv.Atoi(coordinates[1])
	check(err)

	return Edge{nodeA: a, nodeB: b}
}

func lineToCoordinate(line string) (int, Coordinate) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		panic("Malformatted line: " + line)
	}

	nodeID, err := strconv.Atoi(fields[0])
	check(err)

	x, err := strconv.ParseFloat(fields[1], 64)
	check(err)

	y, err := strconv.ParseFloat(fields[2], 64)
	check(err)

	return nodeID, Coordinate{x: x, y: y}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

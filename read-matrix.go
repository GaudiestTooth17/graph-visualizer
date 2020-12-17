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

// readAdjacencyList takes in the name of a file to read and outputs a graph
// and the Reader that it used so that you can continue reading changes from
// another file if you want to.
func readAdjacencyList(file *os.File) (Graph, *bufio.Reader) {
	fileReader := bufio.NewReader(file)

	// read and discard the first line of the file specifying the number of nodes
	// this information is determined implicitly in this program,
	// but not in infection-resistant-network
	_, err := fileReader.ReadString('\n')
	check(err)

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
		// stop at the first blank line
		if len(line) == 1 {
			break
		} else if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}
		nodeID, coordinate := lineToCoordinate(line)
		idToCoordinate[nodeID] = coordinate
	}

	graph := Graph{idToCoordinate: idToCoordinate, edges: edges}
	return graph, fileReader
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

// lineToCoordinate takes a line formatted as [nodeID] [x location] [y location]
// the x and y locations are between -1 and 1
// it returns the nodeID and (x, y)
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

	x, y = normalizeXY(x, y)

	return nodeID, Coordinate{x: x, y: y}
}

// take x and y coordinates in the range [-1, 1] and convert them to be within the
// actual bounds of the window
func normalizeXY(x, y float64) (float64, float64) {
	return ((x + 1) / 2) * maxX, ((y + 1) / 2) * maxY
}

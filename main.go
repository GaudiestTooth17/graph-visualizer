package main

import (
	"fmt"
	"os"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <graph-name>\n", os.Args[0])
		return
	}

	pixelgl.Run(run)
}

func run() {
	graphName := os.Args[1]
	graphFile, err := os.Open(graphName)
	check(err)
	window, imd := makeWindow(graphName)

	// code for static graphs
	// graph, _ := readAdjacencyList(graphName)
	// addEdges(graph, window, imd)
	// addNodes(graph, window, imd)

	// for !window.Closed() {
	// 	drawGraph(graph, window, imd)
	// }

	// code for dynamic graphs
	dynamicGraph := readDynamicNetwork(graphFile)
	graphFile.Close()

	writeStep := newStepWriter()
	step := 1
	for !window.Closed() {
		if window.Pressed(pixelgl.KeyLeft) {
			step = max(step-1, 1)
		} else if window.Pressed(pixelgl.KeyRight) {
			step = min(step+1, len(dynamicGraph.stepToColors)-1)
		} else if window.JustPressed(pixelgl.KeyA) {
			step = max(step-1, 1)
		} else if window.JustPressed(pixelgl.KeyD) {
			step = min(step+1, len(dynamicGraph.stepToColors)-1)
		}
		writeStep(step, window)
		drawDynamicGraph(dynamicGraph, step, window, imd)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

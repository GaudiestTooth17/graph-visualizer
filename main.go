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

	// graph, _ := readAdjacencyList(graphName)
	// addEdges(graph, window, imd)
	// addNodes(graph, window, imd)

	// for !window.Closed() {
	// 	drawGraph(graph, window, imd)
	// }

	dynamicGraph := readDynamicNetwork(graphFile)
	graphFile.Close()

	for !window.Closed() {
		for step := 0; step < len(dynamicGraph.stepToColors) && !window.Closed(); step++ {
			drawDynamicGraph(dynamicGraph, step, window, imd)
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

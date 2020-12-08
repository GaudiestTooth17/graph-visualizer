package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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
	window, imd := makeWindow(graphName)
	graph := readAdjacencyList(graphName)
	addEdges(graph, window, imd)
	addNodes(graph, window, imd)

	for !window.Closed() {
		window.Clear(colornames.Aliceblue)
		imd.Draw(window)
		time.Sleep(40 * time.Millisecond)
		window.Update()
	}
}

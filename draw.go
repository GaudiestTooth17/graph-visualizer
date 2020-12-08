package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	radius        = 2
	edgeThickness = 1
	maxX          = 1500
	maxY          = 900
	stateS        = iota
	stateE        = iota
	stateI        = iota
	stateR        = iota
)

var colorS = colornames.Blue
var colorE = colornames.Green
var colorI = colornames.Crimson
var colorR = colornames.Blueviolet

func makeWindow(title string) (*pixelgl.Window, *imdraw.IMDraw) {
	cfg := pixelgl.WindowConfig{
		Title:     title,
		Bounds:    pixel.R(0, 0, maxX, maxY),
		VSync:     true,
		Resizable: true,
	}
	window, err := pixelgl.NewWindow(cfg)
	check(err)

	imd := imdraw.New(nil)

	return window, imd
}

func addNodes(graph Graph, window *pixelgl.Window, imd *imdraw.IMDraw) {
	imd.Color = colorS
	for _, coordinate := range graph.idToCoordinate {
		imd.Push(pixel.V(coordinate.x, coordinate.y))
		imd.Circle(radius, 0)
	}
	imd.Draw(window)
}

func addEdges(graph Graph, window *pixelgl.Window, imd *imdraw.IMDraw) {
	imd.Color = colornames.Black
	for _, edge := range graph.edges {
		coordinateA := graph.idToCoordinate[edge.nodeA]
		coordinateB := graph.idToCoordinate[edge.nodeB]

		imd.Push(pixel.V(coordinateA.x, coordinateA.y))
		imd.Push(pixel.V(coordinateB.x, coordinateB.y))
		imd.Line(edgeThickness)
	}
	imd.Draw(window)
}

func updateColors(nodeStates map[int]int, window *pixelgl.Window, imd *imdraw.IMDraw) {
	//
}

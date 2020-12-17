package main

import (
	"time"

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
	stateS        = 0
	stateE        = 1
	stateI        = 2
	stateR        = 3
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

func addDynamicNodes(dynamicGraph DynamicNetwork, step int, window *pixelgl.Window, imd *imdraw.IMDraw) {
	for node, coordinate := range dynamicGraph.graph.idToCoordinate {
		imd.Color = dynamicGraph.stepToColors[step][node]
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

func drawGraph(graph Graph, window *pixelgl.Window, imd *imdraw.IMDraw) {
	window.Clear(colornames.Aliceblue)
	imd.Draw(window)
	time.Sleep(40 * time.Millisecond)
	window.Update()
}

func drawDynamicGraph(graph DynamicNetwork, step int, window *pixelgl.Window, imd *imdraw.IMDraw) {
	window.Clear(colornames.Aliceblue)
	imd.Clear()
	addEdges(graph.graph, window, imd)
	addDynamicNodes(graph, step, window, imd)
	imd.Draw(window)
	time.Sleep(250 * time.Millisecond)
	window.Update()
}

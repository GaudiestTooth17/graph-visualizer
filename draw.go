package main

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	radius        = 5
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
var colorR = colornames.Slategrey
var colorBG = colornames.Whitesmoke
var colorEdges = colornames.Black
var colorText = colornames.Black

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
	imd.Color = colorEdges
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
	window.Clear(colorBG)
	imd.Draw(window)
	time.Sleep(40 * time.Millisecond)
	window.Update()
}

func drawDynamicGraph(graph DynamicNetwork, step int, window *pixelgl.Window, imd *imdraw.IMDraw) {
	window.Clear(colorBG)
	imd.Clear()
	addEdges(graph.graph, window, imd)
	addDynamicNodes(graph, step, window, imd)
	imd.Draw(window)
	time.Sleep(40 * time.Millisecond)
	window.Update()
}

func newStepWriter() func(int, *pixelgl.Window) {
	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	screenText := text.New(pixel.V(10, 10), atlas)
	screenText.Color = colorText

	return func(step int, window *pixelgl.Window) {
		screenText.Clear()
		fmt.Fprintf(screenText, "Step %d", step)
		screenText.Draw(window, pixel.IM)
		window.Update()
	}
}

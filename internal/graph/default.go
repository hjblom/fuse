package graph

import "github.com/dominikbraun/graph"

var defaultNodeAttributes = []func(*graph.VertexProperties){
	graph.VertexAttribute("style", "filled"),
	graph.VertexAttribute("shape", "box"),
	graph.VertexAttribute("fillcolor", "orange"),
	graph.VertexAttribute("width", "1.5"),
	graph.VertexAttribute("height", "0.5"),
}

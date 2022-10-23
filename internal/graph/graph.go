package graph

import (
	"bytes"
	"fmt"

	"github.com/hjblom/fuse/internal/config"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type Graph struct {
	components map[string]config.Component
	graph      graph.Graph[string, string]
}

func NewGraph() *Graph {
	g := &Graph{
		components: make(map[string]config.Component),
		graph:      graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles()),
	}
	return g
}

func (g *Graph) AddComponents(components []config.Component) error {
	for _, component := range components {
		if _, ok := g.components[component.Name]; ok {
			return fmt.Errorf("component %s already exists", component.Name)
		}

		// Add component to reverse lookup map
		g.components[component.Name] = component

		// Add vertex to graph
		err := g.addVertex(component.Name)
		if err != nil {
			return fmt.Errorf("failed to add vertex %s: %w", component.Name, err)
		}
	}

	for _, component := range components {
		for _, req := range component.Requires {
			err := g.addEdge(req, component.Name)
			if err != nil {
				return fmt.Errorf("failed to add edge between %s and %s: %w", component.Name, req, err)
			}
		}
	}

	return nil
}

func (g *Graph) AddComponent(component config.Component) error {
	err := g.addVertex(component.Name)
	if err != nil {
		return fmt.Errorf("failed to add vertex %s: %w", component.Name, err)
	}

	for _, req := range component.Requires {
		err := g.addEdge(req, component.Name)
		if err != nil {
			return fmt.Errorf("failed to add edge between %s and %s: %w", component.Name, req, err)
		}
	}

	return nil
}

func (g *Graph) TopologicalSort() ([]string, error) {
	return graph.TopologicalSort(g.graph)
}

func (g *Graph) ToSVG() ([]byte, error) {
	// Convert to DOT
	var dot bytes.Buffer
	err := draw.DOT(g.graph, &dot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert graph to DOT: %w", err)
	}
	graph, err := graphviz.ParseBytes(dot.Bytes())
	if err != nil {
		return nil, err
	}

	// Set graph attributes
	setGraphAttributes(graph)

	// Convert to SVG
	gv := graphviz.New()
	defer gv.Close()

	var svg bytes.Buffer
	err = gv.Render(graph, graphviz.SVG, &svg)
	if err != nil {
		return nil, err
	}

	return svg.Bytes(), nil
}

func (g *Graph) addVertex(name string) error {
	return g.graph.AddVertex(name, defaultNodeAttributes...)
}

func (g *Graph) addEdge(src, dst string) error {
	ln := graph.EdgeAttribute("minlen", "2")
	return g.graph.AddEdge(src, dst, ln)
	// return g.graph.AddEdge(src, dst)
}

func setGraphAttributes(graph *cgraph.Graph) {
	// graph.SetRankDir("LR")
	graph.SetRankSeparator(0.5)
	graph.SetNodeSeparator(1.5)
}

package config

import (
	"bytes"
	"fmt"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

var defaultNodeAttributes = []func(*graph.VertexProperties){
	graph.VertexAttribute("style", "filled"),
	graph.VertexAttribute("shape", "box"),
	graph.VertexAttribute("fillcolor", "orange"),
	graph.VertexAttribute("width", "1.5"),
	graph.VertexAttribute("height", "0.5"),
}

var defaultEdgeAttributes = []func(*graph.EdgeProperties){
	graph.EdgeAttribute("minlen", "2"),
}

type Module struct {
	Path     string     `yaml:"path"`
	Packages []*Package `yaml:"packages"`

	packages map[string]*Package
	graph    graph.Graph[string, string]
}

func NewModule(path string) *Module {
	return &Module{
		Path:     path,
		Packages: []*Package{},
	}
}

func (c *Module) Validate() error {
	// If the graph has been populated, assume validation has been done.
	if c.graph != nil {
		return nil
	} else {
		// Else populate the graph and validate it.
		c.packages = make(map[string]*Package)
		c.graph = graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())
	}

	// Validate module name
	if c.Path == "" {
		return fmt.Errorf("module name is required")
	}

	// Validate packages
	for _, pkg := range c.Packages {
		if _, ok := c.packages[pkg.RelativePath()]; ok {
			return fmt.Errorf("package %s already exists", pkg.RelativePath())
		}

		// Add pkg to reverse lookup map
		c.packages[pkg.RelativePath()] = pkg

		// Add vertex to graph
		err := c.graph.AddVertex(pkg.RelativePath(), defaultNodeAttributes...)
		if err != nil {
			delete(c.packages, pkg.RelativePath())
			return fmt.Errorf("failed to add vertex %s: %w", pkg.RelativePath(), err)
		}
	}

	for _, pkg := range c.Packages {
		for _, req := range pkg.Requires {
			err := c.graph.AddEdge(req, pkg.RelativePath(), defaultEdgeAttributes...)
			if err != nil {
				return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.RelativePath(), req, err)
			}
		}
	}

	return nil
}

func (c *Module) AddPackage(pkg *Package) error {
	// Ensure that the graph has been validated
	if err := c.Validate(); err != nil {
		return err
	}

	// If package already exists, return an error
	if _, ok := c.packages[pkg.RelativePath()]; ok {
		return fmt.Errorf("package %s already exists", pkg.RelativePath())
	}

	// Add pkg to reverse lookup map
	c.packages[pkg.RelativePath()] = pkg

	// Add vertex to graph
	err := c.graph.AddVertex(pkg.RelativePath(), defaultNodeAttributes...)
	if err != nil {
		delete(c.packages, pkg.RelativePath())
		return fmt.Errorf("failed to add vertex %s: %w", pkg.RelativePath(), err)
	}

	// Add edges to graph
	for _, req := range pkg.Requires {
		err := c.graph.AddEdge(req, pkg.RelativePath(), defaultEdgeAttributes...)
		if err != nil {
			return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.RelativePath(), req, err)
		}
	}

	// Add pkg to config
	c.Packages = append(c.Packages, pkg)

	return nil
}

func (c *Module) TopologicalPackageOrder() ([]*Package, error) {
	// Ensure that the graph has been validated
	if err := c.Validate(); err != nil {
		return nil, err
	}

	// Do a topological sort of the graph
	packages, err := graph.TopologicalSort(c.graph)
	if err != nil {
		return nil, fmt.Errorf("failed to sort packages: %w", err)
	}

	// Use result to reverse lookup slice of packages
	result := make([]*Package, len(packages))
	for idx, pkg := range packages {
		result[idx] = c.packages[pkg]
	}

	return result, nil
}

func (c *Module) ToDOT() ([]byte, error) {
	// Ensure that the graph has been validated
	if err := c.Validate(); err != nil {
		return nil, err
	}

	var dot bytes.Buffer
	err := draw.DOT(c.graph, &dot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert graph to DOT: %w", err)
	}
	return dot.Bytes(), nil
}

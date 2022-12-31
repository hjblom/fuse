package config

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

type Module struct {
	Path     string     `yaml:"path"`
	Packages []*Package `yaml:"packages,omitempty"`

	// Reverse lookup map of packages [id]
	packages map[string]*Package
	paths    map[string]bool
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
	}

	// Validate module name
	if c.Path == "" {
		return fmt.Errorf("module name is required")
	}

	// Create instance of graph and packages
	c.packages = make(map[string]*Package)
	c.paths = make(map[string]bool)
	c.graph = graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())

	// Validate packages
	for _, pkg := range c.Packages {
		if _, ok := c.packages[pkg.ID]; ok {
			return fmt.Errorf("cannot add package %s with id %s, already exists", pkg.Name, pkg.ID)
		}
		if _, ok := c.paths[pkg.FullPath(c.Path)]; ok {
			return fmt.Errorf("cannot add package %s with path %s, already exists", pkg.Name, pkg.FullPath(c.Path))
		}

		// Add pkg to reverse lookup map
		c.packages[pkg.ID] = pkg
		c.paths[pkg.FullPath(c.Path)] = true

		// Add vertex to graph
		err := c.graph.AddVertex(pkg.ID, defaultNodeAttributes...)
		if err != nil {
			delete(c.packages, pkg.ID)
			return fmt.Errorf("failed to add vertex %s: %w", pkg.ID, err)
		}
	}

	for _, pkg := range c.Packages {
		for _, req := range pkg.Requires {
			err := c.graph.AddEdge(req, pkg.ID, defaultEdgeAttributes...)
			if err != nil {
				return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.ID, req, err)
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
		delete(c.packages, pkg.Name)
		return fmt.Errorf("failed to add vertex %s: %w", pkg.Name, err)
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

	// Convert to dot
	var dot bytes.Buffer
	err := draw.DOT(c.graph, &dot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert graph to DOT: %w", err)
	}
	return dot.Bytes(), nil
}

// ToSVG attempts to convert the module's graph to SVG using the dot command.
// ToSVG's logic is based on https://github.com/google/pprof/blob/main/internal/driver/webui.go#L339
func (c *Module) ToSVG() ([]byte, error) {
	// Ensure that the graph has been validated
	if err := c.Validate(); err != nil {
		return nil, err
	}

	// Convert graph to DOT
	dot, err := c.ToDOT()
	if err != nil {
		return nil, err
	}

	// Convert DOT to SVG using dot command
	cmd := exec.Command("dot", "-Tsvg")
	svg := &bytes.Buffer{}
	cmd.Stdin = bytes.NewReader(dot)
	cmd.Stderr = os.Stderr
	cmd.Stdout = svg
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("dot program failed to execute: %w", err)
	}
	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("dot program failed to execute: %w", err)
	}

	return svg.Bytes(), nil
}

func (c *Module) GetPackage(pkgName string) *Package {
	// Ensure that the graph has been validated
	if err := c.Validate(); err != nil {
		return nil
	}

	pkg, ok := c.packages[pkgName]
	if !ok {
		return nil
	}
	return pkg
}

func (c *Module) GetPackageOutDegree(pkg *Package) int {
	am, err := c.graph.AdjacencyMap()
	if err != nil {
		return 0
	}
	if p, ok := am[pkg.ID]; ok {
		return len(p)
	}
	return 0
}

var defaultNodeAttributes = []func(*graph.VertexProperties){
	graph.VertexAttribute("style", "filled"),
	graph.VertexAttribute("shape", "box"),
	graph.VertexAttribute("fillcolor", "lightblue1"),
	graph.VertexAttribute("width", "1.5"),
	graph.VertexAttribute("height", "0.5"),
}

var defaultEdgeAttributes = []func(*graph.EdgeProperties){
	graph.EdgeAttribute("minlen", "2"),
}

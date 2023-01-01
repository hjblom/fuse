package config

import (
	"bytes"
	"fmt"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

var reservedAliases = map[string]bool{
	// Config c
	"c": true,
}

type Module struct {
	Path     string     `yaml:"path"`
	Packages []*Package `yaml:"packages,omitempty"`

	packages map[string]*Package
	paths    map[string]bool
	aliases  map[string]bool
	graph    graph.Graph[string, string]
}

func NewModule() *Module {
	m := &Module{
		Packages: []*Package{},
	}
	return m
}

func (m *Module) Validate() error {
	// Validate module name
	if m.Path == "" {
		return fmt.Errorf("module name is required")
	}

	m.packages = make(map[string]*Package)
	m.paths = make(map[string]bool)
	m.aliases = make(map[string]bool)
	m.graph = graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())

	// Validate packages
	for _, pkg := range m.Packages {
		if _, ok := m.packages[pkg.ID]; ok {
			return fmt.Errorf("cannot add package \"%s\" with id \"%s\", already exists", pkg.Name, pkg.ID)
		}
		if _, ok := m.paths[pkg.FullPath(m.Path)]; ok {
			return fmt.Errorf("cannot add package \"%s\" with path \"%s\", already exists", pkg.Name, pkg.FullPath(m.Path))
		}
		if _, ok := m.aliases[pkg.Alias]; ok {
			return fmt.Errorf("cannot add package \"%s\" with alias \"%s\", already exists", pkg.Name, pkg.Alias)
		}
		if _, ok := reservedAliases[pkg.Alias]; ok {
			return fmt.Errorf("cannot add package \"%s\" with alias \"%s\", reserved", pkg.Name, pkg.Alias)
		}

		// Add pkg to reverse lookup map
		m.packages[pkg.ID] = pkg
		m.paths[pkg.FullPath(m.Path)] = true

		// Add vertex to graph
		err := m.graph.AddVertex(pkg.ID, defaultNodeAttributes...)
		if err != nil {
			return fmt.Errorf("failed to add vertex %s: %w", pkg.ID, err)
		}
	}

	for _, pkg := range m.Packages {
		for _, req := range pkg.Requires {
			err := m.graph.AddEdge(req, pkg.ID, defaultEdgeAttributes...)
			if err != nil {
				return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.ID, req, err)
			}
		}
	}

	return nil
}

func (m *Module) AddPackage(pkg *Package) error {
	if _, ok := m.packages[pkg.ID]; ok {
		return fmt.Errorf("cannot add package \"%s\" with id \"%s\", already exists", pkg.Name, pkg.ID)
	}
	if _, ok := m.paths[pkg.FullPath(m.Path)]; ok {
		return fmt.Errorf("cannot add package \"%s\" with path \"%s\", already exists", pkg.Name, pkg.FullPath(m.Path))
	}
	if _, ok := m.aliases[pkg.Alias]; ok {
		return fmt.Errorf("cannot add package \"%s\" with alias \"%s\", already exists", pkg.Name, pkg.Alias)
	}
	if _, ok := reservedAliases[pkg.Alias]; ok {
		return fmt.Errorf("cannot add package \"%s\" with alias \"%s\", reserved", pkg.Name, pkg.Alias)
	}

	// Add vertex to graph
	err := m.graph.AddVertex(pkg.RelativePath(), defaultNodeAttributes...)
	if err != nil {
		return fmt.Errorf("failed to add vertex %s: %w", pkg.Name, err)
	}

	// Add edges to graph
	for _, req := range pkg.Requires {
		err := m.graph.AddEdge(req, pkg.RelativePath(), defaultEdgeAttributes...)
		if err != nil {
			return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.RelativePath(), req, err)
		}
	}

	// Add pkg to config
	m.packages[pkg.RelativePath()] = pkg
	m.paths[pkg.FullPath(m.Path)] = true
	m.aliases[pkg.Alias] = true
	m.Packages = append(m.Packages, pkg)

	return nil
}

func (m *Module) TopologicalPackageOrder() ([]*Package, error) {
	// Do a topological sort of the graph
	packages, err := graph.TopologicalSort(m.graph)
	if err != nil {
		return nil, fmt.Errorf("failed to sort packages: %w", err)
	}

	// Use result to reverse lookup slice of packages
	result := make([]*Package, len(packages))
	for idx, pkg := range packages {
		result[idx] = m.packages[pkg]
	}

	return result, nil
}

func (m *Module) GetPackage(pkgName string) *Package {
	pkg, ok := m.packages[pkgName]
	if !ok {
		return nil
	}
	return pkg
}

func (m *Module) GetPackageOutDegree(pkg *Package) int {
	am, err := m.graph.AdjacencyMap()
	if err != nil {
		return 0
	}
	if p, ok := am[pkg.ID]; ok {
		return len(p)
	}
	return 0
}

func (m *Module) ToDOT() ([]byte, error) {
	// Convert to dot
	var dot bytes.Buffer
	err := draw.DOT(m.graph, &dot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert graph to DOT: %w", err)
	}
	return dot.Bytes(), nil
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

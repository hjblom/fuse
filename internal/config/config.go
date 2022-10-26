package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hjblom/fuse/internal/util/osi"
	"gopkg.in/yaml.v3"

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

type Config struct {
	Module   string     `yaml:"module"`
	Packages []*Package `yaml:"packages"`

	packages map[string]*Package
	graph    graph.Graph[string, string]

	// Reading / Writing
	os osi.Interface
}

type Package struct {
	Name     string   `yaml:"package"`
	Path     string   `yaml:"path,omitempty"`
	Tags     []string `yaml:"tags,omitempty"`
	Requires []string `yaml:"requires,omitempty"`
}

func NewConfig() *Config {
	return &Config{
		os: osi.NewOS(),
	}
}

func NewPackage(name, path string, requires, tags []string) *Package {
	return &Package{
		Name:     name,
		Path:     path,
		Requires: requires,
		Tags:     tags,
	}
}

func (c *Config) Read(path string) error {
	// Read file
	data, err := c.os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

func (c *Config) Write(path string) error {
	// Setup YAML encoder
	data := &bytes.Buffer{}
	enc := yaml.NewEncoder(data)
	defer enc.Close()
	enc.SetIndent(2)

	// Marshal
	err := enc.Encode(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write file
	err = c.os.WriteFile(path, data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func (c *Config) Validate() error {
	// If the graph has been populated, assume validation has been done.
	if c.graph != nil {
		return nil
	} else {
		// Else populate the graph and validate it.
		c.packages = make(map[string]*Package)
		c.graph = graph.New(graph.StringHash, graph.Directed(), graph.PreventCycles())
	}

	// Validate module name
	if c.Module == "" {
		return fmt.Errorf("module name is required")
	}

	// Validate packages
	for _, pkg := range c.Packages {
		if _, ok := c.packages[pkg.FullPath()]; ok {
			return fmt.Errorf("package %s already exists", pkg.FullPath())
		}

		// Add pkg to reverse lookup map
		c.packages[pkg.FullPath()] = pkg

		// Add vertex to graph
		err := c.graph.AddVertex(pkg.FullPath(), defaultNodeAttributes...)
		if err != nil {
			delete(c.packages, pkg.FullPath())
			return fmt.Errorf("failed to add vertex %s: %w", pkg.FullPath(), err)
		}
	}

	for _, pkg := range c.Packages {
		for _, req := range pkg.Requires {
			err := c.graph.AddEdge(req, pkg.FullPath(), defaultEdgeAttributes...)
			if err != nil {
				return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.FullPath(), req, err)
			}
		}
	}

	return nil
}

func (c *Config) AddPackage(pkg *Package) error {
	// Validate config
	if err := c.Validate(); err != nil {
		return err
	}

	// If package already exists, return an error
	if _, ok := c.packages[pkg.FullPath()]; ok {
		return fmt.Errorf("package %s already exists", pkg.FullPath())
	}

	// Add pkg to reverse lookup map
	c.packages[pkg.FullPath()] = pkg

	// Add vertex to graph
	err := c.graph.AddVertex(pkg.FullPath(), defaultNodeAttributes...)
	if err != nil {
		delete(c.packages, pkg.FullPath())
		return fmt.Errorf("failed to add vertex %s: %w", pkg.FullPath(), err)
	}

	// Add edges to graph
	for _, req := range pkg.Requires {
		err := c.graph.AddEdge(req, pkg.FullPath(), defaultEdgeAttributes...)
		if err != nil {
			return fmt.Errorf("failed to add edge between %s and %s: %w", pkg.FullPath(), req, err)
		}
	}

	// Add pkg to config
	c.Packages = append(c.Packages, pkg)

	return nil
}

func (c *Config) TopologicalSort() ([]*Package, error) {
	packages, err := graph.TopologicalSort(c.graph)
	if err != nil {
		return nil, fmt.Errorf("failed to sort packages: %w", err)
	}

	result := make([]*Package, len(packages))
	for idx, pkg := range packages {
		result[idx] = c.packages[pkg]
	}

	return result, nil
}

func (c *Config) ToDOT() ([]byte, error) {
	var dot bytes.Buffer
	err := draw.DOT(c.graph, &dot)
	if err != nil {
		return nil, fmt.Errorf("failed to convert graph to DOT: %w", err)
	}
	return dot.Bytes(), nil
}

func (p *Package) FullPath() string {
	return fmt.Sprintf("%s/%s", p.Path, p.Name)
}

func (p *Package) GoPackageName() string {
	return strings.ToUpper(p.Name[0:1]) + string(p.Name[1:])
}

func (p *Package) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("package name is required")
	}
	return nil
}

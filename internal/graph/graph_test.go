package graph

import (
	"testing"

	"github.com/hjblom/fuse/internal/config"
	"github.com/stretchr/testify/assert"
)

var testComponentsGood = []config.Package{
	{
		Name:     "Server",
		Requires: []string{"Client"},
	},
	{
		Name: "Client",
	},
}

var testComponentsLoop = []config.Package{
	{
		Name:     "A",
		Requires: []string{"B"},
	},
	{
		Name:     "B",
		Requires: []string{"A"},
	},
}

var testComponentsRepeated = []config.Package{
	{
		Name: "A",
	},
	{
		Name: "A",
	},
}

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	assert.NotNil(t, g)
}

func TestAddComponentsSuccess(t *testing.T) {
	g := NewGraph()
	err := g.AddComponents(testComponentsGood)
	assert.NoError(t, err)
}

func TestAddComponentsErrorLoop(t *testing.T) {
	g := NewGraph()
	err := g.AddComponents(testComponentsLoop)
	assert.Error(t, err)
}

func TestAddComponentsErrorRepeated(t *testing.T) {
	g := NewGraph()
	err := g.AddComponents(testComponentsRepeated)
	assert.Error(t, err)
}

func TestAddComponentSuccess(t *testing.T) {
	// Given
	g := NewGraph()
	g.AddComponents(testComponentsGood)

	// When
	err := g.AddComponent(config.Package{
		Name:     "Database",
		Requires: []string{"Server"},
	})

	// Then
	assert.NoError(t, err)
}

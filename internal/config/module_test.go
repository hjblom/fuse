package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	_ "embed"
)

var testConfig = `
module:
  path: path
  packages:
    - package: client
      path: internal
    - package: server
      path: internal/services
      requires:
        - internal/client
`

func TestConfigParse(t *testing.T) {
	c := Module{}
	err := yaml.Unmarshal([]byte(testConfig), &c)
	if err != nil {
		t.Errorf("failed to unmarshal config: %v", err)
	}

	assert.NoError(t, err)
}

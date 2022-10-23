package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	_ "embed"
)

var testConfig = `
components:
  - name: Server
    tags:
      - service
    requires:
      - Client
      - Client2
  - name: Client
    requires:
      - ConfigManager
      - ConfigManager2
  - name: ConfigManager
  - name: ConfigManager2
  - name: Client2
    requires:
     - ConfigManager2
  - name: Logger
    tags:
     - singleton
`

func TestConfigParse(t *testing.T) {
	c := Config{}
	err := yaml.Unmarshal([]byte(testConfig), &c)
	if err != nil {
		t.Errorf("failed to unmarshal config: %v", err)
	}

	assert.NoError(t, err)
}

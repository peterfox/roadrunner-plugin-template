package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigInitDefaults(t *testing.T) {
	cfg := Config{}

	cfg.InitDefaults()

	assert.Equal(t, "foobar", cfg.Value)
}

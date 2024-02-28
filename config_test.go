package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This will examine the config global, and ensure the values match config.yaml
func TestLoadConfig(t *testing.T) {

	assert.IsType(t, globalConfig{}, config, "type check")

	if reflect.DeepEqual(config, new(globalConfig)) {
		t.Error("config is empty.")
	}

	assert.False(t, config.Collector.Debug, "Configuration - Collector.Debug")
	assert.False(t, config.Collector.SampleData, "Configuration - Collector.SampleData")

	assert.NotEmpty(t, config.Vrops.Host, "Configuration - Vrops.Host")
	assert.NotEmpty(t, config.Vrops.PageSize, "Configuration - Vrops.PageSize")
	assert.NotEmpty(t, config.Vrops.Pass, "Configuration - Vrops.Pass")
	assert.NotEmpty(t, config.Vrops.ResourceKindKeys, "Configuration - Vrops.ResourceKindKeys")
	assert.Equal(t, "username", config.Vrops.User, "Configuration - Vrops.User")

}

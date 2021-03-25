package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigs(t *testing.T) {
	expectConfig := config{
		ChartsOutputDir:      "test-charts",
		ChartsManifestFile:   "test_chart_manifest.yaml",
		ChartsValuesDir:      "test-charts/values",
		ClusterChartsBaseDir: "test-base/charts",
	}

	os.Setenv("HELKU_OUTPUT_DIR", expectConfig.ChartsOutputDir)
	os.Setenv("HELKU_CHARTS_MANIFEST_FILE", expectConfig.ChartsManifestFile)
	os.Setenv("HELKU_CHARTS_VALUES_DIR", expectConfig.ChartsValuesDir)
	os.Setenv("HELKU_CLUSTER_BASE_DIR", expectConfig.ClusterChartsBaseDir)

	cfg.loadConfigs()
	assert := assert.New(t)
	assert.Equal(expectConfig, cfg, "Expected config to be equal")
}

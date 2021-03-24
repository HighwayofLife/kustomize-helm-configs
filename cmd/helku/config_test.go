package main

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadConfigs(t *testing.T) {
	expectConfig := config{
		ChartsOutputDir:    "test-charts",
		ChartsManifestFile: "test_chart_manifest.yaml",
		ChartsValuesDir:    "test-charts/values",
		ClusterBaseDir:     "test-base/charts",
	}

	os.Setenv("HELKU_OUTPUT_DIR", expectConfig.ChartsOutputDir)
	os.Setenv("HELKU_CHARTS_MANIFEST_FILE", expectConfig.ChartsManifestFile)
	os.Setenv("HELKU_CHARTS_VALUES_DIR", expectConfig.ChartsValuesDir)
	os.Setenv("HELKU_CLUSTER_BASE_DIR", expectConfig.ClusterBaseDir)

	cfg.loadConfigs()
	assert.Equal(t, expectConfig, cfg)
}

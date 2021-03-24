package main

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadConfigs(t *testing.T) {
	expectConfig := config{
		ChartsOutputDir:    "test_charts",
		ChartsManifestFile: "test_chart_manifest.yaml",
	}

	os.Setenv("HELM_KUSTOMIZE_OUTPUT_DIR", expectConfig.ChartsOutputDir)
	os.Setenv("HELM_KUSTOMIZE_CHARTS_MANIFEST_FILE", expectConfig.ChartsManifestFile)

	cfg.loadConfigs()
	assert.Equal(t, expectConfig, cfg)
	os.Unsetenv("HELM_KUSTOMIZE_OUTPUT_DIR")
	os.Unsetenv("HELM_KUSTOMIZE_CHARTS_MANIFEST_FILE")
}

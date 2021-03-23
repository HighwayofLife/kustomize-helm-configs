package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadChartsManifest(t *testing.T) {
	cfg.ChartsOutputDir = "../../charts"
	cfg.ChartsManifestFile = "../../chart_manifest.yaml"

	_, err := loadChartsManifest()
	assert.Assert(t, err)
	assert.NilError(t, err)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
}

func TestDownloadHelmCharts(t *testing.T) {
	var charts Charts
	cfg.ChartsOutputDir = "../../charts"
	cfg.ChartsManifestFile = "../../chart_manifest.yaml"

	charts.Charts = append(charts.Charts, Chart{
		Repo: "https://helm.datadoghq.com",
		Name: "datadog",
	})

	err := downloadHelmCharts(charts)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
}

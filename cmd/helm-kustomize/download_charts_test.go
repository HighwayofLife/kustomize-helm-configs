package main

import (
	"io/ioutil"
	"os"
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
	var (
		chartRepo = "https://helm.datadoghq.com"
		chartName = "datadog"
	)

	charts.Charts = append(charts.Charts, Chart{
		Repo: chartRepo,
		Name: chartName,
	})

	err := downloadHelmCharts(charts)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	if _, err := os.Stat(cfg.ChartsOutputDir + "/" + chartName); os.IsNotExist(err) {
		t.Errorf("Expected chart directory %s to exist, but does not", chartName)
	}
}

func TestRemoveFiles(t *testing.T) {
	dir := "/tmp/test"
	file := dir + "/write-test"

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		t.Errorf("Expected to create directories: %s. Got error: %s", dir, err.Error())
	}

	bytes := []byte("test\n")
	ioutil.WriteFile(file, bytes, 0644)

	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("Expected to create file: %s, got error: %s", file, err.Error())
	}

	err = removeFiles(dir)
	if err != nil {
		t.Errorf("Expected to remove files: %s, got error %s", dir, err.Error())
	}

	if _, err := os.Stat(file); err == nil {
		t.Errorf("Expected to remove file: %s, but file still exists!", file)
	}

	err = removeFiles("/tmp/fake-file-test")
	if err != nil {
		t.Errorf("Tried to remove non-existant file, should ignore, got error: %s", err.Error())
	}
}

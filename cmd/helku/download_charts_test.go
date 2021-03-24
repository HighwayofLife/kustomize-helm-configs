package main

import (
	"io/ioutil"
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestLoadChartsManifest(t *testing.T) {
	cfg.ChartsManifestFile = "../../chart_manifest.yaml"

	_, err := loadChartsManifest()
	assert.Assert(t, err)
	assert.NilError(t, err)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
}

func testChartsHelper() Charts {
	var charts Charts

	cfg.ChartsOutputDir = "../../test-charts"
	cfg.ChartsValuesDir = cfg.ChartsOutputDir + "/values"
	cfg.ChartsManifestFile = "../../chart_manifest.yaml"

	charts.Charts = append(charts.Charts, Chart{
		Repo: "https://helm.datadoghq.com",
		Name: "datadog",
	})

	return charts
}

func TestDownloadHelmCharts(t *testing.T) {
	charts := testChartsHelper()

	err := downloadHelmCharts(charts)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}

	chartName := charts.Charts[0].Name

	if _, err := os.Stat(cfg.ChartsOutputDir + "/" + chartName); os.IsNotExist(err) {
		t.Errorf("Expected chart directory %s to exist, but does not", chartName)
	}

	t.Cleanup(func() {
		os.RemoveAll(cfg.ChartsOutputDir)
	})
}

func TestCreateValuesOverride(t *testing.T) {
	charts := testChartsHelper()
	valuesFile := cfg.ChartsValuesDir + "/" + charts.Charts[0].Name + ".yaml"

	if _, err := os.Stat(valuesFile); err == nil {
		t.Errorf("Values override file already exists: %s", valuesFile)
	}

	err := createValuesOverride(charts)
	if err != nil {
		t.Errorf("Expected no error creating Values override, got %s", err.Error())
	}

	if _, err := os.Stat(valuesFile); os.IsNotExist(err) {
		t.Errorf("Expected values file, but got file not found. %s", err.Error())
	}

	t.Cleanup(func() {
		os.RemoveAll(valuesFile)
	})
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

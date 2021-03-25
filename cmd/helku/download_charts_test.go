package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadChartsManifest(t *testing.T) {
	cfg.ChartsManifestFile = "../../chart_manifest.yaml"

	_, err := loadChartsManifest()
	assert := assert.New(t)
	assert.NoError(err, "Expects no error")
}

func testChartsHelper() Charts {
	cfg.ChartsOutputDir = "../../test-charts"
	cfg.ChartsValuesDir = cfg.ChartsOutputDir + "/values"
	cfg.ChartsManifestFile = "../../chart_manifest.yaml"

	return Charts{
		[]Chart{
			{
				Repo: "https://helm.datadoghq.com",
				Name: "datadog",
			},
		},
	}
}

func TestDownloadHelmCharts(t *testing.T) {
	charts := testChartsHelper()

	assert := assert.New(t)
	err := downloadHelmCharts(charts)
	assert.NoError(err, "expects no error")

	chartName := charts.Charts[0].Name
	assert.DirExists(cfg.ChartsOutputDir + "/" + chartName)

	t.Cleanup(func() {
		os.RemoveAll(cfg.ChartsOutputDir)
	})
}

func TestCreateValuesOverride(t *testing.T) {
	charts := testChartsHelper()
	valuesFile := cfg.ChartsValuesDir + "/" + charts.Charts[0].Name + ".yaml"

	assert := assert.New(t)
	assert.NoFileExists(valuesFile, "Expects values file to not exist")

	err := createValuesOverride(charts)
	assert.NoError(err, "expects no error")
	assert.FileExists(valuesFile, "expects values file to exist")

	t.Cleanup(func() {
		os.Remove(valuesFile)
	})
}

func TestRemoveFiles(t *testing.T) {
	dir := "/tmp/test"
	file := dir + "/write-test"

	err := os.MkdirAll(dir, 0755)
	assert := assert.New(t)
	assert.NoError(err, "expects no error")

	bytes := []byte("test\n")
	ioutil.WriteFile(file, bytes, 0644)

	assert.FileExists(file, "Expects file to be created")

	err = removeFiles(dir)
	assert.NoError(err, "Expects no error")

	assert.NoFileExists(file, "Expects file to have been deleted")

	err = removeFiles("/tmp/fake-file-test")
	assert.NoError(err, "Expected to no error")
	assert.NoFileExists("/tmp/fake-file-test", "Expects file to be deleted")
}

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func writeTestYamlManifest(t *testing.T) error {
	charts := Charts{
		[]Chart{
			{
				Name: "kube-state-metrics",
				Repo: "https://kubernetes.github.io/kube-state-metrics",
			},
		},
	}
	b, err := yaml.Marshal(charts)
	if err != nil {
		return err
	}
	err = os.WriteFile("test_manifest.yaml", b, 0644)
	if err != nil {
		return err
	}

	t.Cleanup(func() {
		os.Remove("test_manifest.yaml")
	})

	return nil
}

func TestCliAppDownload(t *testing.T) {
	assert := assert.New(t)
	writeTestYamlManifest(t)
	var (
		expectManifestFile = "test_manifest.yaml"
		expectOutputDir    = "test-charts"
		expectValuesDir    = "test-charts/values"
		expectChart        = "test-charts/kube-state-metrics"
	)

	args := os.Args[0:1]
	args = append(args, "download",
		"-manifest="+expectManifestFile,
		"-destination="+expectOutputDir,
		"-values-dir="+expectValuesDir,
	)

	err := initCliApp(args)
	assert.NoError(err, "Expected no error")

	assert.Equal(expectManifestFile, cfg.ChartsManifestFile, "Manifest files should be the same")

	assert.Equal(expectOutputDir, cfg.ChartsOutputDir, "Expect Charts Output dir to be the same")

	assert.Equal(expectValuesDir, cfg.ChartsValuesDir, "Expect chart values dir to be the same")

	assert.FileExists(expectChart, "Expects chart to exist")

	t.Cleanup(func() {
		os.RemoveAll(expectOutputDir)
		os.RemoveAll(expectValuesDir)
	})
}

func TestCliAppTemplate(t *testing.T) {
	assert := assert.New(t)
	args := os.Args[0:1]
	args = append(args, "template")
	err := initCliApp(args)
	assert.NoError(err, "Expected no error")
}

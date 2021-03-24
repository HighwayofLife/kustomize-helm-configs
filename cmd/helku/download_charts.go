package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
	helm "helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

type Charts struct {
	Charts []Chart `json:"charts"`
}

type Chart struct {
	Name    string `json:"name"`
	Repo    string `json:"repo"`
	Version string `json:"version,omitempty"`
}

func loadChartsManifest() (Charts, error) {
	var charts Charts

	yamlData, err := ioutil.ReadFile(cfg.ChartsManifestFile)
	if err != nil {
		logger.Errorw("Unable to read file",
			"file", cfg.ChartsManifestFile,
			"error", err.Error(),
		)
		return charts, err
	}

	err = yaml.Unmarshal(yamlData, &charts)
	if err != nil {
		logger.Errorw("Unable to unmarshall yaml data",
			"file", cfg.ChartsManifestFile,
			"error", err.Error(),
		)
		return charts, err
	}

	return charts, nil
}

// loop over charts
// run helm pull into output dir
func downloadHelmCharts(charts Charts) error {
	err := os.MkdirAll(cfg.ChartsOutputDir, 0755)
	if err != nil {
		return err
	}

	settings := cli.New()
	client := helm.NewPull()
	client.DestDir = cfg.ChartsOutputDir
	client.Untar = true

	for _, chart := range charts.Charts {
		client.Version = chart.Version
		client.RepoURL = chart.Repo

		removeFiles(cfg.ChartsOutputDir + "/" + chart.Name + "*")

		logger.Infow("Downloading Helm Chart",
			"chart", chart.Name,
			"repo", chart.Repo,
			"version", chart.Version,
		)
		client.Settings = settings

		_, err := client.Run(chart.Name)
		if err != nil {
			logger.Errorw("Error pulling chart",
				"chart", chart.Name,
				"repo", chart.Repo,
				"version", chart.Version,
				"error", err.Error(),
			)
			return err
		}

		removeFiles(cfg.ChartsOutputDir + "/" + chart.Name + "*tgz")
	}

	return nil
}

func createValuesOverride(charts Charts) error {

	os.MkdirAll(cfg.ChartsValuesDir, 0755)
	for _, chart := range charts.Charts {
		valuesFile := cfg.ChartsValuesDir + "/" + chart.Name + ".yaml"
		// Override already exists
		if _, err := os.Stat(valuesFile); err == nil {
			continue
		}

		bytes := []byte("namespace: default\n")
		err := ioutil.WriteFile(valuesFile, bytes, 0644)
		if err != nil {
			logger.Errorw("Error writing values override for chart",
				"values_file", valuesFile,
				"chart", chart.Name,
				"error", err.Error(),
			)
			return err
		}

		logger.Infow(
			"Created values override for chart",
			"values_file", valuesFile,
			"chart", chart.Name,
		)
	}

	return nil
}

func removeFiles(glob string) error {
	files, err := filepath.Glob(glob)
	if err != nil {
		logger.Fatalw("Error finding files",
			"glob", glob,
			"error", err.Error(),
		)
		return err
	}
	for _, f := range files {
		if err := os.RemoveAll(f); err != nil {
			logger.Fatalw("Unable to delete file",
				"glob", glob,
				"file", f,
				"error", err.Error(),
			)
			return err
		}
	}

	return nil
}

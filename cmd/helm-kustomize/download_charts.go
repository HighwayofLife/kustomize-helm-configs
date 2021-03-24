package main

import (
	"fmt"
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
	fmt.Println("load charts manifest func")

	yamlData, err := ioutil.ReadFile(cfg.ChartsManifestFile)
	if err != nil {
		logger.Errorw("Unable to read file",
			"file", cfg.ChartsManifestFile,
			"error", err.Error(),
		)
		return charts, err
	}
	fmt.Println("Yaml Data")

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

func downloadHelmCharts(charts Charts) error {
	// loop over charts
	// run helm pull into output dir
	err := os.MkdirAll(cfg.ChartsOutputDir, 0755)
	if err != nil {
		return err
	}

	settings := cli.New()
	client := helm.NewPull()
	client.DestDir = cfg.ChartsOutputDir
	client.Untar = true

	for _, chart := range charts.Charts {
		if chart.Version != "" {
			client.Version = chart.Version
		}
		client.RepoURL = chart.Repo

		removeFiles(cfg.ChartsOutputDir + "/" + chart.Name + "*")

		logger.Infow("Downloading Helm Chart",
			"chart", chart.Name,
			"repo", chart.Repo,
			"version", chart.Version,
		)
		client.Settings = settings

		output, err := client.Run(chart.Name)
		if err != nil {
			logger.Errorw("Error pulling chart",
				"chart", chart.Name,
				"repo", chart.Repo,
				"version", chart.Version,
				"output", output,
				"error", err.Error(),
			)
			return err
		}

		removeFiles(cfg.ChartsOutputDir + "/" + chart.Name + "*tgz")
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
package main

import "github.com/caarlos0/env"

type config struct {
	ChartsOutputDir      string `env:"HELKU_OUTPUT_DIR" envDefault:"charts"`
	ChartsManifestFile   string `env:"HELKU_CHARTS_MANIFEST_FILE" envDefault:"chart_manifest.yaml"`
	ChartsValuesDir      string `env:"HELKU_CHARTS_VALUES_DIR" envDefault:"charts/values"`
	ClusterChartsBaseDir string `env:"HELKU_CLUSTER_CHARTS_BASE_DIR" envDefault:"cluster/charts"`
}

func (c *config) loadConfigs() {
	err := env.Parse(c)
	if err != nil {
		logger.Errorf("Unable to parse configs")
	}
}

package main

import "github.com/caarlos0/env"

type config struct {
	ChartsOutputDir    string `env:"HELM_KUSTOMIZE_OUTPUT_DIR" envDefault:"charts"`
	ChartsManifestFile string `env:"HELM_KUSTOMIZE_CHARTS_MANIFEST_FILE" envDefault:"chart_manifest.yaml"`
	ChartsValuesDir    string `env:"HELM_KUSTOMIZE_CHARTS_VALUES_DIR" envDefault:"charts/values"`
	ClusterBaseDir     string `env:"HELM_KUSTOMIZE_CLUSER_BASE_DIR" envDefault:"cluster/charts"`
}

func (c *config) loadConfigs() {
	err := env.Parse(c)
	if err != nil {
		logger.Errorf("Unable to parse configs")
	}
}

package main

import "github.com/caarlos0/env"

type config struct {
	ChartsOutputDir    string `env:"HELM_KUSTOMIZE_OUTPUT_DIR" envDefault:"charts"`
	ChartsManifestFile string `env:"HELM_KUSTOMIZE_CHARTS_MANIFEST_FILE" envDefault:"chart_manifest.yaml"`
}

func (c *config) loadConfigs() {
	err := env.Parse(c)
	if err != nil {
		logger.Errorf("Unable to parse configs")
	}
}
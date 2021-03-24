package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/urfave/cli/v2"
)

func initCliApp() error {
	app := &cli.App{
		Name:        "helku",
		Version:     "v0.0.1",
		Description: "Download helm charts and apply Kustomize overlays",
		Commands: []*cli.Command{
			{
				Name:    "download",
				Aliases: []string{"pull"},
				Usage:   "Download helm charts from manifest file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "destination",
						Aliases: []string{"d"},
						Usage:   "`DIR` location to write the chart.",
						Value:   cfg.ChartsOutputDir,
					},
				},
				Action: runDownloadCharts,
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "Run helm template on downloaded charts",
				Action:  runHelmTemplate,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatalw("Failed to run app", "error", err.Error())
	}

	return nil
}

func runDownloadCharts(c *cli.Context) error {
	charts, _ := loadChartsManifest()
	downloadHelmCharts(charts)
	createValuesOverride(charts)

	fmt.Fprintf(c.App.Writer, "Downladed charts to %s\n", cfg.ChartsOutputDir)

	return nil
}

func runHelmTemplate(c *cli.Context) error {
	files, err := ioutil.ReadDir(cfg.ChartsOutputDir)
	if err != nil {
		logger.Errorw("Error reading directory",
			"dir", cfg.ChartsOutputDir,
			"error", err.Error(),
		)
	}

	for _, file := range files {
		if !file.Mode().IsDir() {
			continue
		}

		if file.Name() == "values" {
			continue
		}

		data, err := ioutil.ReadFile(file.Name() + "/Chart.yaml")
		if err != nil {
			logger.Errorw(
				"Unable to read file",
				"file", file.Name()+"/Chart.yaml",
				"error", err.Error(),
			)
		}

		match, err := regexp.Match(`^version: (.+)`, data)
		if err != nil {
			logger.Errorw(
				"Error finding version in Chart.yaml",
				"file", file.Name()+"/Chart.yaml",
				"error", err.Error(),
			)
		}
		fmt.Println(match)
	}

	fmt.Fprintf(c.App.Writer, "Generated manifests in %s from %s", cfg.ClusterBaseDir, cfg.ChartsOutputDir)
	return nil
}

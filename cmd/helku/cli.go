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
		Usage:       "Manage helm charts for local cache and kustomize overlays",
		Description: "Download helm charts and apply Kustomize overlays",
		Commands: []*cli.Command{
			{
				Name:    "download",
				Usage:   "Download (pull) helm charts from manifest file",
				Action:  runDownloadCharts,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "destination",
						Aliases: []string{"d"},
						Usage:   "`DIR` location to write the chart.",
						Value:   cfg.ChartsOutputDir,
					},
				},
			},
			{
				Name:   "template",
				Usage:  "Run helm template on downloaded charts",
				Action: runHelmTemplate,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output-dir",
						Aliases: []string{"o"},
						Usage:   "writes the executed templates to files in output-dir `DIR`",
						Value:   cfg.ClusterBaseDir,
					},
				},
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

		chartVersion(file.Name())
	}

	fmt.Fprintf(c.App.Writer, "Would generate manifests in %s from %s", cfg.ClusterBaseDir, cfg.ChartsOutputDir)
	return nil
}

func chartVersion(chartDir string) error {
	data, err := ioutil.ReadFile(chartDir + "/Chart.yaml")
	if err != nil {
		logger.Errorw(
			"Unable to read file",
			"file", chartDir+"/Chart.yaml",
			"error", err.Error(),
		)
		return err
	}

	match, err := regexp.Match(`^version: (.+)`, data)
	if err != nil {
		logger.Errorw(
			"Error finding version in Chart.yaml",
			"file", chartDir+"/Chart.yaml",
			"error", err.Error(),
		)
		return err
	}
	fmt.Println(match)

	return nil
}

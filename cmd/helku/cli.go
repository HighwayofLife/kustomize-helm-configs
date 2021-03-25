package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/urfave/cli/v2"
)

type ChartYAML struct {
	AppVersion string `json:"appVersion"`
	Name       string `json:"name"`
	Version    string `json:"version"`
}

func initCliApp(args []string) error {
	app := &cli.App{
		Name:        "helku",
		Version:     "v0.0.1",
		Usage:       "Manage helm charts for local cache and kustomize overlays",
		Description: "Download helm charts and apply Kustomize overlays",
		Commands: []*cli.Command{
			{
				Name:   "download",
				Usage:  "Download (pull) helm charts from manifest file",
				Action: runDownloadCharts,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "destination",
						Aliases: []string{"d"},
						Usage:   "`DIR` location to write the chart.",
						Value:   cfg.ChartsOutputDir,
					},
					&cli.StringFlag{
						Name:    "manifest",
						Aliases: []string{"m"},
						Usage:   "YAML Manifest `FILE` containing a list of charts",
						Value:   cfg.ChartsManifestFile,
					},
					&cli.StringFlag{
						Name:  "values-dir",
						Usage: "`DIR` for values overrides.",
						Value: cfg.ChartsValuesDir,
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
						Value:   cfg.ClusterChartsBaseDir,
					},
				},
			},
		},
	}

	err := app.Run(args)
	if err != nil {
		logger.Fatalw("Failed to run app", "error", err.Error())
	}

	return nil
}

func runDownloadCharts(c *cli.Context) error {
	if c.String("manifest") != "" {
		cfg.ChartsManifestFile = c.String("manifest")
	}

	if c.String("destination") != "" {
		cfg.ChartsOutputDir = c.String("destination")
	}

	if c.String("values-dir") != "" {
		cfg.ChartsValuesDir = c.String("values-dir")
	}

	charts, _ := loadChartsManifest()
	downloadHelmCharts(charts)
	createValuesOverride(charts)

	fmt.Fprintf(c.App.Writer, "Downladed charts to %s\n", cfg.ChartsOutputDir)

	return nil
}

func runHelmTemplate(c *cli.Context) error {
	if c.String("output-dir") != "" {
		cfg.ClusterChartsBaseDir = c.String("output-dir")
	}

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

		chartVersion(cfg.ChartsOutputDir + "/" + file.Name())
	}

	fmt.Fprintf(c.App.Writer, "Would generate manifests in %s from %s", cfg.ClusterChartsBaseDir, cfg.ChartsOutputDir)
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
	logger.Info("Looking for a regex string...")

	r, err := regexp.Compile(`(?m)^version: (.+)`)
	if err != nil {
		logger.Errorw(
			"Error compiling regexp",
			"file", chartDir+"/Chart.yaml",
			"error", err.Error(),
		)
		return err
	}
	match := r.FindStringSubmatch(string(data))
	if len(match) > 1 {
		fmt.Println(match[1])
	}

	return nil
}

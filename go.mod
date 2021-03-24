module github.com/highwayoflife/kustomize-helm-configs

go 1.16

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/zap v1.16.0
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools/v3 v3.0.3
	helm.sh/helm/v3 v3.5.3
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)

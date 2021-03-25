# kustomize-helm-configs (helku)

## Description
CLI tool to manage a manifest of helm charts, usually 3rd party helm charts and keeping those charts up to date in a local cache, managing a values override, then applying kustomize overlays for more specific customizations that the helm chart might not support, or to manage configurations across many different clusters and environments.

## Usage
Run `helku` to see the full help and usage.

This package is designed to read a chart manifest file. An example is located in the repo and follows this format:
```yaml
charts:
  - name: <chart-name>
    repo: <chart-repo-url>
    version: <chart-version> ## optional, omit to always pull latest chart
```

```sh
helku download --manifest <manifest-file> ## Download all helm charts in the chart manifest file
helku template --output-dir <cluster/charts> ## Generate kubernetes manifests from downloaded charts
```

## Motivation
After running helm charts in production environments, we found that many 3rd party helm charts don't have sufficient support for customizations that are required in production environments, or the level of control needed. Modifying the helm charts or the generated manifests locally made keeping those upstream charts up to date extremely challenging and consequently, many charts fell way behind the upstream.
This tool follows the paradigm of using Kustomize as a post processor, but with a more specific workflow that is needed for GitOps tools like ArgoCD or FluxCD without using a Helm controller/operator. This allows the helm charts to be kept in local caches which benefit from normal Pull-requests, diffs, and reviews. So making modifications to the helm values file becomes a more straightforward and easier process, and ultimately, all it takes to upgrade the chart while keeping modifications in place is to run the `helku download` command.

## Contributions Guidelines

* If you find a bug or have a feature request, submit an issue.
* File an issue first prior to submitting a PR!
* Ensure all exported items are properly commented
* If applicable, submit a test suite against your PR

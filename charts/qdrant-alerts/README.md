# Qdrant alerts

[Qdrant alerts by AppsCode](https://github.com/appscode/alerts) - Qdrant alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/qdrant-alerts --version=v2025.6.30
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2025.6.30
```

## Introduction

This chart deploys Qdrant alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `qdrant`:

```bash
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2025.6.30
```

The command deploys Qdrant alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `qdrant`:

```bash
$ helm uninstall qdrant -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `qdrant-alerts` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Qdrant</code>                              |
| metadata.resource.name                                                        |                                               | <code>qdrants</code>                             |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>"qdrant-quickstart"</code>                 |
| metadata.release.namespace                                                    | Release namespace                             | <code>"demo"</code>                              |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantDown.enabled                           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantDown.duration                          |                                               | <code>"30s"</code>                               |
| form.alert.groups.database.rules.qdrantDown.severity                          |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.qdrantPhaseCritical.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantPhaseCritical.duration                 |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantPhaseCritical.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantClusterNoBaselineNode.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantClusterNoBaselineNode.duration         |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantClusterNoBaselineNode.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantClusterNoBaselineNode.val              |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.qdrantRestarted.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantRestarted.duration                     |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantRestarted.severity                     |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantRestarted.val                          |                                               | <code>180</code>                                 |
| form.alert.groups.database.rules.qdrantHighCPULoad.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantHighCPULoad.duration                   |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantHighCPULoad.severity                   |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantHighCPULoad.val                        |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.qdrantHighHeapMemoryUsed.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantHighHeapMemoryUsed.duration            |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantHighHeapMemoryUsed.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantHighHeapMemoryUsed.val                 |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.qdrantHighDataregionOffHeapUsed.enabled      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantHighDataregionOffHeapUsed.duration     |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantHighDataregionOffHeapUsed.severity     |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantHighDataregionOffHeapUsed.val          |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.qdrantJVMPausesTotalDuration.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.qdrantJVMPausesTotalDuration.duration        |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.qdrantJVMPausesTotalDuration.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.qdrantJVMPausesTotalDuration.val             |                                               | <code>0</code>                                   |
| form.alert.groups.database.rules.diskUsageHigh.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.val                            |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.diskUsageHigh.duration                       |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskUsageHigh.severity                       |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.diskAlmostFull.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.val                           |                                               | <code>95</code>                                  |
| form.alert.groups.database.rules.diskAlmostFull.duration                      |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.diskAlmostFull.severity                      |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.enabled                                         |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.enabled                                          |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                               |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                            |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                            |
| grafana.enabled                                                               |                                               | <code>true</code>                                |
| grafana.version                                                               |                                               | <code>7.5.5</code>                               |
| grafana.jobName                                                               |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                   |                                               | <code>""</code>                                  |
| grafana.apikey                                                                |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2025.6.30 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i qdrant appscode/qdrant-alerts -n demo --create-namespace --version=v2025.6.30 --values values.yaml
```

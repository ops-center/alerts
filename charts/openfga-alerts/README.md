# OpenFGA alerts

[OpenFGA alerts by AppsCode](https://github.com/appscode/alerts) - OpenFGA alerts for monitoring

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/openfga-alerts --version=v2025.6.30
$ helm upgrade -i openfga appscode/openfga-alerts -n demo --create-namespace --version=v2025.6.30
```

## Introduction

This chart deploys OpenFGA alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `openfga`:

```bash
$ helm upgrade -i openfga appscode/openfga-alerts -n demo --create-namespace --version=v2025.6.30
```

The command deploys OpenFGA alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `openfga`:

```bash
$ helm uninstall openfga -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `openfga-alerts` chart and their default values.

|                            Parameter                            | Description |              Default               |
|-----------------------------------------------------------------|-------------|------------------------------------|
| metadata.resource.group                                         |             | <code>alerts.appscode.com</code>   |
| metadata.resource.kind                                          |             | <code>OpenFGAAlerts</code>         |
| metadata.resource.name                                          |             | <code>openfgaalerts</code>         |
| metadata.resource.scope                                         |             | <code>Namespaced</code>            |
| metadata.resource.version                                       |             | <code>v1alpha1</code>              |
| metadata.release.name                                           |             | <code>""</code>                    |
| metadata.release.namespace                                      |             | <code>""</code>                    |
| form.alert.enabled                                              |             | <code>warning</code>               |
| form.alert.labels.release                                       |             | <code>kube-prometheus-stack</code> |
| form.alert.annotations                                          |             | <code>{}</code>                    |
| form.alert.additionalRuleLabels                                 |             | <code>{}</code>                    |
| form.alert.groups.database.enabled                              |             | <code>warning</code>               |
| form.alert.groups.database.rules.highRequestLatency.enabled     |             | <code>true</code>                  |
| form.alert.groups.database.rules.highRequestLatency.duration    |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.highRequestLatency.val         |             | <code>1.0</code>                   |
| form.alert.groups.database.rules.highRequestLatency.severity    |             | <code>critical</code>              |
| form.alert.groups.database.rules.highErrorRate.enabled          |             | <code>true</code>                  |
| form.alert.groups.database.rules.highErrorRate.duration         |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.highErrorRate.val              |             | <code>5</code>                     |
| form.alert.groups.database.rules.highErrorRate.severity         |             | <code>critical</code>              |
| form.alert.groups.database.rules.lowCheckCacheHitRatio.enabled  |             | <code>true</code>                  |
| form.alert.groups.database.rules.lowCheckCacheHitRatio.duration |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.lowCheckCacheHitRatio.val      |             | <code>70</code>                    |
| form.alert.groups.database.rules.lowCheckCacheHitRatio.severity |             | <code>warning</code>               |
| form.alert.groups.database.rules.highSQLConnections.enabled     |             | <code>true</code>                  |
| form.alert.groups.database.rules.highSQLConnections.duration    |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.highSQLConnections.val         |             | <code>80</code>                    |
| form.alert.groups.database.rules.highSQLConnections.severity    |             | <code>critical</code>              |
| form.alert.groups.database.rules.highGoroutines.enabled         |             | <code>true</code>                  |
| form.alert.groups.database.rules.highGoroutines.duration        |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.highGoroutines.val             |             | <code>10000</code>                 |
| form.alert.groups.database.rules.highGoroutines.severity        |             | <code>warning</code>               |
| form.alert.groups.database.rules.highGCDuration.enabled         |             | <code>true</code>                  |
| form.alert.groups.database.rules.highGCDuration.duration        |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.highGCDuration.val             |             | <code>1.0</code>                   |
| form.alert.groups.database.rules.highGCDuration.severity        |             | <code>warning</code>               |
| form.alert.groups.database.rules.highHeapMemory.enabled         |             | <code>true</code>                  |
| form.alert.groups.database.rules.highHeapMemory.duration        |             | <code>"5m"</code>                  |
| form.alert.groups.database.rules.highHeapMemory.val             |             | <code>4294967296 # 4GB</code>      |
| form.alert.groups.database.rules.highHeapMemory.severity        |             | <code>warning</code>               |
| form.alert.groups.database.rules.highAllocationRate.enabled     |             | <code>true</code>                  |
| form.alert.groups.database.rules.highAllocationRate.duration    |             | <code>"2m"</code>                  |
| form.alert.groups.database.rules.highAllocationRate.val         |             | <code>104857600 # 100MB/s</code>   |
| form.alert.groups.database.rules.highAllocationRate.severity    |             | <code>warning</code>               |
| grafana.enabled                                                 |             | <code>false</code>                 |
| grafana.version                                                 |             | <code>8.2.3</code>                 |
| grafana.jobName                                                 |             | <code>openfga</code>               |
| grafana.url                                                     |             | <code>""</code>                    |
| grafana.apikey                                                  |             | <code>""</code>                    |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i openfga appscode/openfga-alerts -n demo --create-namespace --version=v2025.6.30 --set metadata.resource.group=alerts.appscode.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i openfga appscode/openfga-alerts -n demo --create-namespace --version=v2025.6.30 --values values.yaml
```

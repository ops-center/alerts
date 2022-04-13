# MongoDB alerts

[MongoDB alerts by AppsCode](https://github.com/appscode/alerts) - MongoDB alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/mongodb --version=v0.1.0
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys MongoDB alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `mongodb`:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0
```

The command deploys MongoDB alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `mongodb`:

```bash
$ helm uninstall mongodb -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `mongodb` chart and their default values.

|                         Parameter                          |                  Description                  |                     Default                      |
|------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                    |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                     |                                               | <code>MongoDB</code>                             |
| metadata.resource.name                                     |                                               | <code>mongodbs</code>                            |
| metadata.resource.scope                                    |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                  |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                      | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                 | Release namespace                             | <code>""</code>                                  |
| spec.alert.enabled                                         | # Enable PrometheusRule alerts                | <code>true</code>                                |
| spec.alert.labels                                          | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| spec.alert.annotations                                     | # Annotations for default rules               | <code>{}</code>                                  |
| spec.alert.additionalRuleLabels                            | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| spec.alert.rules.mongodbVirtualMemoryUsage.enabled         |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbVirtualMemoryUsage.val             |                                               | <code>2097152 # 2GB</code>                       |
| spec.alert.rules.mongodbVirtualMemoryUsage.duration        |                                               | <code>"1m"</code>                                |
| spec.alert.rules.mongodbVirtualMemoryUsage.severity        |                                               | <code>warning</code>                             |
| spec.alert.rules.mongodbReplicationLag.enabled             |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbReplicationLag.val                 |                                               | <code>10</code>                                  |
| spec.alert.rules.mongodbReplicationLag.duration            |                                               | <code>"0m"</code>                                |
| spec.alert.rules.mongodbReplicationLag.severity            |                                               | <code>crititcal</code>                           |
| spec.alert.rules.mongodbNumberCursorsOpen.enabled          |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbNumberCursorsOpen.val              |                                               | <code>10000</code>                               |
| spec.alert.rules.mongodbNumberCursorsOpen.duration         |                                               | <code>"2m"</code>                                |
| spec.alert.rules.mongodbNumberCursorsOpen.severity         |                                               | <code>warning</code>                             |
| spec.alert.rules.mongodbCursorsTimeouts.enabled            |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbCursorsTimeouts.val                |                                               | <code>100</code>                                 |
| spec.alert.rules.mongodbCursorsTimeouts.duration           |                                               | <code>"2m"</code>                                |
| spec.alert.rules.mongodbCursorsTimeouts.severity           |                                               | <code>warning</code>                             |
| spec.alert.rules.mongodbTooManyConnections.enabled         |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbTooManyConnections.val             |                                               | <code>80 # percentage</code>                     |
| spec.alert.rules.mongodbTooManyConnections.duration        |                                               | <code>"2m"</code>                                |
| spec.alert.rules.mongodbTooManyConnections.severity        |                                               | <code>warning</code>                             |
| spec.alert.rules.mongoDBPhaseCritical.enabled              |                                               | <code>true</code>                                |
| spec.alert.rules.mongoDBPhaseCritical.duration             |                                               | <code>"3m"</code>                                |
| spec.alert.rules.mongoDBPhaseCritical.severity             |                                               | <code>warning</code>                             |
| spec.alert.rules.mongoDBDown.enabled                       |                                               | <code>true</code>                                |
| spec.alert.rules.mongoDBDown.duration                      |                                               | <code>"30s"</code>                               |
| spec.alert.rules.mongoDBDown.severity                      |                                               | <code>crititcal</code>                           |
| spec.alert.rules.mongodbHighLatency.enabled                |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbHighLatency.val                    |                                               | <code>250000</code>                              |
| spec.alert.rules.mongodbHighLatency.duration               |                                               | <code>"10m"</code>                               |
| spec.alert.rules.mongodbHighLatency.severity               |                                               | <code>warning</code>                             |
| spec.alert.rules.mongodbHighTicketUtilization.enabled      |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbHighTicketUtilization.val          |                                               | <code>75 # percentage</code>                     |
| spec.alert.rules.mongodbHighTicketUtilization.duration     |                                               | <code>"10m"</code>                               |
| spec.alert.rules.mongodbHighTicketUtilization.severity     |                                               | <code>warning</code>                             |
| spec.alert.rules.mongodbRecurrentCursorTimeout.enabled     |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbRecurrentCursorTimeout.val         |                                               | <code>0</code>                                   |
| spec.alert.rules.mongodbRecurrentCursorTimeout.duration    |                                               | <code>"30m"</code>                               |
| spec.alert.rules.mongodbRecurrentCursorTimeout.severity    |                                               | <code>warning</code>                             |
| spec.alert.rules.mongodbRecurrentMemoryPageFaults.enabled  |                                               | <code>true</code>                                |
| spec.alert.rules.mongodbRecurrentMemoryPageFaults.val      |                                               | <code>0</code>                                   |
| spec.alert.rules.mongodbRecurrentMemoryPageFaults.duration |                                               | <code>"30m"</code>                               |
| spec.alert.rules.mongodbRecurrentMemoryPageFaults.severity |                                               | <code>warning</code>                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

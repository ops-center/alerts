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

|                       Parameter                       | Description |              Default               |
|-------------------------------------------------------|-------------|------------------------------------|
| namespace                                             |             | <code>demo</code>                  |
| dbName                                                |             | <code>mg-rs</code>                 |
| alert.ruleSelector.app                                |             | <code>kube-prometheus-stack</code> |
| alert.ruleSelector.release                            |             | <code>prometheus</code>            |
| alert.rules.mongodbVirtualMemoryUsage.enabled         |             | <code>true</code>                  |
| alert.rules.mongodbVirtualMemoryUsage.val             |             | <code>2097152 # 2GB</code>         |
| alert.rules.mongodbVirtualMemoryUsage.duration        |             | <code>"1m"</code>                  |
| alert.rules.mongodbVirtualMemoryUsage.severity        |             | <code>warning</code>               |
| alert.rules.mongodbReplicationLag.enabled             |             | <code>true</code>                  |
| alert.rules.mongodbReplicationLag.val                 |             | <code>10</code>                    |
| alert.rules.mongodbReplicationLag.duration            |             | <code>"0m"</code>                  |
| alert.rules.mongodbReplicationLag.severity            |             | <code>crititcal</code>             |
| alert.rules.mongodbNumberCursorsOpen.enabled          |             | <code>true</code>                  |
| alert.rules.mongodbNumberCursorsOpen.val              |             | <code>10000</code>                 |
| alert.rules.mongodbNumberCursorsOpen.duration         |             | <code>"2m"</code>                  |
| alert.rules.mongodbNumberCursorsOpen.severity         |             | <code>warning</code>               |
| alert.rules.mongodbCursorsTimeouts.enabled            |             | <code>true</code>                  |
| alert.rules.mongodbCursorsTimeouts.val                |             | <code>100</code>                   |
| alert.rules.mongodbCursorsTimeouts.duration           |             | <code>"2m"</code>                  |
| alert.rules.mongodbCursorsTimeouts.severity           |             | <code>warning</code>               |
| alert.rules.mongodbTooManyConnections.enabled         |             | <code>true</code>                  |
| alert.rules.mongodbTooManyConnections.val             |             | <code>80 # percentage</code>       |
| alert.rules.mongodbTooManyConnections.duration        |             | <code>"2m"</code>                  |
| alert.rules.mongodbTooManyConnections.severity        |             | <code>warning</code>               |
| alert.rules.mongoDBPhaseCritical.enabled              |             | <code>true</code>                  |
| alert.rules.mongoDBPhaseCritical.duration             |             | <code>"3m"</code>                  |
| alert.rules.mongoDBPhaseCritical.severity             |             | <code>warning</code>               |
| alert.rules.mongoDBDown.enabled                       |             | <code>true</code>                  |
| alert.rules.mongoDBDown.duration                      |             | <code>"30s"</code>                 |
| alert.rules.mongoDBDown.severity                      |             | <code>crititcal</code>             |
| alert.rules.mongodbHighLatency.enabled                |             | <code>true</code>                  |
| alert.rules.mongodbHighLatency.val                    |             | <code>250000</code>                |
| alert.rules.mongodbHighLatency.duration               |             | <code>"10m"</code>                 |
| alert.rules.mongodbHighLatency.severity               |             | <code>warning</code>               |
| alert.rules.mongodbHighTicketUtilization.enabled      |             | <code>true</code>                  |
| alert.rules.mongodbHighTicketUtilization.val          |             | <code>75 # percentage</code>       |
| alert.rules.mongodbHighTicketUtilization.duration     |             | <code>"10m"</code>                 |
| alert.rules.mongodbHighTicketUtilization.severity     |             | <code>warning</code>               |
| alert.rules.mongodbRecurrentCursorTimeout.enabled     |             | <code>true</code>                  |
| alert.rules.mongodbRecurrentCursorTimeout.val         |             | <code>0</code>                     |
| alert.rules.mongodbRecurrentCursorTimeout.duration    |             | <code>"30m"</code>                 |
| alert.rules.mongodbRecurrentCursorTimeout.severity    |             | <code>warning</code>               |
| alert.rules.mongodbRecurrentMemoryPageFaults.enabled  |             | <code>true</code>                  |
| alert.rules.mongodbRecurrentMemoryPageFaults.val      |             | <code>0</code>                     |
| alert.rules.mongodbRecurrentMemoryPageFaults.duration |             | <code>"30m"</code>                 |
| alert.rules.mongodbRecurrentMemoryPageFaults.severity |             | <code>warning</code>               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0 --set namespace=demo
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

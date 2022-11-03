# ProxySQL alerts

[ProxySQL alerts by AppsCode](https://github.com/appscode/alerts) - ProxySQL alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/proxysql --version=v0.1.0
$ helm upgrade -i proxysql appscode/proxysql -n demo --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys ProxySQL alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `proxysql`:

```bash
$ helm upgrade -i proxysql appscode/proxysql -n demo --create-namespace --version=v0.1.0
```

The command deploys ProxySQL alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `proxysql`:

```bash
$ helm uninstall proxysql -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `proxysql` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>ProxySQL</code>                             |
| metadata.resource.name                                                        |                                               | <code>proxysqls</code>                            |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>true</code>                                |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLInstanceDown.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLInstanceDown.duration                |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.proxySQLInstanceDown.severity                |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.proxySQLServiceDown.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLServiceDown.duration                 |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.proxySQLServiceDown.severity                 |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.proxySQLTooManyConnections.enabled           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLTooManyConnections.duration          |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.proxySQLTooManyConnections.val               |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.proxySQLTooManyConnections.severity          |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.proxySQLHighThreadsRunning.enabled           |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLHighThreadsRunning.duration          |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.proxySQLHighThreadsRunning.val               |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.proxySQLHighThreadsRunning.severity          |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.proxySQLSlowQueries.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLSlowQueries.duration                 |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.proxySQLSlowQueries.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.proxySQLRestarted.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLRestarted.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.proxySQLRestarted.val                        |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.proxySQLRestarted.severity                   |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.proxySQLHighQPS.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLHighQPS.duration                     |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.proxySQLHighQPS.val                          |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.proxySQLHighQPS.severity                     |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.proxySQLHighIncomingBytes.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLHighIncomingBytes.duration           |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.proxySQLHighIncomingBytes.val                |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.proxySQLHighIncomingBytes.severity           |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.proxySQLHighOutgoingBytes.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.proxySQLHighOutgoingBytes.duration           |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.proxySQLHighOutgoingBytes.val                |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.proxySQLHighOutgoingBytes.severity           |                                               | <code>critical</code>                            |
| form.alert.groups.cluster.enabled                                             |                                               | <code>true</code>                                |
| form.alert.groups.cluster.rules.proxysqlCLusterSyncFailure.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.cluster.rules.proxysqlCLusterSyncFailure.val                |                                               | <code>0.1</code>                                 |
| form.alert.groups.cluster.rules.proxysqlCLusterSyncFailure.duration           |                                               | <code>"5m"</code>                                |
| form.alert.groups.cluster.rules.proxysqlCLusterSyncFailure.severity           |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.enabled                                         |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.enabled                                          |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                               |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                            |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i proxysql appscode/proxysql -n demo --create-namespace --version=v0.1.0 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i proxysql appscode/proxysql -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

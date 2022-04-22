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

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>MongoDB</code>                             |
| metadata.resource.name                                                        |                                               | <code>mongodbs</code>                            |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| spec.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>true</code>                                |
| spec.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| spec.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| spec.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| spec.alert.groups.database.enabled                                            |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbVirtualMemoryUsage.enabled            |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbVirtualMemoryUsage.val                |                                               | <code>2097152 # 2GB</code>                       |
| spec.alert.groups.database.rules.mongodbVirtualMemoryUsage.duration           |                                               | <code>"1m"</code>                                |
| spec.alert.groups.database.rules.mongodbVirtualMemoryUsage.severity           |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongodbReplicationLag.enabled                |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbReplicationLag.val                    |                                               | <code>10</code>                                  |
| spec.alert.groups.database.rules.mongodbReplicationLag.duration               |                                               | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.mongodbReplicationLag.severity               |                                               | <code>crititcal</code>                           |
| spec.alert.groups.database.rules.mongodbNumberCursorsOpen.enabled             |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbNumberCursorsOpen.val                 |                                               | <code>10000</code>                               |
| spec.alert.groups.database.rules.mongodbNumberCursorsOpen.duration            |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.mongodbNumberCursorsOpen.severity            |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongodbCursorsTimeouts.enabled               |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbCursorsTimeouts.val                   |                                               | <code>100</code>                                 |
| spec.alert.groups.database.rules.mongodbCursorsTimeouts.duration              |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.mongodbCursorsTimeouts.severity              |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongodbTooManyConnections.enabled            |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbTooManyConnections.val                |                                               | <code>80 # percentage</code>                     |
| spec.alert.groups.database.rules.mongodbTooManyConnections.duration           |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.mongodbTooManyConnections.severity           |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongoDBPhaseCritical.enabled                 |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongoDBPhaseCritical.duration                |                                               | <code>"3m"</code>                                |
| spec.alert.groups.database.rules.mongoDBPhaseCritical.severity                |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongoDBDown.enabled                          |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongoDBDown.duration                         |                                               | <code>"30s"</code>                               |
| spec.alert.groups.database.rules.mongoDBDown.severity                         |                                               | <code>crititcal</code>                           |
| spec.alert.groups.database.rules.mongodbHighLatency.enabled                   |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbHighLatency.val                       |                                               | <code>250000</code>                              |
| spec.alert.groups.database.rules.mongodbHighLatency.duration                  |                                               | <code>"10m"</code>                               |
| spec.alert.groups.database.rules.mongodbHighLatency.severity                  |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongodbHighTicketUtilization.enabled         |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbHighTicketUtilization.val             |                                               | <code>75 # percentage</code>                     |
| spec.alert.groups.database.rules.mongodbHighTicketUtilization.duration        |                                               | <code>"10m"</code>                               |
| spec.alert.groups.database.rules.mongodbHighTicketUtilization.severity        |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongodbRecurrentCursorTimeout.enabled        |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbRecurrentCursorTimeout.val            |                                               | <code>0</code>                                   |
| spec.alert.groups.database.rules.mongodbRecurrentCursorTimeout.duration       |                                               | <code>"30m"</code>                               |
| spec.alert.groups.database.rules.mongodbRecurrentCursorTimeout.severity       |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.mongodbRecurrentMemoryPageFaults.enabled     |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.mongodbRecurrentMemoryPageFaults.val         |                                               | <code>0</code>                                   |
| spec.alert.groups.database.rules.mongodbRecurrentMemoryPageFaults.duration    |                                               | <code>"30m"</code>                               |
| spec.alert.groups.database.rules.mongodbRecurrentMemoryPageFaults.severity    |                                               | <code>warning</code>                             |
| spec.alert.groups.provisioner.enabled                                         |                                               | <code>true</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                               | <code>true</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                               | <code>"1m"</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                               | <code>critical</code>                            |
| spec.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                               | <code>true</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                               | <code>"15m"</code>                               |
| spec.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                               | <code>warning</code>                             |
| spec.alert.groups.opsManager.enabled                                          |                                               | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                               | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                               | <code>"0m"</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                               | <code>info</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                               | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                               | <code>"30m"</code>                               |
| spec.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                               | <code>critical</code>                            |
| spec.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                               | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                               | <code>"0m"</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                               | <code>critical</code>                            |
| spec.alert.groups.stash.enabled                                               |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.backupSessionFailed.enabled                     |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.backupSessionFailed.duration                    |                                               | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.backupSessionFailed.severity                    |                                               | <code>critical</code>                            |
| spec.alert.groups.stash.rules.restoreSessionFailed.enabled                    |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.restoreSessionFailed.duration                   |                                               | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.restoreSessionFailed.severity                   |                                               | <code>critical</code>                            |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.enabled               |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.duration              |                                               | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.val                   |                                               | <code>18000</code>                               |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.severity              |                                               | <code>warning</code>                             |
| spec.alert.groups.stash.rules.repositoryCorrupted.enabled                     |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.repositoryCorrupted.duration                    |                                               | <code>"5m"</code>                                |
| spec.alert.groups.stash.rules.repositoryCorrupted.severity                    |                                               | <code>critical</code>                            |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.enabled             |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.duration            |                                               | <code>"5m"</code>                                |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.val                 |                                               | <code>10737418240 # 10GB</code>                  |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.severity            |                                               | <code>waring</code>                              |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled              |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.duration             |                                               | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.val                  |                                               | <code>1800 # 30 minute</code>                    |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.severity             |                                               | <code>waring</code>                              |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled             |                                               | <code>true</code>                                |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration            |                                               | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                 |                                               | <code>1800 # 30 minute</code>                    |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity            |                                               | <code>waring</code>                              |
| spec.alert.groups.schemaManager.enabled                                       |                                               | <code>true</code>                                |
| spec.alert.groups.schemaManager.rules.schemaPendingForTooLong.enabled         |                                               | <code>true</code>                                |
| spec.alert.groups.schemaManager.rules.schemaPendingForTooLong.duration        |                                               | <code>"30m"</code>                               |
| spec.alert.groups.schemaManager.rules.schemaPendingForTooLong.severity        |                                               | <code>warning</code>                             |
| spec.alert.groups.schemaManager.rules.schemaInProgressForTooLong.enabled      |                                               | <code>true</code>                                |
| spec.alert.groups.schemaManager.rules.schemaInProgressForTooLong.duration     |                                               | <code>"30m"</code>                               |
| spec.alert.groups.schemaManager.rules.schemaInProgressForTooLong.severity     |                                               | <code>warning</code>                             |
| spec.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.enabled     |                                               | <code>true</code>                                |
| spec.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.duration    |                                               | <code>"30m"</code>                               |
| spec.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.severity    |                                               | <code>warning</code>                             |
| spec.alert.groups.schemaManager.rules.schemaFailed.enabled                    |                                               | <code>true</code>                                |
| spec.alert.groups.schemaManager.rules.schemaFailed.duration                   |                                               | <code>"0m"</code>                                |
| spec.alert.groups.schemaManager.rules.schemaFailed.severity                   |                                               | <code>warning</code>                             |
| spec.alert.groups.schemaManager.rules.schemaExpired.enabled                   |                                               | <code>true</code>                                |
| spec.alert.groups.schemaManager.rules.schemaExpired.duration                  |                                               | <code>"0m"</code>                                |
| spec.alert.groups.schemaManager.rules.schemaExpired.severity                  |                                               | <code>warning</code>                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mongodb appscode/mongodb -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

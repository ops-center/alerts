# Redis alerts

[Redis alerts by AppsCode](https://github.com/appscode/alerts) - Redis alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/redis --version=v0.1.0
$ helm upgrade -i redis appscode/redis -n demo --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys Redis alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `redis`:

```bash
$ helm upgrade -i redis appscode/redis -n demo --create-namespace --version=v0.1.0
```

The command deploys Redis alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `redis`:

```bash
$ helm uninstall redis -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `redis` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>Redis</code>                               |
| metadata.resource.name                                                        |                                               | <code>redises</code>                             |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| spec.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>true</code>                                |
| spec.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| spec.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| spec.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| spec.alert.groups.database.enabled                                            |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisInstanceDown.enabled                    |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisInstanceDown.duration                   |                                               | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.redisInstanceDown.severity                   |                                               | <code>critical</code>                            |
| spec.alert.groups.database.rules.redisRestarted.enabled                       |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisRestarted.duration                      |                                               | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.redisRestarted.severity                      |                                               | <code>critical</code>                            |
| spec.alert.groups.database.rules.redisTooManyConnections.enabled              |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisTooManyConnections.duration             |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.redisTooManyConnections.val                  |                                               | <code>80</code>                                  |
| spec.alert.groups.database.rules.redisTooManyConnections.severity             |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.redisqlNotEnoughConnections.enabled          |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisqlNotEnoughConnections.duration         |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.redisqlNotEnoughConnections.val              |                                               | <code>5</code>                                   |
| spec.alert.groups.database.rules.redisqlNotEnoughConnections.severity         |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.redisSlowQueries.enabled                     |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisSlowQueries.duration                    |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.redisSlowQueries.severity                    |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.redisqlReplicationLag.enabled                |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisqlReplicationLag.duration               |                                               | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.redisqlReplicationLag.val                    |                                               | <code>30s</code>                                 |
| spec.alert.groups.database.rules.redisqlReplicationLag.severity               |                                               | <code>critical</code>                            |
| spec.alert.groups.database.rules.redisqlHighRollbackRate.enabled              |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisqlHighRollbackRate.duration             |                                               | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.redisqlHighRollbackRate.val                  |                                               | <code>0.02</code>                                |
| spec.alert.groups.database.rules.redisqlHighRollbackRate.severity             |                                               | <code>warning</code>                             |
| spec.alert.groups.database.rules.redisqlSplitBrain.enabled                    |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisqlSplitBrain.duration                   |                                               | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.redisqlSplitBrain.severity                   |                                               | <code>critical</code>                            |
| spec.alert.groups.database.rules.redisqlTooManyLocksAcquired.enabled          |                                               | <code>true</code>                                |
| spec.alert.groups.database.rules.redisqlTooManyLocksAcquired.duration         |                                               | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.redisqlTooManyLocksAcquired.val              |                                               | <code>0.20</code>                                |
| spec.alert.groups.database.rules.redisqlTooManyLocksAcquired.severity         |                                               | <code>critical</code>                            |
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


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i redis appscode/redis -n demo --create-namespace --version=v0.1.0 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i redis appscode/redis -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

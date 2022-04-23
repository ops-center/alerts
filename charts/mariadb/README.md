# MariaDB alerts

[MariaDB alerts by AppsCode](https://github.com/appscode/alerts) - MariaDB alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/mariadb --version=v0.1.0
$ helm upgrade -i mariadb appscode/mariadb -n demo --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys MariaDB alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `mariadb`:

```bash
$ helm upgrade -i mariadb appscode/mariadb -n demo --create-namespace --version=v0.1.0
```

The command deploys MariaDB alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `mariadb`:

```bash
$ helm uninstall mariadb -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `mariadb` chart and their default values.

|                                   Parameter                                   |                  Description                  |                     Default                      |
|-------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                               | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                               | <code>MariaDB</code>                             |
| metadata.resource.name                                                        |                                               | <code>mariadbs</code>                            |
| metadata.resource.scope                                                       |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                            | # Enable PrometheusRule alerts                | <code>true</code>                                |
| form.alert.labels                                                             | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                        | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLInstanceDown.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLInstanceDown.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLInstanceDown.severity                   |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mySQLServiceDown.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLServiceDown.duration                    |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLServiceDown.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mySQLTooManyConnections.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLTooManyConnections.duration             |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mySQLTooManyConnections.val                  |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.mySQLTooManyConnections.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mySQLHighThreadsRunning.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLHighThreadsRunning.duration             |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mySQLHighThreadsRunning.val                  |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.mySQLHighThreadsRunning.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mySQLSlowQueries.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLSlowQueries.duration                    |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mySQLSlowQueries.severity                    |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mySQLInnoDBLogWaits.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLInnoDBLogWaits.duration                 |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLInnoDBLogWaits.val                      |                                               | <code>10</code>                                  |
| form.alert.groups.database.rules.mySQLInnoDBLogWaits.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mySQLRestarted.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLRestarted.duration                      |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLRestarted.val                           |                                               | <code>60</code>                                  |
| form.alert.groups.database.rules.mySQLRestarted.severity                      |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.mySQLHighQPS.enabled                         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLHighQPS.duration                        |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLHighQPS.val                             |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.mySQLHighQPS.severity                        |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mySQLHighIncomingBytes.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLHighIncomingBytes.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLHighIncomingBytes.val                   |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.mySQLHighIncomingBytes.severity              |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mySQLHighOutgoingBytes.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLHighOutgoingBytes.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.mySQLHighOutgoingBytes.val                   |                                               | <code>1048576 # 1MB</code>                       |
| form.alert.groups.database.rules.mySQLHighOutgoingBytes.severity              |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.mySQLTooManyOpenFiles.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.mySQLTooManyOpenFiles.duration               |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.mySQLTooManyOpenFiles.val                    |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.mySQLTooManyOpenFiles.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.cluster.enabled                                             |                                               | <code>true</code>                                |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.enabled       |                                               | <code>true</code>                                |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.val           |                                               | <code>0.1</code>                                 |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.duration      |                                               | <code>"5m"</code>                                |
| form.alert.groups.cluster.rules.galeraReplicationLatencyTooLong.severity      |                                               | <code>warning</code>                             |
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
| form.alert.groups.stash.enabled                                               |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.duration                    |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.restoreSessionFailed.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.severity                   |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.val                   |                                               | <code>18000</code>                               |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.repositoryCorrupted.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.duration                    |                                               | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.duration            |                                               | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.val                 |                                               | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.severity            |                                               | <code>waring</code>                              |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.duration             |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.val                  |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.severity             |                                               | <code>waring</code>                              |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration            |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                 |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity            |                                               | <code>waring</code>                              |
| form.alert.groups.schemaManager.enabled                                       |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.duration        |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.enabled      |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.duration     |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.severity     |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.enabled     |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.duration    |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.severity    |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaFailed.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaFailed.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.schemaManager.rules.schemaFailed.severity                   |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaExpired.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaExpired.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.schemaManager.rules.schemaExpired.severity                  |                                               | <code>warning</code>                             |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mariadb appscode/mariadb -n demo --create-namespace --version=v0.1.0 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mariadb appscode/mariadb -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

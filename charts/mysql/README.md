# MySQL alerts

[MySQL alerts by AppsCode](https://github.com/appscode/alerts) - MySQL alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/mysql --version=v0.1.0
$ helm upgrade -i mysql appscode/mysql -n demo --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys MySQL alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `mysql`:

```bash
$ helm upgrade -i mysql appscode/mysql -n demo --create-namespace --version=v0.1.0
```

The command deploys MySQL alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `mysql`:

```bash
$ helm uninstall mysql -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `mysql` chart and their default values.

|                                      Parameter                                      | Description |              Default               |
|-------------------------------------------------------------------------------------|-------------|------------------------------------|
| namespace                                                                           |             | <code>demo</code>                  |
| dbName                                                                              |             | <code>my-group</code>              |
| alert.ruleSelector.app                                                              |             | <code>kube-prometheus-stack</code> |
| alert.ruleSelector.release                                                          |             | <code>prometheus</code>            |
| alert.groups.database.enabled                                                       |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLInstanceDown.enabled                               |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLInstanceDown.duration                              |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLInstanceDown.severity                              |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLServiceDown.enabled                                |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLServiceDown.duration                               |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLServiceDown.severity                               |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLTooManyConnections.enabled                         |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLTooManyConnections.duration                        |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLTooManyConnections.val                             |             | <code>80</code>                    |
| alert.groups.database.rules.mySQLTooManyConnections.severity                        |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLHighThreadsRunning.enabled                         |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighThreadsRunning.duration                        |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLHighThreadsRunning.val                             |             | <code>60</code>                    |
| alert.groups.database.rules.mySQLHighThreadsRunning.severity                        |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLSlowQueries.enabled                                |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLSlowQueries.duration                               |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLSlowQueries.severity                               |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLInnoDBLogWaits.enabled                             |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLInnoDBLogWaits.duration                            |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLInnoDBLogWaits.val                                 |             | <code>10</code>                    |
| alert.groups.database.rules.mySQLInnoDBLogWaits.severity                            |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLRestarted.enabled                                  |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLRestarted.duration                                 |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLRestarted.val                                      |             | <code>60</code>                    |
| alert.groups.database.rules.mySQLRestarted.severity                                 |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLHighQPS.enabled                                    |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighQPS.duration                                   |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLHighQPS.val                                        |             | <code>1000</code>                  |
| alert.groups.database.rules.mySQLHighQPS.severity                                   |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLHighIncomingBytes.enabled                          |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighIncomingBytes.duration                         |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLHighIncomingBytes.val                              |             | <code>1048576 # 1MB</code>         |
| alert.groups.database.rules.mySQLHighIncomingBytes.severity                         |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLHighOutgoingBytes.enabled                          |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighOutgoingBytes.duration                         |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLHighOutgoingBytes.val                              |             | <code>1048576 # 1MB</code>         |
| alert.groups.database.rules.mySQLHighOutgoingBytes.severity                         |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLTooManyOpenFiles.enabled                           |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLTooManyOpenFiles.duration                          |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLTooManyOpenFiles.val                               |             | <code>80</code>                    |
| alert.groups.database.rules.mySQLTooManyOpenFiles.severity                          |             | <code>warning</code>               |
| alert.groups.group.enabled                                                          |             | <code>true</code>                  |
| alert.groups.group.rules.mySQLHighReplicationDelay.enabled                          |             | <code>true</code>                  |
| alert.groups.group.rules.mySQLHighReplicationDelay.val                              |             | <code>0.5 # second</code>          |
| alert.groups.group.rules.mySQLHighReplicationDelay.duration                         |             | <code>"5m"</code>                  |
| alert.groups.group.rules.mySQLHighReplicationDelay.severity                         |             | <code>warning</code>               |
| alert.groups.group.rules.mySQLHighReplicationTransportTime.enabled                  |             | <code>true</code>                  |
| alert.groups.group.rules.mySQLHighReplicationTransportTime.val                      |             | <code>0.5 # second</code>          |
| alert.groups.group.rules.mySQLHighReplicationTransportTime.duration                 |             | <code>"5m"</code>                  |
| alert.groups.group.rules.mySQLHighReplicationTransportTime.severity                 |             | <code>warning</code>               |
| alert.groups.group.rules.mySQLHighReplicationApplyTime.enabled                      |             | <code>true</code>                  |
| alert.groups.group.rules.mySQLHighReplicationApplyTime.val                          |             | <code>0.5 # second</code>          |
| alert.groups.group.rules.mySQLHighReplicationApplyTime.duration                     |             | <code>"5m"</code>                  |
| alert.groups.group.rules.mySQLHighReplicationApplyTime.severity                     |             | <code>warning</code>               |
| alert.groups.group.rules.mySQLReplicationHighTransactionTime.enabled                |             | <code>true</code>                  |
| alert.groups.group.rules.mySQLReplicationHighTransactionTime.val                    |             | <code>0.5 # second</code>          |
| alert.groups.group.rules.mySQLReplicationHighTransactionTime.duration               |             | <code>"5m"</code>                  |
| alert.groups.group.rules.mySQLReplicationHighTransactionTime.severity               |             | <code>warning</code>               |
| alert.groups.kubedb.enabled                                                         |             | <code>true</code>                  |
| alert.groups.kubedb.rules.kubeDBMySQLPhaseNotReady.enabled                          |             | <code>true</code>                  |
| alert.groups.kubedb.rules.kubeDBMySQLPhaseNotReady.duration                         |             | <code>"1m"</code>                  |
| alert.groups.kubedb.rules.kubeDBMySQLPhaseNotReady.severity                         |             | <code>critical</code>              |
| alert.groups.kubedb.rules.kubeDBMySQLPhaseCritical.enabled                          |             | <code>true</code>                  |
| alert.groups.kubedb.rules.kubeDBMySQLPhaseCritical.duration                         |             | <code>"15m"</code>                 |
| alert.groups.kubedb.rules.kubeDBMySQLPhaseCritical.severity                         |             | <code>warning</code>               |
| alert.groups.opsrequest.enabled                                                     |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestOnProgress.enabled               |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestOnProgress.duration              |             | <code>"0m"</code>                  |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestOnProgress.severity              |             | <code>warning</code>               |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestStatusProgressingToLong.enabled  |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestStatusProgressingToLong.duration |             | <code>"30m"</code>                 |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestStatusProgressingToLong.severity |             | <code>critical</code>              |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestFailed.enabled                   |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestFailed.duration                  |             | <code>"0m"</code>                  |
| alert.groups.opsrequest.rules.kubeDBMySQLOpsRequestFailed.severity                  |             | <code>critical</code>              |
| alert.groups.stash.enabled                                                          |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashBackupSessionFailed.enabled                      |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashBackupSessionFailed.duration                     |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mySQLStashBackupSessionFailed.severity                     |             | <code>critical</code>              |
| alert.groups.stash.rules.mySQLStashRestoreSessionFailed.enabled                     |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashRestoreSessionFailed.duration                    |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mySQLStashRestoreSessionFailed.severity                    |             | <code>critical</code>              |
| alert.groups.stash.rules.mySQLStashNoBackupSessionForTooLong.enabled                |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashNoBackupSessionForTooLong.duration               |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mySQLStashNoBackupSessionForTooLong.val                    |             | <code>18000</code>                 |
| alert.groups.stash.rules.mySQLStashNoBackupSessionForTooLong.severity               |             | <code>warning</code>               |
| alert.groups.stash.rules.mySQLStashRepositoryCorrupted.enabled                      |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashRepositoryCorrupted.duration                     |             | <code>"5m"</code>                  |
| alert.groups.stash.rules.mySQLStashRepositoryCorrupted.severity                     |             | <code>critical</code>              |
| alert.groups.stash.rules.mySQLStashRepositoryStorageRunningLow.enabled              |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashRepositoryStorageRunningLow.duration             |             | <code>"5m"</code>                  |
| alert.groups.stash.rules.mySQLStashRepositoryStorageRunningLow.val                  |             | <code>10737418240 # 10GB</code>    |
| alert.groups.stash.rules.mySQLStashRepositoryStorageRunningLow.severity             |             | <code>waring</code>                |
| alert.groups.stash.rules.mySQLStashBackupSessionPeriodTooLong.enabled               |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashBackupSessionPeriodTooLong.duration              |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mySQLStashBackupSessionPeriodTooLong.val                   |             | <code>1800 # 30 minute</code>      |
| alert.groups.stash.rules.mySQLStashBackupSessionPeriodTooLong.severity              |             | <code>waring</code>                |
| alert.groups.stash.rules.mySQLStashRestoreSessionPeriodTooLong.enabled              |             | <code>true</code>                  |
| alert.groups.stash.rules.mySQLStashRestoreSessionPeriodTooLong.duration             |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mySQLStashRestoreSessionPeriodTooLong.val                  |             | <code>1800 # 30 minute</code>      |
| alert.groups.stash.rules.mySQLStashRestoreSessionPeriodTooLong.severity             |             | <code>waring</code>                |
| alert.groups.schema.enabled                                                         |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaPendingForTooLong.enabled                |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaPendingForTooLong.duration               |             | <code>"30m"</code>                 |
| alert.groups.schema.rules.kubeDBMySQLSchemaPendingForTooLong.severity               |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMySQLSchemaInProgressForTooLong.enabled             |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaInProgressForTooLong.duration            |             | <code>"30m"</code>                 |
| alert.groups.schema.rules.kubeDBMySQLSchemaInProgressForTooLong.severity            |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMySQLSchemaTerminatingForTooLong.enabled            |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaTerminatingForTooLong.duration           |             | <code>"30m"</code>                 |
| alert.groups.schema.rules.kubeDBMySQLSchemaTerminatingForTooLong.severity           |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMySQLSchemaFailed.enabled                           |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaFailed.duration                          |             | <code>"0m"</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaFailed.severity                          |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMySQLSchemaExpired.enabled                          |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaExpired.duration                         |             | <code>"0m"</code>                  |
| alert.groups.schema.rules.kubeDBMySQLSchemaExpired.severity                         |             | <code>warning</code>               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mysql appscode/mysql -n demo --create-namespace --version=v0.1.0 --set namespace=demo
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mysql appscode/mysql -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

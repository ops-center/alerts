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

|                                       Parameter                                       | Description |              Default               |
|---------------------------------------------------------------------------------------|-------------|------------------------------------|
| namespace                                                                             |             | <code>demo</code>                  |
| dbName                                                                                |             | <code>sample-mariadb</code>        |
| alert.ruleSelector.app                                                                |             | <code>kube-prometheus-stack</code> |
| alert.ruleSelector.release                                                            |             | <code>prometheus</code>            |
| alert.groups.database.enabled                                                         |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLInstanceDown.enabled                                 |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLInstanceDown.duration                                |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLInstanceDown.severity                                |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLServiceDown.enabled                                  |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLServiceDown.duration                                 |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLServiceDown.severity                                 |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLTooManyConnections.enabled                           |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLTooManyConnections.duration                          |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLTooManyConnections.val                               |             | <code>80</code>                    |
| alert.groups.database.rules.mySQLTooManyConnections.severity                          |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLHighThreadsRunning.enabled                           |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighThreadsRunning.duration                          |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLHighThreadsRunning.val                               |             | <code>60</code>                    |
| alert.groups.database.rules.mySQLHighThreadsRunning.severity                          |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLSlowQueries.enabled                                  |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLSlowQueries.duration                                 |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLSlowQueries.severity                                 |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLInnoDBLogWaits.enabled                               |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLInnoDBLogWaits.duration                              |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLInnoDBLogWaits.val                                   |             | <code>10</code>                    |
| alert.groups.database.rules.mySQLInnoDBLogWaits.severity                              |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLRestarted.enabled                                    |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLRestarted.duration                                   |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLRestarted.val                                        |             | <code>60</code>                    |
| alert.groups.database.rules.mySQLRestarted.severity                                   |             | <code>warning</code>               |
| alert.groups.database.rules.mySQLHighQPS.enabled                                      |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighQPS.duration                                     |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLHighQPS.val                                          |             | <code>1000</code>                  |
| alert.groups.database.rules.mySQLHighQPS.severity                                     |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLHighIncomingBytes.enabled                            |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighIncomingBytes.duration                           |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLHighIncomingBytes.val                                |             | <code>1048576 # 1MB</code>         |
| alert.groups.database.rules.mySQLHighIncomingBytes.severity                           |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLHighOutgoingBytes.enabled                            |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLHighOutgoingBytes.duration                           |             | <code>"0m"</code>                  |
| alert.groups.database.rules.mySQLHighOutgoingBytes.val                                |             | <code>1048576 # 1MB</code>         |
| alert.groups.database.rules.mySQLHighOutgoingBytes.severity                           |             | <code>critical</code>              |
| alert.groups.database.rules.mySQLTooManyOpenFiles.enabled                             |             | <code>true</code>                  |
| alert.groups.database.rules.mySQLTooManyOpenFiles.duration                            |             | <code>"2m"</code>                  |
| alert.groups.database.rules.mySQLTooManyOpenFiles.val                                 |             | <code>80</code>                    |
| alert.groups.database.rules.mySQLTooManyOpenFiles.severity                            |             | <code>warning</code>               |
| alert.groups.cluster.enabled                                                          |             | <code>true</code>                  |
| alert.groups.cluster.rules.galeraReplicationLatencyTooLong.enabled                    |             | <code>true</code>                  |
| alert.groups.cluster.rules.galeraReplicationLatencyTooLong.val                        |             | <code>0.1</code>                   |
| alert.groups.cluster.rules.galeraReplicationLatencyTooLong.duration                   |             | <code>"5m"</code>                  |
| alert.groups.cluster.rules.galeraReplicationLatencyTooLong.severity                   |             | <code>warning</code>               |
| alert.groups.kubedb.enabled                                                           |             | <code>true</code>                  |
| alert.groups.kubedb.rules.kubeDBMariaDBPhaseNotReady.enabled                          |             | <code>true</code>                  |
| alert.groups.kubedb.rules.kubeDBMariaDBPhaseNotReady.duration                         |             | <code>"1m"</code>                  |
| alert.groups.kubedb.rules.kubeDBMariaDBPhaseNotReady.severity                         |             | <code>critical</code>              |
| alert.groups.kubedb.rules.kubeDBMariaDBPhaseCritical.enabled                          |             | <code>true</code>                  |
| alert.groups.kubedb.rules.kubeDBMariaDBPhaseCritical.duration                         |             | <code>"15m"</code>                 |
| alert.groups.kubedb.rules.kubeDBMariaDBPhaseCritical.severity                         |             | <code>warning</code>               |
| alert.groups.opsrequest.enabled                                                       |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestOnProgress.enabled               |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestOnProgress.duration              |             | <code>"0m"</code>                  |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestOnProgress.severity              |             | <code>warning</code>               |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestStatusProgressingToLong.enabled  |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestStatusProgressingToLong.duration |             | <code>"30m"</code>                 |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestStatusProgressingToLong.severity |             | <code>critical</code>              |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestFailed.enabled                   |             | <code>true</code>                  |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestFailed.duration                  |             | <code>"0m"</code>                  |
| alert.groups.opsrequest.rules.kubeDBMariaDBOpsRequestFailed.severity                  |             | <code>critical</code>              |
| alert.groups.stash.enabled                                                            |             | <code>true</code>                  |
| alert.groups.stash.rules.mariaDBStashBackupSessionFailed.enabled                      |             | <code>true</code>                  |
| alert.groups.stash.rules.mariaDBStashBackupSessionFailed.duration                     |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mariaDBStashBackupSessionFailed.severity                     |             | <code>critical</code>              |
| alert.groups.stash.rules.mariaDBStashRestoreSessionFailed.enabled                     |             | <code>true</code>                  |
| alert.groups.stash.rules.mariaDBStashRestoreSessionFailed.duration                    |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mariaDBStashRestoreSessionFailed.severity                    |             | <code>critical</code>              |
| alert.groups.stash.rules.mariaDBStashNoBackupSessionForTooLong.enabled                |             | <code>true</code>                  |
| alert.groups.stash.rules.mariaDBStashNoBackupSessionForTooLong.duration               |             | <code>"0m"</code>                  |
| alert.groups.stash.rules.mariaDBStashNoBackupSessionForTooLong.val                    |             | <code>18000</code>                 |
| alert.groups.stash.rules.mariaDBStashNoBackupSessionForTooLong.severity               |             | <code>warning</code>               |
| alert.groups.stash.rules.mariaDBStashRepositoryCorrupted.enabled                      |             | <code>true</code>                  |
| alert.groups.stash.rules.mariaDBStashRepositoryCorrupted.duration                     |             | <code>"5m"</code>                  |
| alert.groups.stash.rules.mariaDBStashRepositoryCorrupted.severity                     |             | <code>critical</code>              |
| alert.groups.stash.rules.mariaDBStashRepositoryStorageRunningLow.enabled              |             | <code>true</code>                  |
| alert.groups.stash.rules.mariaDBStashRepositoryStorageRunningLow.duration             |             | <code>"5m"</code>                  |
| alert.groups.stash.rules.mariaDBStashRepositoryStorageRunningLow.val                  |             | <code>10737418240 # 10GB</code>    |
| alert.groups.stash.rules.mariaDBStashRepositoryStorageRunningLow.severity             |             | <code>waring</code>                |
| alert.groups.schema.enabled                                                           |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaPendingForTooLong.enabled                |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaPendingForTooLong.duration               |             | <code>"30m"</code>                 |
| alert.groups.schema.rules.kubeDBMariaDBSchemaPendingForTooLong.severity               |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMariaDBSchemaInProgressForTooLong.enabled             |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaInProgressForTooLong.duration            |             | <code>"30m"</code>                 |
| alert.groups.schema.rules.kubeDBMariaDBSchemaInProgressForTooLong.severity            |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMariaDBSchemaTerminatingForTooLong.enabled            |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaTerminatingForTooLong.duration           |             | <code>"30m"</code>                 |
| alert.groups.schema.rules.kubeDBMariaDBSchemaTerminatingForTooLong.severity           |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMariaDBSchemaFailed.enabled                           |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaFailed.duration                          |             | <code>"0m"</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaFailed.severity                          |             | <code>warning</code>               |
| alert.groups.schema.rules.kubeDBMariaDBSchemaExpired.enabled                          |             | <code>true</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaExpired.duration                         |             | <code>"0m"</code>                  |
| alert.groups.schema.rules.kubeDBMariaDBSchemaExpired.severity                         |             | <code>warning</code>               |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i mariadb appscode/mariadb -n demo --create-namespace --version=v0.1.0 --set namespace=demo
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i mariadb appscode/mariadb -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

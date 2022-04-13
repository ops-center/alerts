# Elasticsearch alerts

[Elasticsearch alerts by AppsCode](https://github.com/appscode/alerts) - Elasticsearch alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/elasticsearch --version=v0.1.0
$ helm upgrade -i elasticsearch appscode/elasticsearch -n demo --create-namespace --version=v0.1.0
```

## Introduction

This chart deploys Elasticsearch alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install/upgrade the chart with the release name `elasticsearch`:

```bash
$ helm upgrade -i elasticsearch appscode/elasticsearch -n demo --create-namespace --version=v0.1.0
```

The command deploys Elasticsearch alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `elasticsearch`:

```bash
$ helm uninstall elasticsearch -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `elasticsearch` chart and their default values.

|                                   Parameter                                   |                   Description                   |                     Default                      |
|-------------------------------------------------------------------------------|-------------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                       |                                                 | <code>kubedb.com</code>                          |
| metadata.resource.kind                                                        |                                                 | <code>Elasticsearch</code>                       |
| metadata.resource.name                                                        |                                                 | <code>elasticsearches</code>                     |
| metadata.resource.scope                                                       |                                                 | <code>Namespaced</code>                          |
| metadata.resource.version                                                     |                                                 | <code>v1alpha2</code>                            |
| metadata.release.name                                                         | Release name                                    | <code>""</code>                                  |
| metadata.release.namespace                                                    | Release namespace                               | <code>""</code>                                  |
| spec.alert.enabled                                                            | # Enable PrometheusRule alerts                  | <code>true</code>                                |
| spec.alert.labels                                                             | # Labels for default rules                      | <code>{"release":"kube-prometheus-stack"}</code> |
| spec.alert.annotations                                                        | # Annotations for default rules                 | <code>{}</code>                                  |
| spec.alert.additionalRuleLabels                                               | # Additional labels for PrometheusRule alerts   | <code>{}</code>                                  |
| spec.alert.groups.database.enabled                                            |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.enabled        |                                                 | <code>false</code>                               |
| spec.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.val            | The heap usage is over 90%                      | <code>90</code>                                  |
| spec.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.duration       |                                                 | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchHeapUsageTooHigh.severity       |                                                 | <code>critical</code>                            |
| spec.alert.groups.database.rules.elasticsearchHeapUsageWarning.enabled        |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchHeapUsageWarning.val            | The heap usage is over 80%                      | <code>80</code>                                  |
| spec.alert.groups.database.rules.elasticsearchHeapUsageWarning.duration       |                                                 | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchHeapUsageWarning.severity       |                                                 | <code>warning</code>                             |
| spec.alert.groups.database.rules.elasticsearchDiskOutOfSpace.enabled          |                                                 | <code>false</code>                               |
| spec.alert.groups.database.rules.elasticsearchDiskOutOfSpace.val              | The disk usage is over 90%. Value range: 0-100. | <code>90</code>                                  |
| spec.alert.groups.database.rules.elasticsearchDiskOutOfSpace.duration         |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchDiskOutOfSpace.severity         |                                                 | <code>critical</code>                            |
| spec.alert.groups.database.rules.elasticsearchDiskSpaceLow.enabled            |                                                 | <code>false</code>                               |
| spec.alert.groups.database.rules.elasticsearchDiskSpaceLow.val                | The disk usage is over 80%. Value range: 0-100. | <code>80</code>                                  |
| spec.alert.groups.database.rules.elasticsearchDiskSpaceLow.duration           |                                                 | <code>"2m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchDiskSpaceLow.severity           |                                                 | <code>warning</code>                             |
| spec.alert.groups.database.rules.elasticsearchClusterRed.enabled              |                                                 | <code>false</code>                               |
| spec.alert.groups.database.rules.elasticsearchClusterRed.duration             |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchClusterRed.severity             |                                                 | <code>critical</code>                            |
| spec.alert.groups.database.rules.elasticsearchClusterYellow.enabled           |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchClusterYellow.duration          |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchClusterYellow.severity          |                                                 | <code>warning</code>                             |
| spec.alert.groups.database.rules.elasticsearchHealthyNodes.enabled            |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchHealthyNodes.val                | should have at least 3 healthy nodes            | <code>3</code>                                   |
| spec.alert.groups.database.rules.elasticsearchHealthyNodes.duration           |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchHealthyNodes.severity           |                                                 | <code>critical</code>                            |
| spec.alert.groups.database.rules.elasticsearchHealthyDataNodes.enabled        |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchHealthyDataNodes.val            | should have at least 3 healthy data nodes       | <code>3</code>                                   |
| spec.alert.groups.database.rules.elasticsearchHealthyDataNodes.duration       |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchHealthyDataNodes.severity       |                                                 | <code>critical</code>                            |
| spec.alert.groups.database.rules.elasticsearchRelocatingShards.enabled        |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchRelocatingShards.duration       |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchRelocatingShards.severity       |                                                 | <code>info</code>                                |
| spec.alert.groups.database.rules.elasticsearchInitializingShards.enabled      |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchInitializingShards.duration     |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchInitializingShards.severity     |                                                 | <code>info</code>                                |
| spec.alert.groups.database.rules.elasticsearchUnassignedShards.enabled        |                                                 | <code>true</code>                                |
| spec.alert.groups.database.rules.elasticsearchUnassignedShards.duration       |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchUnassignedShards.severity       |                                                 | <code>critical</code>                            |
| spec.alert.groups.database.rules.elasticsearchPendingTasks.enabled            |                                                 | <code>false</code>                               |
| spec.alert.groups.database.rules.elasticsearchPendingTasks.duration           |                                                 | <code>"15m"</code>                               |
| spec.alert.groups.database.rules.elasticsearchPendingTasks.severity           |                                                 | <code>warning</code>                             |
| spec.alert.groups.database.rules.elasticsearchNoNewDocuments10m.enabled       |                                                 | <code>false</code>                               |
| spec.alert.groups.database.rules.elasticsearchNoNewDocuments10m.duration      |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.database.rules.elasticsearchNoNewDocuments10m.severity      |                                                 | <code>info</code>                                |
| spec.alert.groups.provisioner.enabled                                         |                                                 | <code>true</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseNotReady.enabled                  |                                                 | <code>true</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseNotReady.duration                 |                                                 | <code>"1m"</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseNotReady.severity                 |                                                 | <code>critical</code>                            |
| spec.alert.groups.provisioner.rules.appPhaseCritical.enabled                  |                                                 | <code>true</code>                                |
| spec.alert.groups.provisioner.rules.appPhaseCritical.duration                 |                                                 | <code>"15m"</code>                               |
| spec.alert.groups.provisioner.rules.appPhaseCritical.severity                 |                                                 | <code>warning</code>                             |
| spec.alert.groups.opsManager.enabled                                          |                                                 | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestOnProgress.enabled               |                                                 | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestOnProgress.duration              |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestOnProgress.severity              |                                                 | <code>warning</code>                             |
| spec.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled  |                                                 | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration |                                                 | <code>"30m"</code>                               |
| spec.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity |                                                 | <code>critical</code>                            |
| spec.alert.groups.opsManager.rules.opsRequestFailed.enabled                   |                                                 | <code>true</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestFailed.duration                  |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.opsManager.rules.opsRequestFailed.severity                  |                                                 | <code>critical</code>                            |
| spec.alert.groups.stash.enabled                                               |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.backupSessionFailed.enabled                     |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.backupSessionFailed.duration                    |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.backupSessionFailed.severity                    |                                                 | <code>critical</code>                            |
| spec.alert.groups.stash.rules.restoreSessionFailed.enabled                    |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.restoreSessionFailed.duration                   |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.restoreSessionFailed.severity                   |                                                 | <code>critical</code>                            |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.enabled               |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.duration              |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.val                   |                                                 | <code>18000</code>                               |
| spec.alert.groups.stash.rules.noBackupSessionForTooLong.severity              |                                                 | <code>warning</code>                             |
| spec.alert.groups.stash.rules.repositoryCorrupted.enabled                     |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.repositoryCorrupted.duration                    |                                                 | <code>"5m"</code>                                |
| spec.alert.groups.stash.rules.repositoryCorrupted.severity                    |                                                 | <code>critical</code>                            |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.enabled             |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.duration            |                                                 | <code>"5m"</code>                                |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.val                 |                                                 | <code>10737418240 # 10GB</code>                  |
| spec.alert.groups.stash.rules.repositoryStorageRunningLow.severity            |                                                 | <code>waring</code>                              |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled              |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.duration             |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.val                  |                                                 | <code>1800 # 30 minute</code>                    |
| spec.alert.groups.stash.rules.backupSessionPeriodTooLong.severity             |                                                 | <code>waring</code>                              |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled             |                                                 | <code>true</code>                                |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration            |                                                 | <code>"0m"</code>                                |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                 |                                                 | <code>1800 # 30 minute</code>                    |
| spec.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity            |                                                 | <code>waring</code>                              |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i elasticsearch appscode/elasticsearch -n demo --create-namespace --version=v0.1.0 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i elasticsearch appscode/elasticsearch -n demo --create-namespace --version=v0.1.0 --values values.yaml
```

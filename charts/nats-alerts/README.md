# NATS alerts

[NATS alerts by AppsCode](https://github.com/appscode/alerts) - NATS alerts for NATS clusters

## TL;DR;

```bash
$ helm repo add appscode oci://ghcr.io/appscode-charts
$ helm repo update
$ helm search repo appscode/nats-alerts --version=v2025.1.0
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2025.1.0
```

## Introduction

This chart deploys NATS alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.21+
- Prometheus operator

## Installing the Chart

To install/upgrade the chart with the release name `nats-alerts`:

```bash
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2025.1.0
```

The command deploys NATS alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `nats-alerts`:

```bash
$ helm uninstall nats-alerts -n monitoring
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `nats-alerts` chart and their default values.

|                                     Parameter                                     |                  Description                  |                     Default                      |
|-----------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                           |                                               | <code>alerts.appscode.com</code>                 |
| metadata.resource.kind                                                            |                                               | <code>NATSAlerts</code>                          |
| metadata.resource.name                                                            |                                               | <code>natsalerts</code>                          |
| metadata.resource.scope                                                           |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                         |                                               | <code>v1alpha2</code>                            |
| metadata.release.name                                                             | Release name                                  | <code>""</code>                                  |
| metadata.release.namespace                                                        | Release namespace                             | <code>""</code>                                  |
| form.alert.enabled                                                                | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                                 | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                            | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                                   | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                                |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsJetStreamHighMemoryUsage.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighMemoryUsage.duration            |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighMemoryUsage.severity            |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.natsJetStreamHighMemoryUsage.val                 |                                               | <code>90</code>                                  |
| form.alert.groups.database.rules.natsJetStreamHighStorageUsage.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighStorageUsage.duration           |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighStorageUsage.severity           |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.natsJetStreamHighStorageUsage.val                |                                               | <code>90</code>                                  |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessages.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessages.duration        |                                               | <code>"10m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessages.severity        |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessages.val             |                                               | <code>5000</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessagesWarning.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessagesWarning.duration |                                               | <code>"10m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessagesWarning.severity |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsJetStreamHighPendingMessagesWarning.val      |                                               | <code>2000</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighAckPending.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighAckPending.duration             |                                               | <code>"10m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamHighAckPending.severity             |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.natsJetStreamHighAckPending.val                  |                                               | <code>2000</code>                                |
| form.alert.groups.database.rules.natsJetStreamConsumerStalled.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamConsumerStalled.duration            |                                               | <code>"15m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamConsumerStalled.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsJetStreamHighMessageCount.enabled            |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighMessageCount.duration           |                                               | <code>"15m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamHighMessageCount.severity           |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsJetStreamHighMessageCount.val                |                                               | <code>1000000</code>                             |
| form.alert.groups.database.rules.natsJetStreamHighStreamSize.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamHighStreamSize.duration             |                                               | <code>"15m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamHighStreamSize.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsJetStreamHighStreamSize.val                  |                                               | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.database.rules.natsJetStreamNoIngestion.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsJetStreamNoIngestion.duration                |                                               | <code>"15m"</code>                               |
| form.alert.groups.database.rules.natsJetStreamNoIngestion.severity                |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsSuddenConnectionDrop.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsSuddenConnectionDrop.duration                |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.natsSuddenConnectionDrop.severity                |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsSuddenConnectionDrop.val                     |                                               | <code>5</code>                                   |
| form.alert.groups.database.rules.natsHighActiveConnections.enabled                |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsHighActiveConnections.duration               |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.natsHighActiveConnections.severity               |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsHighActiveConnections.val                    |                                               | <code>1000</code>                                |
| form.alert.groups.database.rules.natsSuddenConsumerDrop.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsSuddenConsumerDrop.duration                  |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.natsSuddenConsumerDrop.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsSuddenConsumerDrop.val                       |                                               | <code>2</code>                                   |
| form.alert.groups.database.rules.natsHighTotalConsumers.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.natsHighTotalConsumers.duration                  |                                               | <code>"5m"</code>                                |
| form.alert.groups.database.rules.natsHighTotalConsumers.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.natsHighTotalConsumers.val                       |                                               | <code>100</code>                                 |
| form.alert.groups.provisioner.enabled                                             |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                     |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                     |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                      |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                     |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                     |                                               | <code>critical</code>                            |
| form.alert.groups.opsManager.enabled                                              |                                               | <code>warning</code>                             |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestOnProgress.severity                  |                                               | <code>info</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.enabled      |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.duration     |                                               | <code>"30m"</code>                               |
| form.alert.groups.opsManager.rules.opsRequestStatusProgressingToLong.severity     |                                               | <code>critical</code>                            |
| form.alert.groups.opsManager.rules.opsRequestFailed.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.duration                      |                                               | <code>"0m"</code>                                |
| form.alert.groups.opsManager.rules.opsRequestFailed.severity                      |                                               | <code>critical</code>                            |
| form.alert.groups.stash.enabled                                                   |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.backupSessionFailed.enabled                         |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.duration                        |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionFailed.severity                        |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.restoreSessionFailed.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.duration                       |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionFailed.severity                       |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.duration                  |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.noBackupSessionForTooLong.val                       |                                               | <code>18000</code>                               |
| form.alert.groups.stash.rules.repositoryCorrupted.enabled                         |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.duration                        |                                               | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryCorrupted.severity                        |                                               | <code>critical</code>                            |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.duration                |                                               | <code>"5m"</code>                                |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.severity                |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.repositoryStorageRunningLow.val                     |                                               | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.duration                 |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.severity                 |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.backupSessionPeriodTooLong.val                      |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.enabled                 |                                               | <code>true</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.duration                |                                               | <code>"0m"</code>                                |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.severity                |                                               | <code>warning</code>                             |
| form.alert.groups.stash.rules.restoreSessionPeriodTooLong.val                     |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.kubeStash.enabled                                               |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.backupSessionFailed.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionFailed.duration                    |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionFailed.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.enabled                    |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.duration                   |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionFailed.severity                   |                                               | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.duration              |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.noBackupSessionForTooLong.val                   |                                               | <code>18000</code>                               |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.enabled                     |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.duration                    |                                               | <code>"5m"</code>                                |
| form.alert.groups.kubeStash.rules.repositoryCorrupted.severity                    |                                               | <code>critical</code>                            |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.duration            |                                               | <code>"5m"</code>                                |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.repositoryStorageRunningLow.val                 |                                               | <code>10737418240 # 10GB</code>                  |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.duration             |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.backupSessionPeriodTooLong.val                  |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.duration            |                                               | <code>"0m"</code>                                |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.kubeStash.rules.restoreSessionPeriodTooLong.val                 |                                               | <code>1800 # 30 minute</code>                    |
| form.alert.groups.schemaManager.enabled                                           |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.duration            |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaPendingForTooLong.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.duration         |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaInProgressForTooLong.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.enabled         |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.duration        |                                               | <code>"30m"</code>                               |
| form.alert.groups.schemaManager.rules.schemaTerminatingForTooLong.severity        |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaFailed.enabled                        |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaFailed.duration                       |                                               | <code>"0m"</code>                                |
| form.alert.groups.schemaManager.rules.schemaFailed.severity                       |                                               | <code>warning</code>                             |
| form.alert.groups.schemaManager.rules.schemaExpired.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.schemaManager.rules.schemaExpired.duration                      |                                               | <code>"0m"</code>                                |
| form.alert.groups.schemaManager.rules.schemaExpired.severity                      |                                               | <code>warning</code>                             |
| grafana.enabled                                                                   |                                               | <code>false</code>                               |
| grafana.version                                                                   |                                               | <code>8.2.3</code>                               |
| grafana.jobName                                                                   |                                               | <code>nats-stats</code>                          |
| grafana.url                                                                       |                                               | <code>""</code>                                  |
| grafana.apikey                                                                    |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2025.1.0 --set metadata.resource.group=alerts.appscode.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i nats-alerts appscode/nats-alerts -n monitoring --create-namespace --version=v2025.1.0 --values values.yaml
```

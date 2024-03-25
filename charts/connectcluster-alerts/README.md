# ConnectCluster alerts

[ConnectCluster alerts by AppsCode](https://github.com/appscode/alerts) - ConnectCluster alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/connectcluster-alerts --version=v2023.05.09
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2023.05.09
```

## Introduction

This chart deploys ConnectCluster alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `connectcluster`:

```bash
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2023.05.09
```

The command deploys ConnectCluster alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `connectcluster`:

```bash
$ helm uninstall connectcluster -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `connectcluster-alerts` chart and their default values.

|                                     Parameter                                      |                  Description                  |                     Default                      |
|------------------------------------------------------------------------------------|-----------------------------------------------|--------------------------------------------------|
| metadata.resource.group                                                            |                                               | <code>kafka.kubedb.com</code>                    |
| metadata.resource.kind                                                             |                                               | <code>ConnectCluster</code>                      |
| metadata.resource.name                                                             |                                               | <code>connectclusters</code>                     |
| metadata.resource.scope                                                            |                                               | <code>Namespaced</code>                          |
| metadata.resource.version                                                          |                                               | <code>v1alpha1</code>                            |
| metadata.release.name                                                              | Release name                                  | <code>"connect-cluster"</code>                   |
| metadata.release.namespace                                                         | Release namespace                             | <code>"demo"</code>                              |
| form.alert.enabled                                                                 | # Enable PrometheusRule alerts                | <code>warning</code>                             |
| form.alert.labels                                                                  | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code> |
| form.alert.annotations                                                             | # Annotations for default rules               | <code>{}</code>                                  |
| form.alert.additionalRuleLabels                                                    | # Additional labels for PrometheusRule alerts | <code>{}</code>                                  |
| form.alert.groups.database.enabled                                                 |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.connectClusterWorkerDown.enabled                  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectClusterWorkerDown.duration                 |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.connectClusterWorkerDown.severity                 |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.connectClusterTooManyConnections.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectClusterTooManyConnections.duration         |                                               | <code>"2m"</code>                                |
| form.alert.groups.database.rules.connectClusterTooManyConnections.val              |                                               | <code>80</code>                                  |
| form.alert.groups.database.rules.connectClusterTooManyConnections.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.connectClusterPhaseCritical.enabled               |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectClusterPhaseCritical.duration              |                                               | <code>"3m"</code>                                |
| form.alert.groups.database.rules.connectClusterPhaseCritical.severity              |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.connectclusterConnectorCount.enabled              |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectclusterConnectorCount.duration             |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.connectclusterConnectorCount.severity             |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.connectclusterConnectorCount.val                  |                                               | <code>50</code>                                  |
| form.alert.groups.database.rules.connectclusterCoordinatorRebalanceFailed.enabled  |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectclusterCoordinatorRebalanceFailed.duration |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.connectclusterCoordinatorRebalanceFailed.severity |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.connectclusterCoordinatorRebalanceFailed.val      |                                               | <code>10</code>                                  |
| form.alert.groups.database.rules.connectclusterTasksErrorTotalRetries.enabled      |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectclusterTasksErrorTotalRetries.duration     |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.connectclusterTasksErrorTotalRetries.severity     |                                               | <code>critical</code>                            |
| form.alert.groups.database.rules.connectclusterTasksErrorTotalRetries.val          |                                               | <code>5</code>                                   |
| form.alert.groups.database.rules.connectclusterTaskTotal.enabled                   |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectclusterTaskTotal.duration                  |                                               | <code>"1m"</code>                                |
| form.alert.groups.database.rules.connectclusterTaskTotal.severity                  |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.connectclusterTaskTotal.val                       |                                               | <code>150</code>                                 |
| form.alert.groups.database.rules.connectclusterTaskTotalFailed.enabled             |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectclusterTaskTotalFailed.duration            |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.connectclusterTaskTotalFailed.severity            |                                               | <code>warning</code>                             |
| form.alert.groups.database.rules.connectclusterTaskTotalDestroyed.enabled          |                                               | <code>true</code>                                |
| form.alert.groups.database.rules.connectclusterTaskTotalDestroyed.duration         |                                               | <code>"0m"</code>                                |
| form.alert.groups.database.rules.connectclusterTaskTotalDestroyed.severity         |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.enabled                                              |                                               | <code>warning</code>                             |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration                      |                                               | <code>"1m"</code>                                |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity                      |                                               | <code>critical</code>                            |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled                       |                                               | <code>true</code>                                |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration                      |                                               | <code>"15m"</code>                               |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity                      |                                               | <code>warning</code>                             |
| grafana.enabled                                                                    |                                               | <code>false</code>                               |
| grafana.version                                                                    |                                               | <code>8.2.3</code>                               |
| grafana.jobName                                                                    |                                               | <code>kubedb-databases</code>                    |
| grafana.url                                                                        |                                               | <code>""</code>                                  |
| grafana.apikey                                                                     |                                               | <code>""</code>                                  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2023.05.09 --set metadata.resource.group=kafka.kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i connectcluster appscode/connectcluster-alerts -n demo --create-namespace --version=v2023.05.09 --values values.yaml
```

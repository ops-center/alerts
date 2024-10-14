# Cassandra alerts

[Cassandra alerts by AppsCode](https://github.com/appscode/alerts) - Cassandra alerts for KubeDB

## TL;DR;

```bash
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm search repo appscode/cassandra-alerts --version=v2023.05.09
$ helm upgrade -i cassandra appscode/cassandra-alerts -n demo --create-namespace --version=v2023.05.09
```

## Introduction

This chart deploys Cassandra alerts on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.19+

## Installing the Chart

To install/upgrade the chart with the release name `cassandra`:

```bash
$ helm upgrade -i cassandra appscode/cassandra-alerts -n demo --create-namespace --version=v2023.05.09
```

The command deploys Cassandra alerts on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall the `cassandra`:

```bash
$ helm uninstall cassandra -n demo
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `cassandra-alerts` chart and their default values.

|                                Parameter                                |                  Description                  |                          Default                           |
|-------------------------------------------------------------------------|-----------------------------------------------|------------------------------------------------------------|
| metadata.resource.group                                                 |                                               | <code>kubedb.com</code>                                    |
| metadata.resource.kind                                                  |                                               | <code>Cassandra</code>                                     |
| metadata.resource.name                                                  |                                               | <code>cassandras</code>                                    |
| metadata.resource.scope                                                 |                                               | <code>Namespaced</code>                                    |
| metadata.resource.version                                               |                                               | <code>v1</code>                                            |
| metadata.release.name                                                   | Release name                                  | <code>"coreos-prom-memcd"</code>                           |
| metadata.release.namespace                                              | Release namespace                             | <code>"demo"</code>                                        |
| form.alert.enabled                                                      | # Enable PrometheusRule alerts                | <code>warning</code>                                       |
| form.alert.labels                                                       | # Labels for default rules                    | <code>{"release":"kube-prometheus-stack"}</code>           |
| form.alert.annotations                                                  | # Annotations for default rules               | <code>{}</code>                                            |
| form.alert.additionalRuleLabels                                         | # Additional labels for PrometheusRule alerts | <code>{}</code>                                            |
| form.alert.groups.database.enabled                                      |                                               | <code>warning</code>                                       |
| form.alert.groups.database.rules.cassandraDown.enabled                  |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraDown.duration                 |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.cassandraDown.severity                 |                                               | <code>critical</code>                                      |
| form.alert.groups.database.rules.cassandraServiceRespawn.enabled        |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraServiceRespawn.duration       |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.cassandraServiceRespawn.val            |                                               | <code>180</code>                                           |
| form.alert.groups.database.rules.cassandraServiceRespawn.severity       |                                               | <code>critical</code>                                      |
| form.alert.groups.database.rules.cassandraConnectionThrottled.enabled   |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraConnectionThrottled.duration  |                                               | <code>"2m"</code>                                          |
| form.alert.groups.database.rules.cassandraConnectionThrottled.val       |                                               | <code>10</code>                                            |
| form.alert.groups.database.rules.cassandraConnectionThrottled.severity  |                                               | <code>warning</code>                                       |
| form.alert.groups.database.rules.cassandraConnectionsNoneMinor.enabled  |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraConnectionsNoneMinor.duration |                                               | <code>"2m"</code>                                          |
| form.alert.groups.database.rules.cassandraConnectionsNoneMinor.severity |                                               | <code>warning</code>                                       |
| form.alert.groups.database.rules.cassandraItemsNoneMinor.enabled        |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraItemsNoneMinor.duration       |                                               | <code>"2m"</code>                                          |
| form.alert.groups.database.rules.cassandraItemsNoneMinor.severity       |                                               | <code>warning</code>                                       |
| form.alert.groups.database.rules.cassandraEvictionsLimit.enabled        |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraEvictionsLimit.duration       |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.cassandraEvictionsLimit.val            |                                               | <code>10</code>                                            |
| form.alert.groups.database.rules.cassandraEvictionsLimit.severity       |                                               | <code>critical</code>                                      |
| form.alert.groups.database.rules.cassandraMemoryLimit.enable            |                                               | <code>true</code>                                          |
| form.alert.groups.database.rules.cassandraMemoryLimit.duration          |                                               | <code>"0m"</code>                                          |
| form.alert.groups.database.rules.cassandraMemoryLimit.val               |                                               | <code>33554432 # 32MB</code>                               |
| form.alert.groups.database.rules.cassandraMemoryLimit.severity          |                                               | <code>critical</code>                                      |
| form.alert.groups.provisioner.enabled                                   |                                               | <code>warning</code>                                       |
| form.alert.groups.provisioner.rules.appPhaseNotReady.enabled            |                                               | <code>true</code>                                          |
| form.alert.groups.provisioner.rules.appPhaseNotReady.duration           |                                               | <code>"1m"</code>                                          |
| form.alert.groups.provisioner.rules.appPhaseNotReady.severity           |                                               | <code>critical</code>                                      |
| form.alert.groups.provisioner.rules.appPhaseCritical.enabled            |                                               | <code>true</code>                                          |
| form.alert.groups.provisioner.rules.appPhaseCritical.duration           |                                               | <code>"15m"</code>                                         |
| form.alert.groups.provisioner.rules.appPhaseCritical.severity           |                                               | <code>warning</code>                                       |
| grafana.enabled                                                         |                                               | <code>true</code>                                          |
| grafana.version                                                         |                                               | <code>7.5.5</code>                                         |
| grafana.jobName                                                         |                                               | <code>cassandra-stats</code>                               |
| grafana.url                                                             |                                               | <code>"http://prometheus-grafana.monitoring.svc:80"</code> |
| grafana.apikey                                                          |                                               | <code>""</code>                                            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm upgrade -i`. For example:

```bash
$ helm upgrade -i cassandra appscode/cassandra-alerts -n demo --create-namespace --version=v2023.05.09 --set metadata.resource.group=kubedb.com
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```bash
$ helm upgrade -i cassandra appscode/cassandra-alerts -n demo --create-namespace --version=v2023.05.09 --values values.yaml
```

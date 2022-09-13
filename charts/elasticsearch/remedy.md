## TODOs on Elasticsearch Critical Alerts

### Database Alerts

- #### ElasticsearchHeapUsageTooHigh
  - Get the current `heap.percent` for each node using the [_cat/nodes](https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodes.html) API.
  - High heap usage occurs when the garbage collection process cannot keep up. [Reconfigure](https://kubedb.com/docs/v2022.08.08/guides/elasticsearch/configuration/jvm-options/) the default [JVM options](https://www.elastic.co/guide/en/elasticsearch/reference/current/advanced-configuration.html#set-jvm-heap-size) as per your requirement.
  - Contact AppsCode team
- #### ElasticsearchHeapUsageWarning
  - Get the current `heap.percent` for each node using the [_cat/nodes](https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-nodes.html) API
  - Delete indices with old or unnecessary data. Eliminate empty indices. Keep the number of aggregation buckets (size) in your queries to a minimum. Try reducing the size of the large bulk index requests.
  - [Reconfigure](https://kubedb.com/docs/v2022.08.08/guides/elasticsearch/configuration/jvm-options/) the default [JVM options](https://www.elastic.co/guide/en/elasticsearch/reference/current/advanced-configuration.html#set-jvm-heap-size) as per your requirement
  - Contact AppsCode team
- #### ElasticsearchDiskOutOfSpace
  - Reconfigure storage capacity to your cluster -
    - [Horizontally Scale](https://kubedb.com/docs/v2022.08.08/guides/elasticsearch/concepts/elasticsearch-ops-request/#spechorizontalscaling) up your data nodes with `ElasticsearchOpsRequest`
    - [Vertically Scale](https://kubedb.com/docs/v2022.08.08/guides/elasticsearch/concepts/elasticsearch-ops-request/#specverticalscaling) up your data nodes with `ElasticsearchOpsRequest`
  - [Reduce replicas per shard](https://www.elastic.co/guide/en/elasticsearch/reference/current/size-your-shards.html)
  - Contact AppsCode team
- #### ElasticsearchDiskSpaceLow
  - Reconfigure storage capacity to your cluster -
    - [Horizontally Scale](https://kubedb.com/docs/v2022.08.08/guides/elasticsearch/concepts/elasticsearch-ops-request/#spechorizontalscaling) up your data nodes with `ElasticsearchOpsRequest`
    - [Vertically Scale](https://kubedb.com/docs/v2022.08.08/guides/elasticsearch/concepts/elasticsearch-ops-request/#specverticalscaling) up your data nodes with `ElasticsearchOpsRequest`
  - [Reduce replicas per shard](https://www.elastic.co/guide/en/elasticsearch/reference/current/size-your-shards.html)
  - Contact AppsCode team
- #### ElasticsearchClusterRed
  - A red status indicates that one or more indices do not have allocated primary shards. Find the cause of non-allocation using [_cluster/allocation](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-allocation-explain.html) API.
  - Increase the [low watermark](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-cluster.html#disk-based-shard-allocation) for disk usage to keep the cluster running
- #### ElasticsearchClusterYellow
  - Yellow status indicates that one or more of the replica shards on the Elasticsearch cluster are not allocated to a node. Find the cause of non-allocation using [_cluster/allocation](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-allocation-explain.html) API.
  - There are several reasons why a yellow status can be perfectly normal, and in many cases Elasticsearch will recover to green by itself. If cluster is not recovering to green then,
    - If you have a single node standalone cluster, then it is an expected behaviour as Elasticsearch will never assign a replica to the same node as the primary shard. If you still want your cluster status to be green change the number of replicas on each index to be 0.
    - Try Reconfiguring the [disk based shard allocation settings](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-cluster.html#disk-based-shard-allocation)

- #### ElasticsearchHealthyNodes
  - Elasticsearch cluster should have at least 3 healthy nodes. Check cluster health with [_cluster/health](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html#cluster-health-api-example) API.
- #### ElasticsearchHealthyDataNodes
  - Elasticsearch cluster should have at least 3 healthy data nodes. Check cluster health with [_cluster/health](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-health.html#cluster-health-api-example) API.
- #### ElasticsearchRelocatingShards
  - Shard allocation is the process of allocating shards to nodes. This can happen during initial recovery, replica allocation, rebalancing, or when nodes are added or removed
- #### ElasticsearchInitializingShards
  - Indicates that the shards are getting initialized. Use [_cat/recovery](https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-recovery.html) or [index/_recovery](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-recovery.html) API to view information about ongoing and completed shard recoveries.
- #### ElasticsearchUnassignedShards
  - Indicated that Elasticsearch has unassigned shards. View the unassigned shards with [_cat/shards](https://www.elastic.co/guide/en/elasticsearch/reference/current/diagnose-unassigned-shards.html#diagnose-unassigned-shards) API.
- #### ElasticsearchPendingTasks
  - Cluster pending tasks are updates to the cluster state which may have been initiated directly by a user or by the cluster itself. View the pending cluster tasks using [_cluster/pending_tasks](https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-pending.html#cluster-pending) API.
- #### ElasticsearchNoNewDocuments10M
  - Indicates the cluster haven't received any new documents for the past 10 minutes!

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `ElasticsearchOpsRequest` is ongoing on same database, Wait until it completes.
  - If some nodes of the Elasticsearch Cluster are not `Up`, Try restarting those nodes one at a time.
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `ElasticsearchOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team
- #### OpsRequestFailed
  - Contact AppsCode team

### Stash Alerts
- #### BackupSessionFailed
  - [Describe the BackupSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-backupsession)
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - [View the Backup Job/Sidecar log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-backup-jobsidecar-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### RestoreSessionFailed
  - [Describe the RestoreSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-restoresession)
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - [View the Restore Job/Init-Container log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-restore-jobinit-container-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### NoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - [Describe the BackupConfiguration](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#backupconfiguration-notready)
  - Check the conditions of BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### RepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is true
  - Contact AppsCode team
- #### RepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
  - Update RetentionPolicy of the BackupConfiguration to free up storage
- #### BackupSessionPeriodTooLong | RestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is true
  - Check the `Elasticsearch` CRs status
  - Contact AppsCode team

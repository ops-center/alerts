## TODOs on ZooKeeper Critical Alerts

### Database Alerts

- #### ZooKeeperDown
  - Describe the `ZooKeeper` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team
- #### ZooKeeperTooManyNodes
  - Describe the ZooKeeper CR and check current znode count
  - Investigate which applications are creating excessive znodes — check for stale or unnecessary znodes
  - Perform cleanup if needed, or tune heap/memory settings if the load is legitimate
  - Contact AppsCode team
- #### ZooKeeperTooBigMemory
  - Increase ZooKeeper memory/heap size
  - Contact AppsCode team
- #### ZooKeeperTooManyWatch
  - Describe the `ZooKeeper` CR and check watch count. Also, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### ZooKeeperTooManyConnections
  - Increase ZooKeeper `maxClientCnxns` in config
  - Contact AppsCode team
- #### ZooKeeperTooManyOpenFiles
  - Increase ZooKeeper `open_files_limit` in config
  - Contact AppsCode team
- #### ZooKeeperLeaderElection
  - Describe the `ZooKeeper` CR and check pod restarts and logs
  - Verify ZooKeeper ensemble stability — leader election should not happen frequently
  - Contact AppsCode team
- #### ZooKeeperTooLongFsyncTime
  - Describe the `ZooKeeper` CR, check storage performance — slow disk I/O may cause long fsync
  - Contact AppsCode team
- #### ZooKeeperTooLongSnapshotTime
  - Describe the `ZooKeeper` CR, check storage performance and heap usage
  - Contact AppsCode team
- #### ZooKeeperTooHighAvgLatency
  - Describe the `ZooKeeper` CR, check ZooKeeper pod CPU/memory utilization and network latency
  - Contact AppsCode team
- #### ZooKeeperJvmMemoryFilingUp
  - Describe the `ZooKeeper` CR, check for memory leaks or improper GC tuning
  - Contact AppsCode team
- #### DiskUsageHigh
  - Describe the `ZooKeeper` CR, investigate persistent volume usage and clean up old data
  - Contact AppsCode team
- #### DiskAlmostFull
  - Describe the `ZooKeeper` CR, investigate persistent volume usage and clean up old data
  - Contact AppsCode team


### KubeDB Provisioner

- #### AppPhaseNotReady
  - Describe the `ZooKeeper` CR, check the reason from conditions.
  - Check zookeeper pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### AppPhaseCritical
  - check for `zookeeper` container's logs on the node that is not connected with `Primary` node.
  - If some nodes of the ZooKeeper Cluster are not `Up`, Try restarting those nodes one at a time.
  - Contact AppsCode team if this persists for more than 30 minutes.


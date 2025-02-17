## TODOs on Cassandra Critical Alerts

### Database Alerts

- #### CassandraDown
    - Describe the `Cassandra` CR, check the reason from conditions and try restarting the pods.
    - Contact AppsCode team.
- #### CassandraServiceRespawn
    - Describe the `Cassandra` CR.
    - Contact AppsCode team.
- #### CassandraConnectionTimeouts
    - Describe the `Cassandra` CR, check for network issues or misconfigurations causing the connection timeouts.
    - Investigate Cassandra node status, and verify connectivity to other nodes in the cluster.
    - Contact AppsCode team if the issue persists.

- #### CassandraDroppedMessages
    - Describe the `Cassandra` CR and check the reason for message drops.
    - Investigate if the Cassandra cluster is under heavy load or if there's a network bottleneck.
    - Ensure sufficient resources are allocated for Cassandra nodes.
    - Contact AppsCode team if the issue persists.

- #### CassandraHighReadLatency
    - Describe the `Cassandra` CR and check the read performance.
    - Investigate Cassandra logs for any issues related to disk I/O or read-heavy queries.
    - Check if there are any hot spots in the cluster affecting read performance.
    - Contact AppsCode team if the issue persists.

- #### CassandraHighWriteLatency
    - Describe the `Cassandra` CR and check for write performance issues.
    - Investigate if the system is overloaded with writes or if there are issues with disk I/O.
    - Check if there is sufficient disk space and resources for write operations.
    - Contact AppsCode team if the issue persists.

- #### CassandraMemoryLimit
    - Describe the `Cassandra` CR, and check the current memory usage and allocation for Cassandra nodes.
    - Investigate if there are memory leaks or improper memory limits in the configuration.
    - Adjust the memory settings in the configuration and restart the pods if needed.
    - Contact AppsCode team if the issue persists.
- 
### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact with AppsCode Team.
- #### AppPhaseCritical
    - If any `CassandraOpsRequest` is ongoing on same database, Wait until it completes.
    - Contact with AppsCode team if this persists for more than 30 minutes.

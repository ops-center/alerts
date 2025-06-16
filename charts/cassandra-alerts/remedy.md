## TODOs on Cassandra Critical Alerts

### Database Alerts

- #### CassandraDown
    - Describe the `Cassandra` CR, check the reason from conditions and try restarting the pods.
    - Check the Seed nodes. You may check `CASSANDRA_SEEDS` environment variable of that pod.
    - Contact AppsCode team.
- #### CassandraServiceRespawn
    - Describe the `Cassandra` CR.
    - Check for invalid or conflicting configuration in `/etc/cassandra/cassandra.yaml`, `/etc/cassandra/cassandra-env.sh`, or `/etc/cassandra/jvm*.options` files. You may get errors in `/var/log/cassandra/system.log` in case of invalid or conflicting configurations.
    - Contact AppsCode team.
- #### CassandraConnectionTimeouts
    - Describe the `Cassandra` CR, check for network issues or misconfigurations causing the connection timeouts.
    - Investigate Cassandra node status, and verify connectivity to other nodes in the cluster by using `nodetool status` command.
    - Check CPU and Memory usage.
    - Check thread pool stats by using `nodetool tpstats` and check for pending values of `ReadStage`,`MutationStage`,`CompactionExecutor` and `MemtableFlushWriter`.
    - Contact AppsCode team if the issue persists.

- #### CassandraDroppedMessages
    - Describe the `Cassandra` CR and check the reason for message drops.
    - Investigate if the Cassandra cluster is under heavy load or if there's a network bottleneck.
    - Ensure sufficient resources are allocated for Cassandra nodes.
    - Check for `write_request_timeout`, `read_request_timeout` values from `/etc/cassandra/cassandra.yaml`.
    - Contact AppsCode team if the issue persists.

- #### CassandraHighReadLatency
    - Describe the `Cassandra` CR and check the read performance.
    - Check `ReadStage` pending count from `nodetool tpstats`.
    - Misconfigured of poorly designed `partition key` can cause high read latency. Uneven node load from `nodetool status` may indicate a poorly designed parition key. Check your partition key design.
    - Investigate Cassandra logs for any issues related to disk I/O or read-heavy queries.
    - Check if there are any hot spots in the cluster affecting read performance.
    - Contact AppsCode team if the issue persists.

- #### CassandraHighWriteLatency
    - Describe the `Cassandra` CR and check for write performance issues.
    - Investigate if the system is overloaded with writes or if there are issues with disk I/O.
    - Check `concurrent_writes` value from `/etc/cassandra/cassandra.yaml` file. Increase the value if needed.
    - Monitor and reduce dropped mutations via `nodetool tpstats`.
    - Check if there is sufficient disk space and resources for write operations.
    - Contact AppsCode team if the issue persists.

- #### CassandraMemoryLimit
    - Describe the `Cassandra` CR, and check the current memory usage and allocation for Cassandra nodes.
    - Investigate possible memory leaks or improper memory limits in the configuration. You may check `key cache`, `row cache`, `counter cache` from `nodetool info` to help identify memory leaks.
    - Adjust the memory settings in the configuration and restart the pods if needed.
    - Reduce `memtable_heap_space`, `key_cache_size`, `row_cache_size` from `/etc/cassandra/cassandra.yaml` if necessary.
    - Contact AppsCode team if the issue persists.

### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact with AppsCode Team.
- #### AppPhaseCritical
    - If any `CassandraOpsRequest` is ongoing on same database, Wait until it completes.
    - Contact with AppsCode team if this persists for more than 30 minutes.

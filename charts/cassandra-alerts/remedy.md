## TODOs on Cassandra Critical Alerts

### Database Alerts

- #### CassandraDown
    - Describe the `Cassandra` CR, check the reason from conditions and try restarting the pods.
    - Contact AppsCode team.
- #### CassandraServiceRespawn
    - Describe the `Cassandra` CR.
    - Contact AppsCode team.
- #### CassandraConnectionThrottled
    - Increase `Cassandra` configuration `max_connections`
- #### CassandraConnectionsNoneMinor
    - Try to increase `Cassandra` connections.
- #### CassandraItemsNoneMinor
    - Check `Cassandra` items and try to increase it.
- #### CassandraEvictionsLimit | CassandraMemoryLimit
    - Increase `Cassandra` configuration `memory-limit`.
    - Use [KubeDB](https://kubedb.com/docs/latest/guides/cassandra/) custom configuration
    - Monitoring the cache `Hit Ratios` will also help.

### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact with AppsCode Team.
- #### AppPhaseCritical
    - If any `CassandraOpsRequest` is ongoing on same database, Wait until it completes.
    - Contact with AppsCode team if this persists for more than 30 minutes.

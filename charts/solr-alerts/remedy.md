## TODOs on singlestore Critical Alerts

### Database Alerts

- #### SolrDownShards
    - try restarting the pods . May happen if data nodes are not running.
    - Contact AppsCode team
- #### SolrRecoveryFailedShards
    - Shards failed to recover. Check zookeeper connection. Solr may not able to find zookeeper.
    - Contact AppsCode team
- #### SolrHighQPS
    - We can scale number of shard per collection or db nodes to handle high QPS.
- #### SolrHighThreadsRunning
    - We can configure jvm memory or thread pool to avoid this warning.
- #### SolrHighHeapSize
    - This is directlly connected to jvm heap. So we have to configure jvm memory to avoid this warning
- #### SolrHighBufferSize
    - Any changes in indices are first made in the buffer. Adjusting parameters like maxDox or maxTime which creating collections may dodge this warning
- #### SolrHighPoolSize
    - Adjust thread pool size in solrconfig.xml may help to avoid this warning. 


### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - If any `SinglestoreOpsRequest` is ongoing on same database, Wait until it completes.
    - Contact AppsCode team if this persists for more than 30 minutes.


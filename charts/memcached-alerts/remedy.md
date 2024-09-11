## TODOs on Memcached Critical Alerts

### Database Alerts

- #### MemcachedDown
    - Describe the `Memcached` CR, check the reason from conditions and try restarting the pods.
    - Contact AppsCode team.
- #### MemcachedServiceRespawn
    - Describe the `Memcached` CR.
    - Contact AppsCode team.
- #### MemcachedConnectionThrottled
    - Increase `Memcached` configuration `max_connections`
- #### MemcachedConnectionsNoneMinor
    - Try to increase `Memcached` connections.
- #### MemcachedItemsNoneMinor
    - Check `Memcached` items and try to increase it.
- #### MemcachedEvictionsLimit | MemcachedMemoryLimit
    - Increase `Memcached` configuration `memory-limit`.
    - Use [KubeDB](https://kubedb.com/docs/latest/guides/memcached/) custom configuration
    - Monitoring the cache `Hit Ratios` will also help.

### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact with AppsCode Team.
- #### AppPhaseCritical
    - If any `MemcachedOpsRequest` is ongoing on same database, Wait until it completes.
    - Contact with AppsCode team if this persists for more than 30 minutes.

## TODOs on ProxySQL Critical Alerts

### Database Alerts

- #### ProxySQLInstanceDown
    - Describe the `ProxySQL` CR, check the reason from conditions and try restarting the pods
    - Contact AppsCode team
- #### ProxySQLServiceDown
    - Describe the `ProxySQL` CR check the conditions
    - Contact AppsCode team
- #### ProxySQLTooManyConnections
    - Increase ProxySQL variable `mysql-max_connections`
- #### ProxySQLHighThreadsRunning
    - Increasing ProxySQL variable `mysql-max_connections` may help.
    - Also try tuning ProxySQL for memory optimization
- #### ProxySQLSlowQueries
    - Look into the `stats_mysql_query_digest` table.
- #### ProxySQLRestarted
    - Check if `ProxySQL` CR is in Ready status
    - Contact AppsCode team if status is not updated.
- #### ProxySQLHighQPS | ProxySQLHighIncomingBytes | ProxySQLHighOutgoingBytes
    - Scale ProxySQL using [KubeDB](https://kubedb.com/docs/latest/guides/proxysql/) Scaling OpsRequest

### ProxySQL Cluster Alerts

- #### ProxySQLCLusterSyncFailure
    - Wait for all pods to be `Up`.
    - Contact AppsCode team

### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - If any `ProxySQLOpsRequest` is ongoing on same proxysql server, Wait until it completes.
    - If some nodes of the ProxySQL Cluster are not `Up`, Try restarting those nodes one at a time.
    - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
    - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
    - If any `ProxySQLSQLOpsRequest` is ongoing on same proxysql server, Wait until it completes.
    - Contact AppsCode team
- #### OpsRequestFailed
    - Contact AppsCode team
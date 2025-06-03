## TODOs on ProxySQL Critical Alerts

### Database Alerts

- #### ProxySQLInstanceDown
    - Describe the `ProxySQL` CR using `kubectl describe proxysql <name> -n <namespace>` and check the Events and Conditions sections for reasons or errors.
    - Check the KubeDB Operator logs for errors
    - Check pod status and logs: 
      - kubectl get pods -n <namespace> -l app=<proxysql-label>  
      - kubectl describe pod <proxysql-pod> -n <namespace>
      - kubectl logs -n <namespace> <proxysql-pod> -c proxysql
      - Look for crash loops, readiness/liveness probe failures, or resource errors and try restarting the pods
    - Verify resource limits (CPU/Memory): Inspect your ProxySQL container’s CPU/Memory settings under the CR or PetSet. If pods cannot start due to insufficient resources, adjust as needed.
    - Contact AppsCode team: If the issue persists, reach out with collected logs and events.
- #### ProxySQLServiceDown
  - Describe the `ProxySQL` CR using `kubectl describe proxysql <name> -n <namespace>` and check the Events and Conditions sections for reasons or errors.
  - Check the KubeDB Operator logs for errors
  - Check pods status and logs:
    - kubectl get pods -n <namespace> -l app=<proxysql-label>
    - kubectl describe pod <proxysql-pod> -n <namespace>
    - kubectl logs -n <namespace> <proxysql-pod> -c proxysql
    - Look for crash loops, readiness/liveness probe failures, or resource errors and try restarting the pods
  - Verify resource limits (CPU/Memory): Inspect your ProxySQL container’s CPU/Memory settings under the CR or PetSet. If pods cannot start due to insufficient resources, adjust as needed.
  - Contact AppsCode team: If the issue persists, reach out with collected logs and events.
- #### ProxySQLTooManyConnections
    - Inspect Current Connection Count
    - Increase ProxySQL variable `mysql-max_connections`
      - Create a ProxySQLOpsRequest of type Reconfigure to increase max connections
      - Scale Horizontally: Create a ProxySQLOpsRequest of type HorizontalScaling
    - Contact AppsCode Team: If reconfiguration and scaling do not resolve, gather metrics from stats tables and share with the team.
- #### ProxySQLSlowQueries
    - Look into the `stats_mysql_query_digest` table.
    - Identify high-frequency or long-running queries.
    - Optimize Slow Queries: 
       - You can think of adding indexes or rewriting queries.
       - Consider caching in ProxySQL if the same queries repeat often.
    - Contact AppsCode Team: If you need advice on ProxySQL query routing or caching strategies.
- #### ProxySQLRestarted
    - Check if `ProxySQL` CR is in Ready status: 
        - `kubectl get proxysql <name> -n <namespace> -o yaml`
        - Verify `.status.phase` is Ready.
    - Review Pod Events/Logs: If pods restarted unexpectedly, check events for OOMKill or crash loops
       - kubectl describe pod <proxysql-pod> -n <namespace>
       - kubectl logs <proxysql-pod> -n <namespace> -c proxysql
    - Contact AppsCode Team: If restarts continue or status does not return to Ready.
- #### ProxySQLHighQPS | ProxySQLHighIncomingBytes | ProxySQLHighOutgoingBytes
    - Measure Current QPS & Traffic: using commands like `SELECT sum(count_star) AS total_queries FROM stats_mysql_query_digest;`
    - Scale ProxySQL Horizontally or Vertically using Scaling OpsRequests
    - Tune Connection Pooling & Caching

### ProxySQL Cluster Alerts

- #### ProxySQLCLusterSyncFailure
    - Wait for all nodes/pods to be `Ready`. 
       - kubectl get pods -l app=<proxysql-label> -n <namespace>
       - Ensure all cluster members show Running and Ready: 1/1.
    - Check Cluster Status via Admin Interface: `SELECT hostname, status FROM proxysql_servers;` 
      - Look for nodes in status other than ONLINE.
    - Restart One Node at a Time: If a specific node is stuck, restart that pod and let it rejoin:
       - `kubectl delete pod <proxysql-node-pod> -n <namespace>`
    - Contact AppsCode Team: If sync never completes.

### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - If any `ProxySQLOpsRequest` is ongoing on same proxysql server, Wait until it completes.
    - If some nodes of the ProxySQL Cluster are not `Up`, Try restarting those nodes one at a time.
    - Contact the AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
    - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
    - If any `ProxySQLSQLOpsRequest` is ongoing on same proxysql server, Wait until it completes.
    - Contact AppsCode team
- #### OpsRequestFailed
    - Contact AppsCode team
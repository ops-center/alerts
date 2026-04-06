## TODOs on HanaDB Critical Alerts

### Database Alerts

- #### HanaDBInstanceDown
  - Check HanaDB pod status and exporter container logs.
  - Verify the stats service and ServiceMonitor are healthy.
  - If the database pod is restarting, inspect HANA trace files and coordinator logs.
- #### HanaDBServiceDown
  - Verify all `up` targets for the HanaDB stats service.
  - Check whether the service selector still points to healthy pods.
- #### HanaDBRestarted
  - Check HanaDB pod events and container logs to understand why the pod restarted.
  - Verify whether the restart was caused by a node disruption, OOM kill, or probe failure.
  - If restarts continue, inspect HANA trace files and coordinator logs.
- #### HanaDBHighCPUUsage
  - Check the HANA host CPU usage and expensive SQL statements.
  - Review recent workload spikes or maintenance operations.
  - Increase available CPU or scale the workload if high usage persists.
- #### HanaDBHighMemoryUsage
  - Check HANA memory consumers and allocation limit settings.
  - Review recent workload spikes or large in-memory operations.
  - Increase available memory or tune the workload if memory pressure persists.
- #### HanaDBReplicationNotActive
  - Check `hdbnsutil -sr_state` and `systemReplicationStatus.py` on the cluster members.
  - Inspect the coordinator logs and verify primary/secondary connectivity.
  - Confirm the secondary site is active and replication status is `ACTIVE`.
- #### DiskUsageHigh / DiskAlmostFull
  - Check PVC usage and cleanup old trace or backup files if appropriate.
  - Expand storage if the volume is expected to continue growing.

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Describe the HanaDB object and inspect the conditions.
  - Check operator logs and the HANA pod logs.
- #### AppPhaseCritical
  - Check system replication health and coordinator logs.
  - Identify the unhealthy member and inspect its HANA trace/log output.

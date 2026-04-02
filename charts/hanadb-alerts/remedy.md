## TODOs on HanaDB Critical Alerts

### Database Alerts

- #### HanaDBInstanceDown
  - Check HanaDB pod status and exporter container logs.
  - Verify the stats service and ServiceMonitor are healthy.
  - If the database pod is restarting, inspect HANA trace files and coordinator logs.
- #### HanaDBServiceDown
  - Verify all `up` targets for the HanaDB stats service.
  - Check whether the service selector still points to healthy pods.
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

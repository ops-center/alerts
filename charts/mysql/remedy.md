## TODOs on MySQL Critical Alerts

### Database Alerts

- #### MySQLInstanceDown
  - Describe the `MySQL` CR and check the reason and try restarting the pods
  - Contact AppsCode team
- #### MySQLServiceDown
  - Describe the `MySQL` CR and Try restarting all the pods
  - Contact AppsCode team
- #### MySQLTooManyConnections
  - Increase mysql variable `max_connections`
- #### MySQLHighThreadsRunning
  - Increasing mysql variable `max_connections` may help. 
  - Also try tuning mysql for memory optimization
- #### MySQLSlowQueries
  - Check Slow Query log file here `/var/log/mysql/mysql-slow.log`
- #### MySQLInnoDBLogWaits
  - Reason for alert: [Innodb_log_waits](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html#statvar_Innodb_log_waits)
  - Try reconfiguring [innodb_log_buffer_size](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_buffer_size)
- #### MySQLRestarted
  - Check if `MySQL CR` is in Ready status
  - Contact AppsCode team if status is not updated.
- #### MySQLHighQPS | MySQLHighIncomingBytes | MySQLHighOutgoingBytes
  - Scale MySQL using [KubeDB](https://kubedb.com/docs/v2022.08.08/guides/mysql/) Scaling OpsRequest
- #### MySQLTooManyOpenFiles
  - Increase mysql variable `open_files_limit`

### Group Replication Alerts

- #### MySQLHighReplicationDelay | MySQLHighReplicationTransportTime | MySQLHighReplicationApplyTime | MySQLReplicationHighTransactionTime
  - Try scale MySQL using [KubeDB](https://kubedb.com/docs/v2022.08.08/guides/mysql/) Scaling OpsRequest
  - Contact AppsCode team

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `MySQLOpsRequest` is ongoing on same database, Wait until it completes.
  - If some nodes of the MySQL group are not `Up`, Try restarting those nodes one at a time.
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `MySQLOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team
- #### OpsRequestFailed
  - Contact AppsCode team

### Stash Alerts
- #### BackupSessionFailed
  - Troubleshoot Stash [Backup Failed](https://stash.run/docs/v2022.07.09/guides/troubleshooting/how-to-troubleshoot/#backup-failed)
  - Check the Job created by `backupsession`
  - Check if the `INTEGRITY` of `repository` is true
  - Contact AppsCode team
- #### RestoreSessionFailed
  - Troubleshoot Stash [Restore Failed](https://stash.run/docs/v2022.07.09/guides/troubleshooting/how-to-troubleshoot/#restore-failed)
  - Check the Job created by `restoresession`
  - Contact AppsCode team
- #### NoBackupSessionForTooLong
  - Check if the `backupconfiguration` is not paused
  - Check the Job created by `backupsession`
  - Contact AppsCode team
- #### RepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is true
  - Contact AppsCode team
- #### RepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
- #### BackupSessionPeriodTooLong | RestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is true
  - Check the `MySQL` CRs status
  - Contact AppsCode team


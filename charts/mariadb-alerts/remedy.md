## TODOs on MariaDB Critical Alerts

### Database Alerts

- #### MariaDBInstanceDown
  - Describe the `MariaDB` CR, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### MariaDBServiceDown
  - Describe the `MariaDB` CR check the conditions
  - Contact AppsCode team
- #### MariaDBTooManyConnections
  - Increase MariaDB variable `max_connections`
- #### MariaDBHighThreadsRunning
  - Increasing MariaDB variable `max_connections` may help. 
  - Also try tuning MariaDB for memory optimization
- #### MariaDBSlowQueries
  - Check Slow Query log file here `/var/lib/mysql/'$pod-name'-slow.log`
- #### MariaDBInnoDBLogWaits
  - Reason for alert: [Innodb_log_waits](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html#statvar_Innodb_log_waits)
  - Try reconfiguring [innodb_log_buffer_size](https://dev.my sql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_buffer_size)
- #### MariaDBRestarted
  - Check if `MariaDB` CR is in Ready status
  - Contact AppsCode team if status is not updated.
- #### MariaDBHighQPS | MariaDBHighIncomingBytes | MariaDBHighOutgoingBytes
  - Scale MariaDB using [KubeDB](https://kubedb.com/docs/latest/guides/mariadb/) Scaling OpsRequest
- #### MariaDBTooManyOpenFiles
  - Increase MariaDB variable `open_files_limit`

### MariaDB Cluster Alerts

- #### GaleraReplicationLatencyTooLong
  - Try scale MariaDB using [KubeDB](https://kubedb.com/docs/latest/guides/mariadb/) Scaling OpsRequest
  - Contact AppsCode team

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `MariaDBOpsRequest` is ongoing on same database, Wait until it completes.
  - If some nodes of the MariaDB Cluster are not `Up`, Try restarting those nodes one at a time.
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
  - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
  - If any `MariaDBSQLOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team
- #### OpsRequestFailed
  - Contact AppsCode team

### Stash Alerts
- #### BackupSessionFailed
  - [Describe the BackupSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-backupsession)
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - [View the Backup Job/Sidecar log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-backup-jobsidecar-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### RestoreSessionFailed
  - [Describe the RestoreSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-restoresession)
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - [View the Restore Job/Init-Container log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-restore-jobinit-container-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### NoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - [Describe the BackupConfiguration](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#backupconfiguration-notready)
  - Check the conditions of BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### RepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is true
  - Contact AppsCode team
- #### RepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
  - Update RetentionPolicy of the BackupConfiguration to free up storage
- #### BackupSessionPeriodTooLong | RestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is true
  - Check the `MariaDB` CRs status
  - Contact AppsCode team

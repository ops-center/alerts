## TODOs on MongoDB Critical Alerts

### Database Alerts

- #### MongoDBDown
  - Describe the `MongoDB` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team
- #### MongodbReplicationLag
  - Check if there is any network issue which is causing the replication lag.
- #### MongodbTooManyConnections
  - Increase `net.maxIncomingConnections` value using [MongoDBOpsRequest](https://kubedb.com/docs/latest/guides/mongodb/reconfigure/overview/)
- #### MongodbVirtualMemoryUsage
  - Increase storage and/or memory using [VerticalScaling OpsRequest](https://kubedb.com/docs/latest/guides/mongodb/scaling/vertical-scaling/overview/) and [VolumeExpansion OpsRequest](https://kubedb.com/docs/latest/guides/mongodb/volume-expansion/overview/)
- #### MongoDBRecurrentMemoryPageFaults
  - Increase Memory using [VerticalScaling OpsRequest](https://kubedb.com/docs/latest/guides/mongodb/scaling/vertical-scaling/overview/)
  - Index your data properly
- #### MongodbNumberCursorsOpen
  - Check for slow queries using `db.currentOp()` command
  - Check if application that is using mongodb is closing cursors properly
- #### MongoDBHighLatency | MongodbCursorsTimeouts | MongoDBRecurrentCursorTimeout | MongoDBHighTicketUtilization
  - Check for slow queries using `db.currentOp()` command
  - Check database pod logs
  - Check for network and disk performance
  - Contact AppsCode team if this issue persists for more than a day

### KubeDB Provisioner

- #### KubeDBMongoDBPhaseNotReady
  - Describe the `MongoDB` CR, check the reason from conditions.
  - Check mongodb pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### KubeDBMongoDBPhaseCritical
  - If any `MongoDBOpsRequest` is progressing for the same database, wait until it completes.
  - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
  - Describe the pods that are not Ready and check their events and conditions. 
  - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### KubeDBMongoDBOpsRequestOnProgress
  - Just a reminder that an Ops Request is currently processing for a database, nothing to worry about it.
- #### KubeDBMongoDBOpsRequestStatusProgressingToLong
  - Some ops request might take long depending on the process
  - Check the KubeDB ops manager operator pod logs
  - Contact AppsCode team
- #### KubeDBMongoDBOpsRequestFailed
  - Describe the MongoDBOpsRequest and check the events and conditions.
  - Check the KubeDB ops manager operator pod logs.
  - Contact AppsCode team.

### Stash Alerts

- #### MongoDBStashBackupSessionFailed
  - [Describe the BackupSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-backupsession)
  - Check the conditions in the BackupSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the BackupSession
  - [View the Backup Job/Sidecar log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-backup-jobsidecar-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### MongoDBStashRestoreSessionFailed
  - [Describe the RestoreSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-restoresession)
  - Check the conditions in the RestoreSession
  - Check the reasons of the `false` conditions (if any)
  - Check the events of the RestoreSession
  - [View the Restore Job/Init-Container log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-restore-jobinit-container-log)
  - Check if the `INTEGRITY` of Repository is true
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### MongoDBStashNoBackupSessionForTooLong
  - Check if the BackupConfiguration is not `Paused`
  - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
  - [Describe the BackupConfiguration](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#backupconfiguration-notready)
  - Check the conditions of BackupConfiguration
  - Check the reasons of the `false` conditions (if any)
  - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
  - Contact AppsCode team
- #### MongoDBStashRepositoryCorrupted
  - Check if the `INTEGRITY` of `repository` is true
  - Contact AppsCode team
- #### MongoDBStashRepositoryStorageRunningLow
  - Increase the volume size of `repository` backend
  - Update RetentionPolicy of the BackupConfiguration to free up storage
- #### MongoDBStashBackupSessionPeriodTooLong | MongoDBStashRestoreSessionPeriodTooLong
  - Check if the `INTEGRITY` of `repository` is true
  - Check the `MongoDB` CRs status
  - Contact AppsCode team

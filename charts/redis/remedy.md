## TODOs on Redis Critical Alerts

### Database Alerts

- #### RedisDown
    - Describe the `Redis` CR, check the reason from conditions and try restarting the pods
    - Contact AppsCode team
- #### RedisMissingMaster
  - Describe the `Redis` CR, check the reason from conditions and try restarting the `replica` pods
  - Contact AppsCode team
- #### RedisTooManyMasters
  - Contact AppsCode team
- #### RedisDisconnectedSlaves
  - Describe the `Redis` CR, check the reason from conditions and try restarting the `replica` pods
  - Contact AppsCode team
- #### RedisTooManyConnections
    - Increase Redis variable `max_connections`
- #### RedisRejectedConnections
    - Check if `Redis` CR is in `Ready` state. If it's `Critical` state then figure out which node is not Running and Restart it. 
    - If it's a `Redis Shard Cluster` then make sure the client connecting in `Cluster` mode
    - Contact AppsCode team
### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - If any `RedisOpsRequest` is ongoing on same database, Wait until it completes.
    - check for `redis` container's logs on the node that is not connected with `Primary` node.
    - If some nodes of the Redis Cluster are not `Up`, Try restarting those nodes one at a time.
    - Contact AppsCode team if this persists for more than 30 minutes.

### KubeDB OpsManager

- #### OpsRequestOnProgress
    - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
    - If any `RedisSQLOpsRequest` is ongoing on same database, Wait until it completes. For `Version Upgrade` or some other `OpsRequest` It may take longer then expectation.
    - Contact AppsCode team
- #### OpsRequestFailed
    - Describe the OpsRequest and Check the conditions in it
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
    - Check the `Redis` CRs status
    - Contact AppsCode team
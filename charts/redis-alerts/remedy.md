## TODOs on Redis Critical Alerts

### Database Alerts

- #### RedisDown
    - Describe the `Redis` CR, check the reason from conditions, and try restarting the pods
    - **Example**: Check the logs for errors
    - Contact AppsCode team
- #### RedisMissingMaster
    - Describe the `Redis` CR, check the reason from conditions, and try restarting the `replica` pods
    - **Example**: Ensure the primary node is reachable
    - Contact AppsCode team
- #### RedisTooManyMasters
    - Verify network connectivity between nodes
    - Contact AppsCode team
- #### RedisDisconnectedSlaves
    - Describe the `Redis` CR, check the reason from conditions, and try restarting the `replica/slave` pods
    - **Example**: Verify network connectivity between nodes
    - Contact AppsCode team
- #### RedisTooManyConnections
    - Increase Redis variable `maxclients` using `RedisOpsRequest` type `Reconfigure`
    - Contact AppsCode team

### KubeDB Provisioner

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - If any `RedisOpsRequest` is ongoing on the same database, wait until it completes
    - Check for `redis` container's logs on the node that is not connected with the `Primary` node
    - If some nodes of the Redis Cluster are not `Up`, try restarting those nodes one at a time
    - Contact AppsCode team if this persists for more than 30 minutes

### KubeDB OpsManager

- #### OpsRequestOnProgress
    - Just a reminder, nothing to worry about
- #### OpsRequestStatusProgressingToLong
    - If any `RedisSQLOpsRequest` is ongoing on the same database, wait until it completes For `Version Upgrade` or some other `OpsRequest`, it may take longer than expected
    - Contact AppsCode team
- #### OpsRequestFailed
    - Describe the OpsRequest and check the conditions in it
    - Contact AppsCode team

### Stash Alerts
- #### BackupSessionFailed
    - [Describe the BackupSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-backupsession)
    - Check the conditions in the BackupSession
    - Check the reasons for the `false` conditions (if any)
    - Check the events of the BackupSession
    - [View the Backup Job/Sidecar log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-backup-jobsidecar-log)
    - Check if the `INTEGRITY` of the Repository is true
    - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
    - Contact AppsCode team
- #### RestoreSessionFailed
    - [Describe the RestoreSession](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#describe-the-restoresession)
    - Check the conditions in the RestoreSession
    - Check the reasons for the `false` conditions (if any)
    - Check the events of the RestoreSession
    - [View the Restore Job/Init-Container log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#view-restore-jobinit-container-log)
    - Check if the `INTEGRITY` of the Repository is true
    - [Check the Stash operator log](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#check-stash-operator-log)
    - Contact AppsCode team
- #### NoBackupSessionForTooLong
    - Check if the BackupConfiguration is not `Paused`
    - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
    - [Describe the BackupConfiguration](https://stash.run/docs/latest/guides/troubleshooting/how-to-troubleshoot/#backupconfiguration-notready)
    - Check the conditions of BackupConfiguration
    - Check the reasons for the `false` conditions (if any)
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
  
### KubeStash Alerts
- #### RedisKubeStashBackupSessionFailed
    - Describe the BackupSession
    - Check the conditions in the BackupSession
    - Check the reasons for the `false` conditions (if any)
    - Check the events of the BackupSession
    - View the Backup Job log
    - Check if the `INTEGRITY` of the Repository is `true`
    - Check the KubeStash operator log
    - Contact AppsCode team
- #### RedisKubeStashRestoreSessionFailed
    - Describe the RestoreSession
    - Check the conditions in the RestoreSession
    - Check the reasons for the `false` conditions (if any)
    - Check the events of the RestoreSession
    - View the Restore Job log
    - Check if the `INTEGRITY` of the Repository is `true`
    - Check the KubeStash operator log
    - Contact AppsCode team
- #### RedisKubeStashNoBackupSessionForTooLong
    - Check if the BackupConfiguration is not `Paused`
    - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
    - Describe the BackupConfiguration
    - Check the conditions of BackupConfiguration
    - Check the reasons for the `false` conditions (if any)
    - Check the KubeStash operator log
    - Contact AppsCode team
- #### RedisKubeStashRepositoryCorrupted
    - Check if the `INTEGRITY` of `repository` is `true`
    - Contact AppsCode team
- #### RedisKubeStashRepositoryStorageRunningLow
    - Increase the volume size of `repository` backend
    - Update RetentionPolicy to free up storage
- #### RedisKubeStashBackupSessionPeriodTooLong | RedisKubeStashRestoreSessionPeriodTooLong
    - Check if the `INTEGRITY` of `repository` is `true`
    - Check the `Redis` CRs status
    - Contact AppsCode team

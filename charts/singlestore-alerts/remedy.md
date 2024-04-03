## TODOs on singlestore Critical Alerts

### Database Alerts

- #### SinglestoreInstanceDown
  - Describe the `Singlestore` CR, check the reason from conditions and try restarting the pods
  - Contact AppsCode team
- #### SinglestoreServiceDown
  - Describe the `Singlestore` CR and Try restarting all the pods
  - Contact AppsCode team
- #### SinglestoreTooManyConnections
  - Increase singlestore variable `max_connections`
- #### SinglestoreHighThreadsRunning
  - Increasing singlestore variable `max_connections` may help. 
  - Also try tuning singlestore for memory optimization
- #### SinglestoreRestarted
  - Check if `Singlestore` CR is in Ready status
  - Contact AppsCode team if status is not updated.
- #### SinglestoreHighQPS | SinglestoreHighIncomingBytes | SinglestoreHighOutgoingBytes
  - Scale Singlestore using [KubeDB](https://kubedb.com/docs/latest/guides/singlestore/) Scaling OpsRequest


### KubeDB Provisioner

- #### AppPhaseNotReady
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `SinglestoreOpsRequest` is ongoing on same database, Wait until it completes.
  - Contact AppsCode team if this persists for more than 30 minutes.


### KubeStash Alerts

- #### SinglestoreKubeStashBackupSessionFailed
    - Describe the BackupSession
    - Check the conditions in the BackupSession
    - Check the reasons of the `false` conditions (if any)
    - Check the events of the BackupSession
    - View the Backup Job log
    - Check if the `INTEGRITY` of Repository is `true`
    - Check the KubeStash operator log
    - Contact AppsCode team
- #### SinglestoreKubeStashRestoreSessionFailed
    - Describe the RestoreSession
    - Check the conditions in the RestoreSession
    - Check the reasons of the `false` conditions (if any)
    - Check the events of the RestoreSession
    - View the Restore Job log
    - Check if the `INTEGRITY` of Repository is `true`
    - Check the KubeStash operator log
    - Contact AppsCode team
- #### SinglestoreKubeStashNoBackupSessionForTooLong
    - Check if the BackupConfiguration is not `Paused`
    - Check if the BackupConfiguration is in `Not Ready` or `Invalid` Phase
    - Describe the BackupConfiguration
    - Check the conditions of BackupConfiguration
    - Check the reasons of the `false` conditions (if any)
    - Check the KubeStash operator log
    - Contact AppsCode team
- #### SinglestoreKubeStashRepositoryCorrupted
    - Check if the `INTEGRITY` of `repository` is `true`
    - Contact AppsCode team
- #### SinglestoreKubeStashRepositoryStorageRunningLow
    - Increase the volume size of `repository` backend
    - Update RetentionPolicy to free up storage
- #### SinglestoreKubeStashBackupSessionPeriodTooLong | SinglestoreKubeStashRestoreSessionPeriodTooLong
    - Check if the `INTEGRITY` of `repository` is `true`
    - Check the `Singlestore` CRs status
    - Contact AppsCode team



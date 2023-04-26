## TODOs on Vault Critical Alerts

### Vault Alerts

- #### VaultSealed
    - Describe the `VaultServer` CR, check the reason from conditions and try restarting the pods
    - Checkout Logs of Unsealer container of VaultServer Pod, see if anything unusual.
    - Contact AppsCode team
- #### VaultAutoPilotNodeUnhealthy
    - Debug Vault Integrated Storage from [here](https://developer.hashicorp.com/vault/tutorials/raft/raft-autopilot)
    - Try Restarting the Pods one by one
    - Contact AppsCode team
- #### VaultLeadershipLoss
    - Try Restarting the Pods one by one.
    - Contact AppsCode team
- #### VaultLeadershipStepsDowns
    - Try Restarting the Pods one by one.
    - Contact AppsCode team
- #### VaultLeadershipSetupFailures
    - Try Restarting the Pods one by one
    - Contact AppsCode team
- #### VaultRequestFailures
    - See the logs of VaultServer Pods
    - You might need more resources on Vault Nodes
    - Contact AppsCode team
- #### VaultResponseFailures
    - See the logs of VaultServer Pods
    - You might need more resources on Vault Nodes
    - Contact AppsCode team
- #### VaultTooManyInfinityTokens
    - Revoke the unused Tokens
    - Check if Any instance creating more tokens than necessary.

### KubeVault Operator Alerts

- #### AppPhaseNotReady
    - Contact AppsCode team
- #### AppPhaseCritical
    - If any `VaultOpsRequest` is ongoing on same database, Wait until it completes.
    - Check for `vault` container's logs on the node to see if anything unusual.
    - If some nodes of the Vault Cluster are not `Up`, Try restarting those nodes one at a time.
    - Contact AppsCode team if this persists for more than 30 minutes.

### OpsManager Alerts

- #### OpsRequestOnProgress
    - Just a reminder, nothing to worry about.
- #### OpsRequestStatusProgressingToLong
    - If any `VaultOpsRequest` is ongoing on same Vault server, Wait until it completes. For `Version Upgrade` or some other `OpsRequest` It may take longer then expectation.
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
    - Check the `VaultServer` CRs status
    - Contact AppsCode team
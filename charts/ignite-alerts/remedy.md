## TODOs on Ignite Critical Alerts

### Database Alerts

- #### IgniteDown:
    - Describe the `Ignite` CR, check the reason from conditions and try restarting the pods
    - Check database pod logs
    - Contact AppsCode team
- #### IgniteClusterNoBaselineNode:
    - Check the baseline topology of cluster
    - Add nodes to baseline topology.
- #### IgniteRestarted:
    - Check the Ignite CR.
    - Check the logs of nodes why it restarted.
- #### IgniteHighCPULoad:
    - Increase the CPU limit of Ignite.
- #### IgniteHighHeapMemoryUsed:
    - Increase heap memory limit of Ignite.
- #### IgniteHighDataregionOffHeapUsed:
    - Increase the dataregion off heap limit of Ignite.
- #### IgniteJVMPausesTotalDuration:
    - Check the current phase of Ignite.
    - Check whether any pods got restarted.
  

### KubeDB Provisioner

- #### KubeDBIgnitePhaseNotReady
    - Describe the `Ignite` CR, check the reason from conditions.
    - Check ignite pod logs
    - Check KubeDB provisioner logs
    - Contact AppsCode team
- #### KubeDBIgnitePhaseCritical
    - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
    - Describe the pods that are not Ready and check their events and conditions.
    - Contact AppsCode team if this persists for more than 30 minutes.
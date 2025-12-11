## TODOs on Qdrant Critical Alerts

### Database Alerts

- #### QdrantDown:
    - Describe the `Qdrant` CR, check the reason from conditions and try restarting the pods
    - Check database pod logs
    - Contact AppsCode team
- #### QdrantRestarted:
    - Check the Qdrant CR.
    - Check the logs of nodes why it restarted.
- #### QdrantHighCPULoad:
    - Increase the CPU limit of Qdrant.

### KubeDB Provisioner

- #### KubeDBQdrantPhaseNotReady
    - Describe the `Qdrant` CR, check the reason from conditions.
    - Check qdrant pod logs
    - Check KubeDB provisioner logs
    - Contact AppsCode team
- #### KubeDBQdrantPhaseCritical
    - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
    - Describe the pods that are not Ready and check their events and conditions.
    - Contact AppsCode team if this persists for more than 30 minutes.
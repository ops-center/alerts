## TODOs on Qdrant Critical Alerts

### Database Alerts

- #### QdrantInstanceDown
  - Describe the `Qdrant` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team
- #### QdrantPhaseCritical
  - Describe the `Qdrant` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team if this persists for more than 30 minutes
- #### QdrantRestarted
  - Check database pod logs to understand why the restart occurred
  - Describe the `Qdrant` CR and check for any issues in conditions
  - Contact AppsCode team if restarts are frequent
- #### QdrantHighCPUUsage
  - Check for runaway queries or high load
  - Scale horizontally using [QdrantOpsRequest](https://kubedb.com/docs/latest/guides/qdrant/scaling/horizontal-scaling/overview/) type `HorizontalScaling`
  - Increase CPU limits using [VerticalScaling OpsRequest](https://kubedb.com/docs/latest/guides/qdrant/scaling/vertical-scaling/overview/)
  - Contact AppsCode team if the issue persists
- #### QdrantHighMemoryUsage
  - Check for memory leaks or high memory consumption
  - Increase memory limits using [VerticalScaling OpsRequest](https://kubedb.com/docs/latest/guides/qdrant/scaling/vertical-scaling/overview/)
  - Contact AppsCode team if the issue persists
- #### QdrantHighPendingOperations
  - Check for slow operations or queue buildup
  - Investigate network issues between Qdrant nodes
  - Check cluster health and node status
  - Contact AppsCode team if the issue persists
- #### QdrantGrpcResponsesFailHigh
  - Check gRPC service logs for errors
  - Verify network connectivity between Qdrant nodes
  - Check for resource constraints (CPU/Memory)
  - Contact AppsCode team if the issue persists
- #### QdrantRestResponsesFailHigh
  - Check REST API service logs for errors
  - Verify network connectivity and cluster health
  - Check for resource constraints (CPU/Memory)
  - Contact AppsCode team if the issue persists
- #### DiskUsageHigh
  - Expand storage using [VolumeExpansion OpsRequest](https://kubedb.com/docs/latest/guides/qdrant/volume-expansion/overview/)
  - Clean up old or unnecessary data
  - Contact AppsCode team
- #### DiskAlmostFull
  - Expand storage using [VolumeExpansion OpsRequest](https://kubedb.com/docs/latest/guides/qdrant/volume-expansion/overview/) immediately
  - Clean up old or unnecessary data to free up space
  - Contact AppsCode team

### KubeDB Provisioner

- #### AppPhaseNotReady
  - Describe the `Qdrant` CR, check the reason from conditions
  - Check Qdrant pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### AppPhaseCritical
  - If any `QdrantOpsRequest` is progressing for the same database, wait until it completes
  - If some pods of the database are not `Ready`, try restarting those pods one at a time
  - Describe the pods that are not Ready and check their events and conditions
  - Contact AppsCode team if this persists for more than 30 minutes

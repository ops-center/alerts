## TODOs on Hazelcast Critical Alerts

### Database Alerts

- #### HazelcastPartitionCountExceed:
  - Number of partition exceeds mean we need more pods
  - Ensure that, all the pods are running.
- #### HazelcastHighHeapPercentage:
  - Happens when java memory exceeds
  - Need to increase heap memory.
- #### HazelcastHighMemoryUsage:
  - Happens if 80% of memory is occupied
  - In crease pod memory. Vertical scaling os request can fix the issue
- #### HazelcastHighPhysicalMemoryUsage:
  - Refers to excessive consumption of RAM 
  - Need to increase RAM
- #### HazelcastHighLatency :
  - Get operations taking too much time
  - It can happen because of excessive traffic.
  - In creasing pod can fix the issue

- #### HazelcastDown
  - Describe the `Hazelcast` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team

### KubeDB Provisioner

- #### KubeDBHazelcastPhaseNotReady
  - Describe the `Hazelcast` CR, check the reason from conditions.
  - Check rabbitmq pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### KubeDBHazelcastPhaseCritical
  - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
  - Describe the pods that are not Ready and check their events and conditions. 
  - Contact AppsCode team if this persists for more than 30 minutes.
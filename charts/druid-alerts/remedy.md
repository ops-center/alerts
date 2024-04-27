## TODOs on Druid Critical Alerts

### Database Alerts

- #### DruidDown:
  - Describe the `RabbitMQ` CR, check the reason from conditions and try restarting the pods
  - Check database pod logs
  - Contact AppsCode team
- #### ZKDisconnected:
  - Try restarting ZooKeeper
  - Check for any mistakes in the configurations of zookeeper connections in `common.runtime.properties` file.
  - Monitor for dead letter exchanges to handle unroutable messages appropriately.
  - If using KubeDB Managed ZooKeeper, then, Contact AppsCode team
- #### HighSegmentUsage:
  - Consider adding more disk space to your Druid cluster
  - Consider adjusting the segment granularity settings to reduce the size of individual segments.
- #### HighJVMMemoryUsage:
  - Consider increasing the JVM heap size allocated to Druid.
  - Review and optimize the configuration settings of Apache Druid to reduce memory consumption.
### KubeDB Provisioner

- #### KubeDBDruidPhaseNotReady
  - Describe the `Druid` CR, check the reason from conditions.
  - Check rabbitmq pod logs
  - Check KubeDB provisioner logs
  - Contact AppsCode team
- #### KubeDBDruidPhaseCritical
  - If some pods of the database are not `Ready`, Try restarting those pods one at a time.
  - Describe the pods that are not Ready and check their events and conditions. 
  - Contact AppsCode team if this persists for more than 30 minutes.
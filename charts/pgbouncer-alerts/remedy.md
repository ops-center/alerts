### Database Alerts

- #### PgBouncerDown
    - Describe the `PgBouncer` CR, check the reason from conditions and try restarting the pods
    - Check pgbouncer pod logs
    - Contact AppsCode team
- #### PgBouncerExporterLastScrapeError
    - Check logs for exporter containers
    - Check if PgBouncer is pingable
- #### PgBouncerTooManyConnections
    - Increase `max_pool` value in `initConfig` field using [PgBouncerOpsRequest](https://kubedb.com/docs/latest/guides/pgbouncer/reconfigure/overview/)
- #### pgbouncerLogPoolerError
    - Check logs for pgbouncer pods

### KubeDB Provisioner

- #### KubeDBPgBouncerPhaseNotReady
    - Describe the `PgBouncer` CR, check the reason from conditions.
    - Check pgbouncer pod logs
    - Check KubeDB provisioner logs
    - Contact AppsCode team
- #### KubeDBPgBouncerPhaseCritical
    - If any `PgBouncerDBOpsRequest` is progressing for the same pgbouncer, wait until it completes.
    - If some pods of the pgbouncer are not `Ready`, Try restarting those pods one at a time.
    - Describe the pods that are not Ready and check their events and conditions.
    - Contact AppsCode team if this persists for more than 30 minutes.

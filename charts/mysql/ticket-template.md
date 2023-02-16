This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern
while using. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have MySQL database in namespace demo.Please refer to [KubeDB Docs](https://kubedb.com/docs/v2023.01.31/guides/mysql/) for more about KubeDB.
```bash
Kubectl get mysql <ns>     // wiil list all the database in a namesapce  
$ kubectl get mysql -n demo
  NAME    VERSION   STATUS   AGE
  mysql   8.0.31    Ready    6h51m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``critical`` ``NotReady``

**Ready:** KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

**Provisioning:** Usually, the Database Phase is set to Provisioning to bootstrapping for the first time. If you find the database phase is stuck in the provisional state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A possible solution is to describe the database phase, check the configuration, operator, and pod logs and find the reason.

you can contact to AppsCode with the following things on hand.
- MySQL object:
    ```bash
    kubectl get mysql -n <ns> <mysql-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the MySQL object 
    ```bash
      kubectl describe mysql -n <ns> <mysql-ojbect-Name> 
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <mysql-ojbect-Name>
    ```
- Describe the pods: if there are multiple pods describe all of them 
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Check the stateful and secret is created
    ```bash
    kubectl get secret -n <ns> | grep <mysql-object-name>
    ```
- Pod logs:  if there are multiple pods, log all of them. stand alone mysql doesn't contain the `mysql-coordinator` container.
    ```bash
    kubectl logs -n <ns> <pod-name> -c mysql
    kubectl logs -n <ns> <pod-name> -c mysql-coordinator  
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmaneger-pod-name>
    ```

**Critical:** Database Phase Critical means some of the database server/pods is not in the cluster and not working properly with the database cluster. 
The reasons behind this could be someDatabase left the cluster maybe for a restart or replication errors or unexpected kills. 
A Possible solution for this could be to find out the servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server. 
You can contact to AppsCode with the following things on hand

- MySQL object:
    ```bash
    kubectl get mysql -n <ns> <mysql-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the MySQL object
    ```bash
      kubectl describe mysql -n <ns> <mysql-ojbect-Name> 
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <mysql-ojbect-Name>
    ```
- Describe the pods: if there are multiple pods describe all of them
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Pod logs:  if there are multiple pods, log all of them. stand alone mysql doesn't contain the `mysql-coordinator` container.
    ```bash
    kubectl logs -n <ns> <pod-name> -c mysql
    kubectl logs -n <ns> <pod-name> -c mysql-coordinator  
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmaneger-pod-name>
    ```

**NotReady:** Database Phase Not Ready means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous. A possible solution is to identify what exactly happened.
Checking the logs from operator and pod containers describing the database object and pods. Restarting the pod might work. Before forcing a cluster fail-over and recover. 
There might be a need for further human intervention to know what exactly happened and what will be the best way to resolve it.
In that case please contact AppsCode with the following information.

- MySQL object:
    ```bash
    kubectl get mysql -n <ns> <mysql-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the MySQL object
    ```bash
      kubectl describe mysql -n <ns> <mysql-ojbect-Name> 
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <mysql-ojbect-Name>
    ```
- Describe the pods: if there are multiple pods describe all of them
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Pod logs:  if there are multiple pods, log all of them. stand alone mysql doesn't contain the `mysql-coordinator` container.
    ```bash
    kubectl logs -n <ns> <pod-name> -c mysql
    kubectl logs -n <ns> <pod-name> -c mysql-coordinator  
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmaneger-pod-name>
    ```
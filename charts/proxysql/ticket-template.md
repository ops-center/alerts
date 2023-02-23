This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed servers follow similar pattern. To troubleshoot and find what to do let's look into the following doc.

> Please refer to [KubeDB Docs](https://kubedb.com/docs/latest/guides/proxysql/) for more about KubeDB ProxySQL.

Let's say you have a KubeDB provisioned ProxySQL running in namespace demo.
```bash
kubectl get proxysql -n <ns>     // will list all the proxysql in a namesapce  
$ kubectl get proxysql -n demo
  NAME            VERSION         STATUS   AGE
  demo-proxysql   2.3.2-debian    Ready    6h51m
```
There are four different phases for a KubeDB managed ProxySQL object.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

**Ready:** KubeDB Managed ProxySQL phase becomes Ready when every proxysql server is connected with the cluster and working properly. There's nothing to worry about.

**Provisioning:** Usually, the ProxySQL phase is set to `Provisioning` while bootstrapping for the first time. If you find the proxysql phase is stuck in the provisioning state for a long time,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the ProxySQL object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- Get the ProxySQL object:
    ```bash
    kubectl get proxysql -n <ns> <proxysql-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml 
    ```
- Describe the ProxySQL object:
    ```bash
      kubectl describe proxysql -n <ns> <proxysql-object-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <proxysql-object-Name>
    ```
- Describe the pods. If there are multiple pods describe all of them:
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Check the related statefulset and secrets:
    ```bash
    kubectl get secret -n <ns> | grep <proxysql-object-name>
    ```
- Get the pod logs. If there are multiple pods, log all of them:
    ```bash
    kubectl logs -n <ns> <pod-name>
    ```
- Get the operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```

**Critical:** ProxySQL Phase  Critical means some server/pods is not in the cluster or failing synchronization with the database cluster
The reasons could be that some proxysql server left the cluster for a restart. Or replication errors or unexpected kills or an ops-request might have caused a reload.
To resolve this, we need to observe the whole situation. For this purpose we should get the log of troubling pod/server, describe it,  by checking the logs, describing the proxysql object, or maybe querying in the working proxysql server.

You can contact to AppsCode with the following things attached,

- Get the ProxySQL object:
    ```bash
    kubectl get proxysql -n <ns> <proxysql-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the ProxySQL object:
    ```bash
      kubectl describe proxysql -n <ns> <proxysql-object-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <proxysql-object-Name>
    ```
- Describe the pods. If there are multiple pods describe all of them:
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Get the pod logs. if there are multiple pods, log all of them:
    ```bash
    kubectl logs -n <ns> <pod-name>
    ```
- Get the operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```

**NotReady:** ProxySQL Phase Not Ready means none of the proxysql servers are working properly. There are several possible reasons for that. Maybe something is misconfigured,
or the server is Killed, replication errors, or something miscellaneous. To resolve this, first we need to know what exactly happened.
Checking the logs from operator and pod containers, describing the proxysql object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

In that case please contact AppsCode with the following information.

- Get the ProxySQL object:
    ```bash
    kubectl get proxysql -n <ns> <proxysql-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the ProxySQL object:
    ```bash
      kubectl describe proxysql -n <ns> <proxysql-object-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <proxysql-object-Name>
    ```
- Describe the pods, if there are multiple pods describe all of them:
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Get the pod logs, if there are multiple pods get all of them:
    ```bash
    kubectl logs -n <ns> <pod-name>
    ```
- Get the operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
    ```
  
This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed servers follow similar pattern. To troubleshoot and find what to do let's look into the following doc.

> Please refer to [KubeDB Docs](https://kubedb.com/docs/v2023.01.31/guides/elasticsearch/) for more about KubeDB Elasticsearch. This documentation is applicable for both Opensearch and Elasticsearch as KueDB manages both of these databases with the same Controller CRD.

Let's say you have a KubeDB provisioned Elasticsearch running in namespace demo.
```bash
Kubectl get elasticsearch -n <ns>     // wiil list all the elasticsearch in a namesapce  
$ kubectl get elasticsearch -n demo
  NAME                 VERSION        STATUS   AGE
  demo-elasticsearch   xpack-8.5.2    Ready    6h51m
```
There are four different phases for a KubeDB managed Elasticsearch object.
``Ready`` ``Provisioning`` ``critical`` ``NotReady``

Below we have discussed what the phases actually mean and what we recommend to do while contacting us with your issues.

**Ready:** KubeDB Managed Elasticsearch phase becomes Ready when every Elasticsearch node is connected with the cluster and working properly. There's nothing to worry about.

**Provisioning:** Usually, the Elasticsearch phase is set to `Provisioning` while bootstrapping for the first time. If you find the Elasticsearch phase is stuck in the provisional state for a long time, there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
To step ahead for solution, first we need to observe the whole situation. For this purpose we need to describe the Elasticsearch, check the configurations, get the operator and pod logs and find the reason.

You can contact to AppsCode team with the following things attached,
- Get the Elasticsearch object yaml:
    ```bash
    kubectl get es -n <ns> <elasticsearch-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml 
    ```
- Describe the Elasticsearch object:
    ```bash
      kubectl describe es -n <ns> <elasticsearch-ojbect-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <elasticsearch-ojbect-Name>
    ```
- Describe the pods. If there are multiple pods describe all of them:
    ```bash
       kubectl describe pod -n <ns> <podName> 
    ```
- Check the related statefulset and secrets:
    ```bash
    kubectl get secret -n <ns> | grep <elasticsearch-object-name>
    ```
- Get the pod logs. If there are multiple pods, log all of them:
    ```bash
    kubectl logs -n <ns> <pod-name>
    ```
- Get the operator logs:
    ```bash
    kubectl logs -n <kubedb-ns> <provisioner-pod-name>
    kubectl logs -n <kubedb-ns> <opsmaneger-pod-name>
    ```

**Critical:** Elasticsearch Phase Critical means some Elasticsearch node/pod is not joined with the cluster or failing synchronization with it.
The reasons could be that some elasticsearch server left the cluster for a restart. Or replication errors or unexpected kills or an ops-request might have caused a reload.
To step ahead for a solution, first we need to observe the whole situation. For this purpose we should get the log of troubling pod/server, describe it,  by checking the logs, describing the Elasticsearch object, or maybe querying in the working elasticsearch node.

You can contact to AppsCode team with the following things attached,

- Get the Elasticsearch object yaml:
    ```bash
    kubectl get es -n <ns> <elasticsearch-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml 
    ```
- Describe the Elasticsearch object:
    ```bash
      kubectl describe es -n <ns> <elasticsearch-ojbect-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <elasticsearch-ojbect-Name>
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
    kubectl logs -n <kubedb-ns> <opsmaneger-pod-name>
    ```

**NotReady:** Elasticsearch Phase Not Ready means none of the elasticsearch pods are working properly. There are several possible reasons for that. Maybe something is misconfigured, or the server is Killed, replication errors, or something miscellaneous. To resolve this, first we need to know what exactly happened.
Checking the logs from operator and pod containers, describing the Elasticsearch object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

In that case please contact AppsCode team with the following information.

- Get the Elasticsearch object yaml:
    ```bash
    kubectl get es -n <ns> <elasticsearch-object-name> -oyaml
    Kubectl get pod -n <demo> <pod-name> -oyaml
    ```
- Describe the Elasticsearch object:
    ```bash
      kubectl describe es -n <ns> <elasticsearch-ojbect-Name> 
    ```
- Describe the StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <elasticsearch-ojbect-Name>
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
    kubectl logs -n <kubedb-ns> <opsmaneger-pod-name>
    ```
  
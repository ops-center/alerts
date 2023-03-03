This template is designed to contact with AppsCode with troubleshooting and support in general. To troubleshoot and find what to do let's get started with the vaultserver phase.

> Please refer to [KubeVault Docs](https://kubevault.com/docs/latest/guides/) for more about KubeVault.

Let's say you have VaultServer instance in a namespace.
```bash
kubectl get vaultserver -n <ns>  // will list all the vaultserver in a namesapce     
$ kubectl get vaultserver -n demo 
  NAME           VERSION   STATUS   AGE
  sample-vs      6.2.5     Ready    1h11m
```
There are four different vaultserver phase you may see in KubeVault managed VaultServer.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

**Ready:** VaultServer phase becomes Ready when every vault server instance is in the cluster and working properly. There's nothing to worry about.

**Provisioning:** Usually, the VaultServer Phase is set to `Provisioning` while bootstrapping for the first time. If you find the vaultserver phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the VaultServer object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- Get the VaultServer object:
    ```bash
    kubectl get vs -n <ns> <vaultserver-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml
    ```
- Describe the VaultServer object:
    ```bash
      kubectl describe vs -n <ns> <vaultserver-object-name>
    ```
- Describe the underlying StatefulSet object:
    ```bash
      kubectl describe sts -n <ns> <vaultserver-object-name>
    ```
- Describe the pods: If there are multiple pods describe all of them one by one
    ```bash
       kubectl describe pod -n <ns> <pod-name>
    ```
- Check the stateful and secret is created:
    ```bash
    kubectl get secret -n <ns> | grep <vaultserver-object-name>
    ```
- Pod logs: If there are multiple pods, log all of them one by one.
    ```bash
    kubectl logs -n <ns> <pod-name> -c vault
    # If the vaultserver is sealed
    kubectl logs -n <ns> <pod-name> -c vault-unsealer
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubevault-ns> <operator-pod-name>
    ```

**Critical:** VaultServer Phase Critical means some server/pods is not in the cluster or failing synchronization with the cluster.
The reasons could be some VaultServer left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the vaultserver object, or maybe querying in the working server.

You can contact to AppsCode with the following things attached,

- VaultServer object:
    ```bash
    kubectl get vs -n <ns> <vaultserver-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml
    ```
- Describe the VaultServer object
    ```bash
      kubectl describe vs -n <ns> <vaultserver-object-Name>
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <vaultserver-object-name>
    ```
- Describe the pods: if there are multiple pods describe all of them on by one.
    ```bash
       kubectl describe pod -n <ns> <pod-name>
    ```
- Pod logs: If there are multiple pods, log all of them one by one.
    ```bash
    kubectl logs -n <ns> <pod-name> -c vault
    # If the vaultserver is sealed
    kubectl logs -n <ns> <pod-name> -c vault-unsealer
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubevault-ns> <operator-pod-name>
    ```

**NotReady:** VaultServer Phase NotReady means none of the servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the server is killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the VaultServer object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

In that case please contact AppsCode with the following information.

- VaultServer object:
    ```bash
    kubectl get vs -n <ns> <vaultserver-object-name> -oyaml
    Kubectl get pod -n <ns> <pod-name> -oyaml
    ```
- Describe the VaultServer object
    ```bash
      kubectl describe vs -n <ns> <vaultserver-object-Name>
    ```
- Describe the StatefulSet object
    ```bash
      kubectl describe sts -n <ns> <statefulset-name>
    ```
- Describe the pods: if there are multiple pods describe all of them on by one
    ```bash
       kubectl describe pod -n <ns> <podName>
    ```
- Pod logs: If there are multiple pods, log all of them one by one.
    ```bash
    kubectl logs -n <ns> <pod-name> -c vault
    # If the vaultserver is sealed
    kubectl logs -n <ns> <pod-name> -c vault-unsealer
    ```
- Operator logs:
    ```bash
    kubectl logs -n <kubevault-ns> <operator-pod-name>
    ```
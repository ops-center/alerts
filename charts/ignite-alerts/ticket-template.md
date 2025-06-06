This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have ignite database in namespace demo.
```bash
Kubectl get ignite -n <ns>     // will list all the database in a namesapce  
$ kubectl get ignite -n demo
  NAME    VERSION   STATUS   AGE
  rm-sh   2.17.0    Ready    6h51m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

# Phases
## Ready
KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

## Provisioning
Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the ignite object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- [Get the ignite object](#get-ignite)
- [Describe ignite object](#describe-ignite)
- Describe the PetSet object: If there are multiple petsets, describe all of them. To do that, first
  [List petsets under a ignite](#list-petsets-under-a-ignite). Then [Describe petsets](#describe-petsets) for each of the petsets.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a ignite](#list-pods-under-a-ignite). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded ignite, you also need to check the [ignite bootstrap container's logs](#ignite-bootstrap-logs).
- [Operator logs](#operator-logs)

## Critical
Database Phase Critical means some of the database server/pods is not in the cluster or failing synchronization with the database cluster.
The reasons could be some Database left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server.

You can contact to AppsCode with the following things attached,
- [Get the ignite object](#get-ignite)
- [Describe ignite object](#describe-ignite)
- Describe the PetSet object: If there are multiple petsets, describe all of them. To do that, first
  [List petsets under a ignite](#list-petsets-under-a-ignite). Then [Describe petsets](#describe-petsets) for each of the petsets.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a ignite](#list-pods-under-a-ignite). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded ignite, you also need to check the [ignite bootstrap container's logs](#ignite-bootstrap-logs).
- [Operator logs](#operator-logs)

## NotReady
Database Phase NotReady means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the ignite object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

You can contact to AppsCode with the following things attached,
- [Get the ignite object](#get-ignite)
- [Describe ignite object](#describe-ignite)
- Describe the PetSet object: If there are multiple petsets, describe all of them. To do that, first
  [List petsets under a ignite](#list-petsets-under-a-ignite). Then [Describe petsets](#describe-petsets) for each of the petsets.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a ignite](#list-pods-under-a-ignite). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded ignite, you also need to check the [ignite bootstrap container's logs](#ignite-bootstrap-logs).
- [Operator logs](#operator-logs)

# How to ?
## Get ignite
```bash
kubectl get ignite -n <namespace> -oyaml <ignite-object-Name> 
```

## Describe ignite
```bash
kubectl describe ignite -n <namespace> <ignite-object-Name> 
```

## List petsets under a ignite
```bash
kubectl get petsets -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=ignites.kubedb.com,app.kubernetes.io/instance=<ignite-object-Name> 
```

## Describe petsets
```bash
kubectl describe petsets -n <namespace> <petsets-Name> 
```

## Check secret
```bash
# Get the secret name
kubectl get rm -n <namespace> <ignite-object-Name> -o jsonpath="{.spec.authSecret.name}"
# Check secret's existence
kubectl get secrets -n <namespace> <ignite-secret-name>
```

## List pods under a ignite
```bash
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=ignites.kubedb.com,app.kubernetes.io/instance=<ignite-object-Name> 
```

## Describe pods
```bash
kubectl describe pods -n <namespace> <pod-Name> 
```

## Get pod logs
```bash
kubectl logs -n <namespace> <pod-name> -c ignite
```

## Ignite bootstrap logs
```bash
# List ignite pods 
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=ignites.kubedb.com,app.kubernetes.io/instance=<ignite-object-Name> 
# Now check logs
kubectl logs -n <ns> <pod-name> -c ignite
```

## Operator logs
```bash
kubectl logs -n <kubedb-ns> <provisioner-pod-name>
kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
```
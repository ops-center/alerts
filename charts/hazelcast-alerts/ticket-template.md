This template is designed to contact with AppsCode with troubleshooting and support in general. You will find all the KubeDB Managed database follows similar pattern. To troubleshoot and find what to do let's get started with the database phase.

Let's say you have hazelcast database in namespace demo.
```bash
Kubectl get hazelcast -n <ns>     // will list all the database in a namesapce  
$ kubectl get hazelcast -n demo
  NAME    VERSION   STATUS   AGE
  hz-sh   4.4.6     Ready    6h51m
```
There are four different db phase you may see in KubeDB managed Database.
``Ready`` ``Provisioning`` ``Critical`` ``NotReady``

# Phases
## Ready
KubeDB Managed Database phase becomes Ready when every database server is in the cluster and working properly. There's nothing to worry about.

## Provisioning
Usually, the Database Phase is set to `Provisioning` while bootstrapping for the first time. If you find the database phase is stuck in the provisioning state,
there may be some misconfiguration, lack of k8s resources, or miscellaneous issues.
A recommended approach is to describe the hazelcast object, check the configuration, operator, and pod logs and find the reason.

You can contact to AppsCode with the following things attached,
- [Get the hazelcast object](#get-hazelcast)
- [Describe hazelcast object](#describe-hazelcast)
- Describe the StatefulSet object: If there are multiple sts, describe all of them. To do that, first
  [List statefulsets under a hazelcast](#list-statefulsets-under-a-hazelcast). Then [Describe statefulsets](#describe-statefulsets) for each of the sts.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a hazelcast](#list-pods-under-a-hazelcast). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded hazelcast, you also need to check the [hazelcasts bootstrap container's logs](#hazelcasts-bootstrap-logs).
- [Operator logs](#operator-logs)

## Critical
Database Phase Critical means some of the database server/pods is not in the cluster or failing synchronization with the database cluster.
The reasons could be some Database left the cluster maybe for a restart or replication errors or unexpected kills.
To resolve this , we need to  find out which servers/pod that is not in the cluster by checking the logs, describing the database object, or maybe querying in the working database server.

You can contact to AppsCode with the following things attached,
- [Get the hazelcast object](#get-hazelcast)
- [Describe hazelcast object](#describe-hazelcast)
- Describe the StatefulSet object: If there are multiple sts, describe all of them. To do that, first
  [List statefulsets under a hazelcast](#list-statefulsets-under-a-hazelcast). Then [Describe statefulsets](#describe-statefulsets) for each of the sts.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a hazelcast](#list-pods-under-a-hazelcast). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded hazelcast, you also need to check the [hazelcasts bootstrap container's logs](#hazelcasts-bootstrap-logs).
- [Operator logs](#operator-logs)

## NotReady
Database Phase NotReady means none of the database servers are working properly. There could several possible reasons for that, maybe something is misconfigured,
maybe the database server is Killed, Replication errors, or something miscellaneous.
To resolve this, first we need to know what exactly happened. Checking the logs from operator and pod containers, describing the hazelcast object and pods is a recommended way to start debugging. Restarting the pod might sometime solve the issue. But, before forcing a cluster fail-over and recover,
there might be a need for human intervention to know what will be the best way to resolve it.

You can contact to AppsCode with the following things attached,
- [Get the hazelcast object](#get-hazelcast)
- [Describe hazelcast object](#describe-hazelcast)
- Describe the StatefulSet object: If there are multiple sts, describe all of them. To do that, first
  [List statefulsets under a hazelcast](#list-statefulsets-under-a-hazelcast). Then [Describe statefulsets](#describe-statefulsets) for each of the sts.
- Describe the pods: If there are multiple pod, describe all of them. To do that, first
  [List pods under a hazelcast](#list-pods-under-a-hazelcast). Then [Describe pods](#describe-pods) for each of the pod.
- [Check the secret is created](#check-secret)
- [Check Pod logs](#get-pod-logs).  If it is a sharded hazelcast, you also need to check the [hazelcasts bootstrap container's logs](#hazelcasts-bootstrap-logs).
- [Operator logs](#operator-logs)

# How to ? 
## Get hazelcast
```bash
kubectl get hazelcast -n <namespace> -oyaml <hazelcast-object-Name> 
```

## Describe hazelcast
```bash
kubectl describe hazelcast -n <namespace> <hazelcast-object-Name> 
```

## List statefulsets under a hazelcast
```bash
kubectl get sts -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=hazelcasts.kubedb.com,app.kubernetes.io/instance=<hazelcast-object-Name> 
```

## Describe statefulsets
```bash
kubectl describe sts -n <namespace> <sts-Name> 
```

## Check secret
```bash
# Get the secret name
kubectl get hz -n <namespace> <hazelcast-object-Name> -o jsonpath="{.spec.authSecret.name}"
# Check secret's existence
kubectl get secrets -n <namespace> <hazelcast-secret-name>
```

## List pods under a hazelcast
```bash
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=hazelcasts.kubedb.com,app.kubernetes.io/instance=<hazelcast-object-Name> 
```

## Describe pods
```bash
kubectl describe pods -n <namespace> <pod-Name> 
```

## Get pod logs 
```bash
kubectl logs -n <namespace> <pod-name> -c hazelcast
```

## hazelcasts bootstrap logs
```bash
# List hazelcasts pods 
kubectl get pods -n <namespace> -l=app.kubernetes.io/component=database,app.kubernetes.io/managed-by=kubedb.com,app.kubernetes.io/name=hazelcasts.kubedb.com,app.kubernetes.io/instance=<hazelcast-object-Name> 
# Now check logs
kubectl logs -n <ns> <pod-name> -c hazelcast
```

## Operator logs
```bash
kubectl logs -n <kubedb-ns> <provisioner-pod-name>
kubectl logs -n <kubedb-ns> <opsmanager-pod-name>
```
Collect these when opening an issue for HanaDB alerts:

```bash
kubectl get hanadb -n <ns> <name> -oyaml
kubectl describe hanadb -n <ns> <name>
kubectl get pods -n <ns> -l app.kubernetes.io/instance=<name> -oyaml
kubectl describe pod -n <ns> <pod-name>
kubectl logs -n <ns> <pod-name> -c hanadb
kubectl logs -n <ns> <pod-name> -c hanadb-coordinator
kubectl logs -n <ns> <pod-name> -c exporter
kubectl logs -n <kubedb-ns> <operator-pod>
```

# Kubernetes notes
## Libraries
+ [Jaeger](https://github.com/jaegertracing/jaeger): tracing
+ [kops](https://github.com/kubernetes/kops): k8s install/upgrade/management
## [kubectl tips and tricks](https://discuss.kubernetes.io/t/kubectl-tips-and-tricks/192/7)
+ `kubectl alpha diff --help`: analyzes two kubernetes resources and prints the lines that are different.
+ `kubectl get pods --watch`: watch pod status
+ `kubectl top node`, `kubectl top pod` to see node/pod cpu/memory usage
+ Test connectivity from pod `jumpod` to service `thesvc`, service are accessible via FQDN 
  in the form `$SVC.$NAMESPACE.svc.cluster.local` [ref](http://kubernetesbyexample.com/sd/):
    - ping: `kubectl exec jumpod -c shell -i -t -- ping thesvc.default.svc.cluster.local` (non default namespace also works)
    - curl: `kubectl exec jumpod -c shell -i -t -- curl http://thesvc/info`
    - nslookup mysql service: `kubectl exec -ti $POD_NAME -- nslookup mysql`
+ Check environment variable: `kubectl exec <pod_name> -- printenv`
## Helm
+ `helm template`: locally render templates
+ `helm list | grep FAILED | awk '{print $1}' | xargs -L1 helm delete`: delete all failed releases

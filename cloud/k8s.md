# Kubernetes notes
## Libraries
+ [Jaeger](https://github.com/jaegertracing/jaeger): tracing
+ [kops](https://github.com/kubernetes/kops): k8s install/upgrade/management

## Tools
+ [Weave Scope](https://www.weave.works/docs/scope/latest/installing/#k8s), a tool to visualize and debug k8s
  - install `kubectl apply -f "https://cloud.weave.works/k8s/scope.yaml?k8s-version=$(kubectl version | base64 | tr -d '\n')"`
  - open scope in browser `kubectl port-forward -n weave "$(kubectl get -n weave pod --selector=weave-scope-component=app -o jsonpath='{.items..metadata.name}')" 4040`
+ [kail: k8s log viewer](https://github.com/boz/kail)
  - install `brew tap boz/repo;brew install boz/repo/kail`
  - [usage](https://github.com/boz/kail#usage): `kail --svc frontend --deploy webapp --log-level debug --ns monitoring`

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
### Port, Target Port and Nodeport [ref](https://vitalflux.com/kubernetes-port-targetport-and-nodeport/):
+ **Port**: Port is the port number which makes a service visible to other services running within the same K8s cluster.  In other words, in case a service wants to invoke another service running within the same Kubernetes cluster, it will be able to do so using port specified against “port” in the service spec file.
+ **Target Port**: Target port is the port on the POD where the service is running.
+ **Nodeport**: Node port is the port on which the service can be accessed from external users using Kube-Proxy
## Helm
+ `helm template`: locally render templates
+ `helm install --dry-run --debug <char_dir>`: check the generated manifests
+ `helm list | grep FAILED | awk '{print $1}' | xargs -L1 helm delete`: delete all failed releases

# Kubernetes notes
## Libraries and Tools
+ [Jaeger](https://github.com/jaegertracing/jaeger): tracing
+ [kops](https://github.com/kubernetes/kops): k8s install/upgrade/management
+ [Weave Scope](https://www.weave.works/docs/scope/latest/installing/#k8s), a tool to visualize and debug k8s
  - install `kubectl apply -f "https://cloud.weave.works/k8s/scope.yaml?k8s-version=$(kubectl version | base64 | tr -d '\n')"`
  - open scope in browser `kubectl port-forward -n weave "$(kubectl get -n weave pod --selector=weave-scope-component=app -o jsonpath='{.items..metadata.name}')" 4040`
+ [kail: k8s log viewer](https://github.com/boz/kail)
  - install `brew tap boz/repo;brew install boz/repo/kail`
  - [usage](https://github.com/boz/kail#usage): `kail --svc frontend --deploy webapp --log-level debug --ns monitoring`
+ [Telepresence](https://github.com/telepresenceio/telepresence): Local development against a remote Kubernetes
+ [squash](https://github.com/solo-io/squash): The debugger for microservices
+ [VSCode Kubernetes Tools](https://github.com/Azure/vscode-kubernetes-tools)
+ [draft](https://github.com/Azure/draft): steamlined k8s deployment, automate building, pushing image, applying yaml
## [kubectl tips and tricks](https://discuss.kubernetes.io/t/kubectl-tips-and-tricks/192/7)
+ `kubectl alpha diff --help`: analyzes two kubernetes resources and prints the lines that are different.
+ `kubectl get pods --watch`: watch pod status
+ `kubectl top node`, `kubectl top pod` to see node/pod cpu/memory usage
+ Test connectivity from pod `jumpod` to service `thesvc`, service are accessible via FQDN
  in the form `$SVC.$NAMESPACE.svc.cluster.local` [ref](http://kubernetesbyexample.com/sd/):
    - ping: `kubectl exec jumpod -c shell -i -t -- ping thesvc.default.svc.cluster.local` (non default namespace also works)
    - curl: `kubectl exec jumpod -c shell -i -t -- curl http://thesvc/info`
    - nslookup mysql service: `kubectl exec -ti $POD_NAME -- nslookup mysql`
+ `kubectl config get-contexts`: check k8s contexts
+ `kubectl config use-context docker-for-desktop`: choose k8s context
+ Check environment variable: `kubectl exec <pod_name> -- printenv`
+ `kubectl --dry-run blabla -o yaml`: generate manifests to stdout
+ `kubectl get deplyment -o yaml`
+ `kubectl port-forward svc/foo-svc 5000`: expose a in-cluster service
+ Zombie processes [issue](https://github.com/helm/charts/issues/2989#issuecomment-351053778): for deployment like rabbitmq, it's health check/liveness probe will create many zombie processes. This happens on older version k8s. To solve it, kubelet add `--pod-infra-container-image=gcr.io/google_containers/pause-amd64:3.1` to use newer pause container to reap zombies. Thank you [almighty pause container](https://www.ianlewis.org/en/almighty-pause-container)
+ `kubectl cluster-info dump`: to debug and diagnose cluster problems
+ `kubectl get pods | grep Evicted | awk '{print $1}' | xargs kubectl delete pod`: Delete evicted pods
+ Accessing pod metadata [ref](https://github.com/luksa/kubernetes-in-action/blob/master/Chapter08/downward-api-env.yaml)
+ `kubectl api-versions`: Find api versions cluster support.
+ Debugging pod: `kubectl describe pod <pod-id>`, `kubectl get pods -o wide`
+ `kubectl exec POD_NAME -c CONTAINER_NAME reboot`: rebbot specific container of pod
### Port, Target Port and Nodeport [ref](https://vitalflux.com/kubernetes-port-targetport-and-nodeport/):
+ **Port**: Port is the port number which makes a service visible to other services running within the same K8s cluster.  In other words, in case a service wants to invoke another service running within the same Kubernetes cluster, it will be able to do so using port specified against “port” in the service spec file.
+ **Target Port**: Target port is the port on the POD where the service is running.
+ **Nodeport**: Node port is the port on which the service can be accessed from external users using kube-proxy
+
## Helm
+ `helm template`: locally render templates
+ `helm install --dry-run --debug <char_dir>`: check the generated manifests
+ `helm list | grep FAILED | awk '{print $1}' | xargs -L1 helm delete`: delete all failed releases
+ [stable/rabbimq/templates/svc.yaml](https://github.com/helm/charts/blob/master/stable/rabbitmq/templates/svc.yaml)
+ [test-vault-status.yaml](https://github.com/banzaicloud/banzai-charts/blob/master/vault/templates/tests/test-vault-status.yaml)

## k8s recipes
+ init container (TODO: local repo)

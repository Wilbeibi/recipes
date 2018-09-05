# Skaffold example
## Steps to run example
1. make and docker build once to let skaffold know how to build
2. Given a k8s running, `skaffold dev` to start listening
4. `make` to generate new executable
4. skaffold will notice it and run docker build and push to public registry,
apply k8s manifests changes
5. Watch the logs from `skaffold dev` output
## References
+ [Skaffold: happy Kubernetes workflows](https://ahmet.im/blog/skaffold/)
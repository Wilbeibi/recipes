# Golang Notes
## vgo [ref](https://zhuanlan.zhihu.com/p/33926171) (TODO: deprecated, called go module now)
+ touch go.mod; vgo build
+ `vgo list -m`: check all dependencies
+ `vgo list -m -u`: check all dependencies and update
+ `vgo test all`: execute all tests
+ `vgo get -u`: upgrade all dependencies
+ `vgo list -u rsc.io/sample`: check all available version(tag) of specific package
+ `vgo get rsc.io/sample@1.3.1`: get specific version and modify go.mod
+ `exclude "rsc.io/sampler" v1.99.99`: exclude version v1.99.99 via add this line to go.mod
+ `replace "rsc.io/quote" v1.5.2 => "github.com/you/quote" v0.0.0-myfork`: replace package
+ `replace "rsc.io/quote" v1.5.2 => "../quote"`: replace with local directory
+ `vgo vendor`: put dependencies in vendor to comfort traditional `go build` person


## READ
+ [tar](https://github.com/GoogleContainerTools/skaffold/blob/master/pkg/skaffold/util/tar.go)
+ [update-fmt.sh](https://github.com/heptio/ark/blob/master/hack/update-fmt.sh)
    - `gofmt -w -s $(find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./pkg/generated/*" -not -name "zz_generated*")`
    - `goimports > /dev/null || go get golang.org/x/tools/cmd/goimports` [about ||](https://unix.stackexchange.com/a/24685/36211)
    - `goimports -w -d $(find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./pkg/generated/*" -not -name "zz_generated*")`
+ [notify](https://github.com/oxequa/realize/blob/master/realize/notify.go)
+ [vault/helper](https://github.com/hashicorp/vault/tree/master/helper)
    + [read password from stdin/file](https://github.com/hashicorp/vault/blob/master/helper/password/password.go)
## Mise
+ [variadic parameter in interface](https://github.com/go-kit/kit/blob/master/metrics/metrics.go)

## My awesome go list
+ [viper](https://github.com/spf13/viper): Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats
+ [cobra](https://github.com/spf13/cobra): My cobra demo app at https://github.com/wilbeibi/cobra-demo
+ [skaffold pkg util](https://github.com/GoogleContainerTools/skaffold/blob/master/pkg/skaffold/util/util.go)
+ [kubewatch k8sutil](https://github.com/bitnami-labs/kubewatch/blob/master/pkg/utils/k8sutil.go)

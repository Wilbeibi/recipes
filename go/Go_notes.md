# Golang Notes

## Go module [ref](https://tonybai.com/2018/11/19/some-changes-in-go-1-11/)
+ `export GO111MODULE=on`: turn on go mod
+ `go mod init`: create go.mod file, if found existing version control tools, it will generate deps from them
+ `go build`: with go.mod exists, go build will fill the content of it and generate go.sum records dep version and hash
+ `go mod verify` to verify dependencies' versions are intact
+ `go mod why golang.org/x/oauth2` check whether oauth2 is a valid dependency
+ `go mod tidy` to clean up go.mod and go.sum
+ `go list -m -versions golang.org/x/text` to check all golang.org/x/text versions
+ `go get golang.org/x/text@v0.1.0` base on previous versions, downgrade it from v0.3.0 to v0.1.0
+ `go build -mod=vendor` to build with vendor for reproduceable build
+ cannot find module providing package issue: `go clean -modcache` and `go build` again


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

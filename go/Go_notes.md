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
+ `go build -mod=vendor` to build with vendor for reproducible build
+ cannot find module providing package issue: `go clean -modcache` and `go build` again

## FAQ
+ [What is x.(T)?](https://golang.org/ref/spec#Type_assertions)
+ [What is omitempty in json/xml parsing](https://www.socketloop.com/tutorials/golang-meaning-of-omitempty-in-struct-s-field-tag)

## READ
+ [tar](https://github.com/GoogleContainerTools/skaffold/blob/master/pkg/skaffold/util/tar.go)
+ [update-fmt.sh](https://github.com/heptio/ark/blob/master/hack/update-fmt.sh)
    - `gofmt -w -s $(find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./pkg/generated/*" -not -name "zz_generated*")`
    - `goimports > /dev/null || go get golang.org/x/tools/cmd/goimports` [about ||](https://unix.stackexchange.com/a/24685/36211)
    - `goimports -w -d $(find . -type f -name "*.go" -not -path "./vendor/*" -not -path "./pkg/generated/*" -not -name "zz_generated*")`
+ [notify](https://github.com/oxequa/realize/blob/master/realize/notify.go)
+ [vault/helper](https://github.com/hashicorp/vault/tree/master/helper)
    + [read password from stdin/file](https://github.com/hashicorp/vault/blob/master/helper/password/password.go)
+ [k8s.io/client-go/util/workqueue](https://github.com/kubernetes/client-go/tree/master/util/workqueue)
+ [skaffold pkg util](https://github.com/GoogleContainerTools/skaffold/blob/master/pkg/skaffold/util/util.go)
+ [kubewatch k8sutil: in cluster client/out of cluster client](https://github.com/bitnami-labs/kubewatch/blob/master/pkg/utils/k8sutil.go#L18-L50)
+ [Pluck](https://github.com/pulumi/pulumi-kubernetes/blob/master/pkg/openapi/openapi.go#L121) like `json.Unmarshal` for unstructured
+ [go tour webcrawler](https://gist.github.com/wilbeibi/7f0c5d8cf60266bbb8f921029c4d7edf)

## Mise
+ [variadic parameter in interface](https://github.com/go-kit/kit/blob/master/metrics/metrics.go)
+ [returns the first IPv4 address of the supplied interface name](https://github.com/weaveworks/common/blob/master/network/interface.go#L9)

## Awesome Go libraries
+ [viper](https://github.com/spf13/viper): Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats
+ [cobra](https://github.com/spf13/cobra): My cobra demo app at https://github.com/wilbeibi/cobra-demo
+ [bleve: lightweight index/search engine](https://github.com/blevesearch/bleve)
    + [demo: beer-search](https://github.com/blevesearch/beer-search)
+ [gocui: Console User Interfaces](https://github.com/jroimartin/gocui)
    + [gitbatch usage](https://github.com/isacikgoz/gitbatch)
+ [go-prompt: interactive prompts in Go](https://github.com/c-bata/go-prompt)
    + [tldr](https://github.com/isacikgoz/tldr)
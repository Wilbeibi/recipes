# Golang Notes
## vgo [ref](https://zhuanlan.zhihu.com/p/33926171)
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

## My awesome go list
+ [viper](https://github.com/spf13/viper): Viper is a complete configuration solution for Go applications including 12-Factor apps. It is designed to work within an application, and can handle all types of configuration needs and formats
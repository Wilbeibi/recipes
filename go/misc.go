// A little copying is better than a little dependency
// From: https://groups.google.com/d/msg/golang-nuts/d0nF_k4dSx4/rPGgfXv6QCoJ
package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func uuid() string {
	f, err := os.Open("/dev/urandom") // ignore errcheck
	if err != nil {
		panic(err)
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// RandomID, from https://github.com/GoogleContainerTools/skaffold/blob/master/pkg/skaffold/util/util.go#L34
func RandomID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", b)
}

func splitAndFields() {
	fmt.Printf("%q\n", strings.Split(" hello  world", " ")) // ["" "hello" "" "world"]
	fmt.Printf("%q\n", strings.Fields(" hello  world"))     // ["hello" "world"]
	// Notice, it's ' ', because c is rune, not string, line 24 use " ", has to be string
	f := func(c rune) bool { return c == ',' || c == ' ' }
	fmt.Printf("%q\n", strings.FieldsFunc(" Hello, hello, world", f))
}

func printVersion() {
	fmt.Printf("Go Version: %s", runtime.Version())
	fmt.Printf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
}

// From github.com/go101/go101/go101.go
func runShellCommand(timeout time.Duration, cmd string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	command := exec.CommandContext(ctx, cmd, args...)
	command.Dir = rootPath
	return command.Output()
}

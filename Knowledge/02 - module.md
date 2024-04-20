# 3. Go Module

A collection of GO packages (dependencies) that are related and dependent in the project.

> Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.

Module management files in Golang consist of 2 files; 

1. `go.mod`: List all dependencies in the project.
2. `go.sum`: Keep the version of dependencies in the project.

## Initialize Go.mod (Go Modules)

Init the Go Module

```sh
$ mkdir go-module
$ cd go-module
$ go mod init example.com/yourusername/go-example
```

From the above snippet. We created a `go.mod` file for managing the dependencies in Go project. `go mod init` told Go to initialize the `go module` management file. 

The `example.com/yourusername/go-example` indicates the location where the go tools can find the source code. Usually it will be Github repo.

Inside `go.mod` file, the `module` indicates the src code location. And below is the version of dependency, here `go` is the only dependency with associated version.

```txt
module example.com/yourusername/go-example

go 1.22.2
```


# Call code from another module

Reference with directory `00 - basic, syntax, modules, packaging/02 - call from another module`

Suppose there are 2 Go modules. Namely; `greetings` and `sayhello`. 

`sayhello` needs to call `Hello()` function from `greetings` module. 

We can add local edit the `go.mod` in `sayhello` module to instruct that source code of `greetings` module is not in the module repository but in local directory.

```sh
$ cd '00 - basic, syntax, modules, packaging/02 - call from another module/sayhello'
$ go mod edit -replace example.com/greeting=./greeting
$ go mod tidy
> go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

**Reference**
- https://go.dev/doc/tutorial/call-module-code
# What?

This package exposes some internal packages from
[`golang.org/x/tools/internal`](https://github.com/golang/tools/tree/master/internal), by simply copying them.

# How?

> **Note:** This section explains how this package is created, **not how it should be used** by the calling code.

Here we've exposed the following packages from `golang.org/x/tools/internal`: `bug`, `diff`, `event`, `fakenet`,
`fuzzy`, `gocommand`, `goroot`, `jsonrpc2`, `jsonrpc2_v2`, `memoize`, `persistent`, `stack`, `testenv` and `xcontext`.

Here's how we've done this:

- First we've copied all these packages from the original source to this repo.
- After that we've adjusted `import` statements by replacing `golang.org/x/tools/internal/<package_name>` with
  `github.com/peske/x-tools-internal/<package_name>` for all the imported packages.

This is done automatically, by using _copy tool_. To use the tool you first need to build it:

```bash
go build .
go generate
```

This will create `x-tools-internal` executable file in the working directory. After that you can use the copy tool in
the following way:

```bash
./x-tools-internal /path/to/golang.org/x/tools
```

The CLI argument provided represents the local path of `golang.org/x/tools` module. The argument is optional, and if not
provided it will default to `$GOHOME/src/golang.org/x/tools`.

Assuming that the source module is located at its default location (`$GOHOME/src/golang.org/x/tools`), previous two
steps can be replaced by simply executing `go generate`, since the [`main.go`](./main.go) file contains the following
lines:

```go
//go:generate go build .
//go:generate ./x-tools-internal
```

# Why?

The original packages are used in some other `golang.org/x/tools` packages that we want to rewrite, but we cannot use
them directly because they are `internal`. For example, these packages are used in
[peske/lsp-srv package](https://github.com/peske/lsp-srv).

# License?

The same license as the original one - [BSD-3-Clause license](./LICENSE). Although almost all the code is created by the
authors of the original module (`The Go Authors`), we've changed copyright here to `Fat Dragon and authors` not to get
the credits, but to protect the original authors of any responsibility if there are any problems in the code that we've
changed. All credits should go to the authors of the original module.

# Version?

Current `main` branch is based on the original repository commit
[3e6f71b](https://github.com/golang/tools/commit/3e6f71bba4359aeb7a301d361ee3cf95e8799599) from January 17, 2023.

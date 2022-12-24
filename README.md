# What?

This package exposes some internal packages from
[`golang.org/x/tools/internal`](https://github.com/golang/tools/tree/master/internal), by simply copying them.

# How?

Here we've exposed the following packages from `golang.org/x/tools/internal`: `bug`, `diff`, `event`, `fakenet`,
`fuzzy`, `gocommand`, `goroot`, `jsonrpc2`, `jsonrpc2_v2`, `memoize`, `persistent`, `stack`, `testenv` and `xcontext`.

Here's how we've done this:

- First we've copied all these packages from the original source to this repo.
- After that we've adjusted `import` statements by replacing `golang.org/x/tools/internal/<package_name>` with
  `github.com/peske/x-tools-internal/<package_name>` for all the imported packages.

All this can be done by using a [copy tool](./_copy_tool). The tool automatically copies packages from
`golang.org/x/tools/internal` into this repository. Usage:

- `cd` into `./_copy_tool` directory, and build the tool by executing `go build -o ../cptool`. Note that the executable
  is stored in the root directory of this repository.
- `cd` into the root directory of this repository and execute: `./cptool /path/to/golang.org/x/tools/internal` (change
  the source path appropriately).

# Why?

The original packages are used in some other `golang.org/x/tools` packages that we want to rewrite, but we cannot use
them directly because they are `internal`. For example, these packages are used in
[peske/lsp-srv package](https://github.com/peske/lsp-srv).

# License?

The same license as the original one - [BSD-3-Clause license](./LICENSE).

# Version?

Current `main` branch is based on the original repository commit
[eb70795](https://github.com/golang/tools/commit/eb70795aaccb8e6c9615c88085ef3414ba04b8c9) from December 17, 2022.

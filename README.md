# What?

This package exposes some internal packages from
[`golang.org/x/tools/internal`](https://github.com/golang/tools/tree/master/internal), by simply copying them.

# How?

We first copy desired packages (currently these are `bug`, `event`, `fakenet`, `fuzzy`, `gocommand`, `jsonrpc2`,
`jsonrpc2_v2`, `memoize`, `persistent` and `xcontext`) into this repository, and perform the following replacements
across all the files:

- `golang.org/x/tools/internal/bug` with `github.com/peske/x-tools-internal/bug`
- `golang.org/x/tools/internal/event` with `github.com/peske/x-tools-internal/event`
- `golang.org/x/tools/internal/fakenet` with `github.com/peske/x-tools-internal/fakenet`
- `golang.org/x/tools/internal/fuzzy` with `github.com/peske/x-tools-internal/fuzzy`
- `golang.org/x/tools/internal/gocommand` with `github.com/peske/x-tools-internal/gocommand`
- `golang.org/x/tools/internal/jsonrpc2_v2` with `github.com/peske/x-tools-internal/jsonrpc2_v2`
- `golang.org/x/tools/internal/jsonrpc2` with `github.com/peske/x-tools-internal/jsonrpc2`
- `golang.org/x/tools/internal/memoize` with `github.com/peske/x-tools-internal/memoize`
- `golang.org/x/tools/internal/persistent` with `github.com/peske/x-tools-internal/persistent`
- `golang.org/x/tools/internal/xcontext` with `github.com/peske/x-tools-internal/xcontext`

Finally, we also remove `golang.org/x/tools/internal/stack/stacktest` dependency. Luckily, this dependency is only used
in a few test files, and it was easy to remove: simply remove it from `import`, and delete (comment out) the lines of
code where it is used (only very few of them). For unit testing we rely on the original package anyway.

# Why?

The original packages are used in some other `golang.org/x/tools` packages that we want to rewrite, but we cannot use
them directly because they are `internal`. For example, these packages are used in
[peske/lsp package](https://github.com/peske/lsp).

# License?

The same license as the original one - [BSD-3-Clause license](./LICENSE).

# Version?

Current `main` branch is based on the original repository commit
[eb70795](https://github.com/golang/tools/commit/eb70795aaccb8e6c9615c88085ef3414ba04b8c9) from December 17, 2022.

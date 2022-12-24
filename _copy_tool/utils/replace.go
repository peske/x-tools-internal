package utils

import "strings"

var Packages = []string{
	"bug",
	"diff",
	"event",
	"fakenet",
	"fuzzy",
	"gocommand",
	"goroot",
	"jsonrpc2",
	"jsonrpc2_v2",
	"memoize",
	"persistent",
	"stack",
	"testenv",
	"xcontext",
}

func Replace(content string) string {
	for _, p := range Packages {
		content = strings.Replace(content,
			"\"golang.org/x/tools/internal/"+p,
			"\"github.com/peske/x-tools-internal/"+p, -1)
	}
	return content
}

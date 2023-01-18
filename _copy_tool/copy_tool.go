package _copy_tool

import (
	"fmt"
	"os"
	"path"
	"strings"
)

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

func CheckDirectoryExists(dir string) error {
	if s, err := os.Stat(dir); err != nil {
		return err
	} else if !s.IsDir() {
		return fmt.Errorf("'%s' is not a directory", dir)
	}
	return nil
}

func replace(fp string, content []byte) []byte {
	if strings.HasSuffix(strings.ToLower(fp), ".md") {
		// We don't want to replace in Markdown files
		return content
	}
	str := string(content)
	for _, p := range Packages {
		str = strings.Replace(str,
			"\"golang.org/x/tools/internal/"+p,
			"\"github.com/peske/x-tools-internal/"+p, -1)
	}
	return []byte(str)
}

func copyFile(src, dst string, replaceFn func(string, []byte) []byte) error {
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	var content []byte
	if content, err = os.ReadFile(src); err != nil {
		return err
	}

	content = replace(dst, content)
	if replaceFn != nil {
		content = replaceFn(dst, content)
	}

	return os.WriteFile(dst, content, srcinfo.Mode())
}

func CopyDir(src string, dst string, replaceFn func(string, []byte) []byte) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = CopyDir(srcfp, dstfp, replaceFn); err != nil {
				return err
			}
		} else {
			if err = copyFile(srcfp, dstfp, replaceFn); err != nil {
				return err
			}
		}
	}
	return nil
}

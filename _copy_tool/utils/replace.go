package utils

import (
	"fmt"
	"log"
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

func Replace(content string) string {
	for _, p := range Packages {
		content = strings.Replace(content,
			"\"golang.org/x/tools/internal/"+p,
			"\"github.com/peske/x-tools-internal/"+p, -1)
	}
	return content
}

func EnsureDir(dir string) {
	if s, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("'%s' does not exist", dir)
		} else {
			log.Fatal(err)
		}
	} else if !s.IsDir() {
		log.Fatalf("'%s' is not a directory", dir)
	}
}

func CopyFile(src, dst string) error {
	var err error
	var srcinfo os.FileInfo

	c, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	str := Replace(string(c))

	err = os.WriteFile(dst, []byte(str), 0700)
	if err != nil {
		return err
	}

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func CopyDir(src string, dst string) error {
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
			if err = CopyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = CopyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

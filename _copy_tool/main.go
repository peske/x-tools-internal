package main

import (
	"fmt"
	"github.com/peske/x-tools-internal/_copy_tool/utils"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Source directory not specified.")
	}
	src := os.Args[1]
	ensureDir(src)
	for _, p := range utils.Packages {
		ensureDir(p)
		ensureDir(filepath.Join(src, p))
	}
	for _, p := range utils.Packages {
		if err := os.RemoveAll(p); err != nil && !os.IsNotExist(err) {
			log.Fatalln(err)
		}
		if err := copyDir(filepath.Join(src, p), p); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Finished")
}

func ensureDir(dir string) {
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

func copyFile(src, dst string) error {
	var err error
	var srcinfo os.FileInfo

	c, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	str := utils.Replace(string(c))

	err = os.WriteFile(dst, []byte(str), 0700)
	if err != nil {
		return err
	}

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

func copyDir(src string, dst string) error {
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
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

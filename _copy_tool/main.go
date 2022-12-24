package main

import (
	"github.com/peske/x-tools-internal/_copy_tool/utils"
	"log"
	"os"
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
		if err := utils.CopyDir(filepath.Join(src, p), p); err != nil {
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

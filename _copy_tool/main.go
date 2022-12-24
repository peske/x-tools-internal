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
	utils.EnsureDir(src)
	for _, p := range utils.Packages {
		utils.EnsureDir(p)
		utils.EnsureDir(filepath.Join(src, p))
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

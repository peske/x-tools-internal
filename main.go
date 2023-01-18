package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/peske/x-tools-internal/_copy_tool"
)

func main() {
	// Get source path
	var src string
	var err error
	switch len(os.Args) {
	case 1:
		h := os.Getenv("GOHOME")
		if h == "" {
			log.Fatalln("Source directory not specified.")
		}
		src = filepath.Join(h, "src", "golang.org", "x", "tools")
	case 2:
		if src, err = filepath.Abs(os.Args[1]); err != nil {
			log.Fatalf("Invalid source path '%s'", os.Args[1])
		}
	default:
		log.Fatalln("Invalid arguments.")
	}

	// Check source validity
	if err = checkSourceValidity(src); err != nil {
		log.Fatalln(err)
	}

	// Delete destination
	dst, _ := filepath.Abs(".") // should not error ever
	for _, p := range _copy_tool.Packages {
		if err = os.RemoveAll(filepath.Join(dst, p)); err != nil && !os.IsNotExist(err) {
			log.Fatalln(err)
		}
	}

	// Copy from source
	for _, p := range _copy_tool.Packages {
		if err = _copy_tool.CopyDir(filepath.Join(src, "internal", p), filepath.Join(dst, p), nil); err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Finished.")
}

func checkSourceValidity(src string) error {
	if err := _copy_tool.CheckDirectoryExists(src); err != nil {
		return err
	}
	src = filepath.Join(src, "internal")
	if err := _copy_tool.CheckDirectoryExists(src); err != nil {
		return err
	}
	for _, p := range _copy_tool.Packages {
		if err := _copy_tool.CheckDirectoryExists(filepath.Join(src, p)); err != nil {
			return err
		}
	}
	return nil
}

//go:generate ./x-tools-internal

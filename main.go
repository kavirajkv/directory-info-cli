package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
)

// file and size to store
type FileInfo struct {
	Path string
	Size int64
	lastusage time.Time
}

// scanDirectory scans a directory and returns a list of files
func scanDirectory(dir string) ([]FileInfo, error) {
	var files []FileInfo


	// ref: https://pkg.go.dev/path/filepath#Walk (p.s with reference to the given )
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, FileInfo{
				Path: path,
				Size: info.Size(),
				lastusage: info.ModTime(),
			})
		}
		return nil
	})

	return files, err
}

func main() {

	//by default here pwd is used to scan
	dir:=flag.String("dir", ".", "Directory path to analyze")
	help:=flag.Bool("help",false,"by default PWD will be used as dir path")
	flag.Parse()

	if *help{
		flag.Usage()
		os.Exit(0)
	}


	usage, err := disk.Usage(*dir)
	if err != nil {
		log.Fatalf("Error getting disk usage: %v", err)
	}
	//note: usage value is in bytes
	fmt.Printf("Total disk usage: %.2f GB\n", float64(usage.Used)/1024/1024/1024)


	fmt.Printf("Scanning directory: %s\n", *dir)
	files, err := scanDirectory(*dir)
	if err != nil {
		log.Fatalf("Error scanning directory: %v", err)
	}


	// List large files
	fmt.Println("\nLarge files in the given directory :")
	for _, file := range files {
		//greater then 100MB can be shown as large files
		if file.Size > 100*1024*1024 { 
			fmt.Printf("%s - %.2f MB - last usage at: %v", file.Path, float64(file.Size)/1024/1024,file.lastusage)
		}
	}	
}


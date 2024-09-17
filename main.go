package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var fileCategories = map[string]string{
	".jpg": "Images", ".jpeg": "Images", ".png": "Images",
	".txt": "Documents", ".docx": "Documents", ".pdf": "Documents",
	".mp3": "Music", ".wav": "Music",
	".mp4": "Videos", ".mkv": "Videos",
}

func main() {
	dirPtr := flag.String("dir", ".", "Directory to organize")
	flag.Parse()

	handleInterrupt()

	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		fmt.Println("Error: Directory does not exist")
		os.Exit(1)
	}

	fmt.Println("Organizing directory:", *dirPtr)
	organizeFiles(*dirPtr)
}

func organizeFiles(dir string) {
	files := listFiles(dir)

	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			category, exists := fileCategories[ext]

			if exists {
				categoryDir := filepath.Join(dir, category)
				if _, err := os.Stat(categoryDir); os.IsNotExist(err) {
					os.Mkdir(categoryDir, 0755)
				}

				srcPath := filepath.Join(dir, file.Name())
				destPath := filepath.Join(categoryDir, file.Name())
				err := os.Rename(srcPath, destPath)
				if err != nil {
					fmt.Println("Error moving file:", err)
				} else {
					fmt.Printf("Moved %s to %s\n", file.Name(), category)
				}
			}
		}
	}
}

func listFiles(dir string) []os.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return nil
	}
	return files
}

func handleInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("\nOperation interrupted. Exiting gracefully...")
		os.Exit(0)
	}()
}

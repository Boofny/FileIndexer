// Package indexer will be the starting point from where this file indexer will start and grow
package indexer 

// TODO: For now the first thing is to build the file tree using recursion and then go routines
import (
	"fmt"
	"os"
	"path/filepath"
)

const tempDir = "./"

// NOTE: for now what i am thinking is doing go rourotines searching rather than searching the whole tree at once then serch all kids at once so on

// Start just the entry point for the code
func Start(){
	dirfs := os.DirFS(tempDir) // may want to have this for the future 
	file, err := dirfs.Open("cmd")
	if err != nil {
		return 
	}
	defer file.Close()

	path, err := filepath.Abs(tempDir)
	if err != nil {
		panic(err)
	}

	tree := NewFileTree(path)
	tree.BuildTree(path)

	for _, v := range tree.Files{
		fmt.Println("file in root file: ", v.Name)
	}
	tree.PrintDirTree("-")
}





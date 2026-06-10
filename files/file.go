// Package files will be the starting point from where this file indexer will start and grow
package files

// TODO: For now the first thing is to build the file tree using recursion and then go routines 
import (
	"fmt"
	"os"
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

	tree := NewFileTree(tempDir)
	tree.BuildTree(tempDir)
	for _, v := range tree.Children {
		fmt.Println("Directory children: ", v.Name)
	}

	for _, i := range tree.Children[2].Files{
		fmt.Println(tree.Children[0].Name)
		fmt.Print("Name: ")
		fmt.Println(i.Name)
	}

	fmt.Print("Root: ", tree.Files[0].Name)
}





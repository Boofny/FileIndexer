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

	// dir, err := os.ReadDir(tempDir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(dir[len(dir)-1].Name())

	// for _, v := range dir{
	// 	fmt.Println(v.Name(), v.IsDir())
	// 	if v.IsDir() {
	// 		fmt.Println(v.Name())
	// 		dirs, err := os.ReadDir(v.Name())
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 		if len(dirs) == 0 {
	// 			continue
	// 		}
	// 		fmt.Println(dirs[0].Name())
	// 	}
	// }
	tree := NewFileTree(tempDir)
	tree.BuildTree(tempDir)
	fmt.Print("Root: ", tree.Name)
}





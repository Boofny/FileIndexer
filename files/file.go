// Package files will be the starting point from where this file indexer will start and grow
package files

import (
	"fmt"
	"log"
	"os"
)

type File struct{
	Name 		string // name of the file
	Type 		string // .txt .go etc...
	Size 		int32
	Content []byte // used for later searching
	IsBin 	bool // checking for binaray files for not allowing them to be open
}

type Directory struct{
	Files []File // could have many files inside of a dir
	Name string // name of the directory
	Childrean []*Directory
}

const tempDir = "./"

// NOTE: for now what i am thinking is doing go rourotines searching rather than searching the whole tree at once then serch all kids at once so on

// Start just the entry point for the code
func Start(){
	dirfs := os.DirFS(tempDir) // may want to have this for the future 
	file, err := dirfs.Open("go.mod")
	if err != nil {
		return 
	}
	defer file.Close()

	// buf := make([]byte, 128)

	// fmt.Println(file.Read(buf))
	// fmt.Println(string(buf))

	err = os.Chdir(tempDir)
	if err != nil {
		log.Fatal(err)
	}
	dir, err := os.ReadDir(tempDir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir[len(dir)-1].Name())
	for _, v := range dir{
		fmt.Println(v.Name(), v.IsDir())
		if v.Name()[0] == '.'{ // using this bear bones way to check dot dirs
			break
		}
	}
}

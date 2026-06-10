package files

import (
	"log"
	"os"
)

const banned = "node_modules"

type File struct{
	Name 		string // name of the file
	Type 		string // .txt .go etc...
	Path 		string // the full path of a file ~/codeBlocks/Thing.java
	Size 		int32  // size of a file in bytes
	IsBin 	bool 	 // checking for binaray files for not allowing them to be open
	Content []byte // used for later searching
}

type Directory struct{
	Files []File 	// could have many files inside of a dir
	Name string 	// name of the directory
	// Type string 	// type of dir like . or not
	Childrean []*Directory
}

func NewFileTree(rootName string)Directory{
	return Directory{
		Name: rootName,
	} // returns nothing for now
}

func (d *Directory)BuildTree(dirName string)Directory{
	dir, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
		// return Directory{} // for error purposes
	}

	for _, v := range dir{
		if v.IsDir() && !IsDotDir(v.Name()) && v.Name() != banned{
			// fmt.Print(v.Name())
			d.Childrean = append(d.Childrean, &Directory{Name: v.Name()})
			d.BuildTree(v.Name())
		}else{ // then the v is a file will add the 
			d.Files = append(d.Files, addChildren(v)...)
		}
	}
	return *d // idk if this will work
}

func (d *Directory)IndexChildren(){

}

// addChildren thinking or returning the file Childrean non dirs
func addChildren(dir os.DirEntry)[]File{
	return nil 
}

// IsDotDir used to check that we dont go into .git or somthing else
func IsDotDir(dirName string)bool{
	return dirName[0] == '.'
}

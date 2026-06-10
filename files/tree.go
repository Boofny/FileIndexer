package files

import (
	"log"
	"os"
	"path/filepath"
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
	Children  []*Directory
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

	for _, v := range dir{ // for now this is JUST getting the other dirs
		if !v.IsDir(){ // then the v is a file will add the 
			filePath := GetFilePath(v.Name())
			size, err := os.Stat(v.Name())
			if err != nil {
				return Directory{}
			}
			d.Files = append(d.Files, File{
				Name: v.Name(), // hope this works
				Type: filepath.Ext(v.Name()),
				Path: filePath,
				Size: int32(size.Size()),
				IsBin: false, // just for now
				Content: nil,	 // also just for now
			})

		}else if v.IsDir() && !IsDotDir(v.Name()) && v.Name() != banned{

			d.Children = append(d.Children, &Directory{Name: v.Name(), Files: d.Files}) // is just sending the same children from the main dir ./
			d.BuildTree(v.Name())
		}
	}
	return *d // idk if this will work
}

func GetFilePath(fileName string)string{
	filePath, err := filepath.Abs(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return filePath
}

func GetFileSize(fileName string)int32{
	size, err := os.Stat(fileName)
	if err != nil {
		panic(err)
	}
	return int32(size.Size())
}

// addChildren thinking or returning the file children  non dirs
func addChildren(dir os.DirEntry)[]File{
	// var file []File
	// for _, v := range dir{
	// 	file = append(file, File{
	// 		Name: v,
	// 	})
	// }
	return nil 
}

// IsDotDir used to check that we dont go into .git or somthing else
func IsDotDir(dirName string)bool{
	return dirName[0] == '.'
}

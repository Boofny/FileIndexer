package indexer

import "fmt"

// print.go is just used for debug prints of the tree

func (d *Directory)PrintDirTree(indent string){
	fmt.Println(indent + d.Name)
	for _, f := range d.Children{
		f.PrintDirTree(indent + indent)
		PrintFilesInDir(*f, indent)
	}
}

func PrintFilesInDir(dir Directory, indent string){
	for _, v := range dir.Files{
		fmt.Println(" " + indent + v.Name)
	}
}

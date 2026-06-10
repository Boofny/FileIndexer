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
		fmt.Printf(" %s %s %s %s %d\n", indent, v.Name, v.Type, v.Path, v.Size)
		// fmt.Print(" " + indent + v.Name + string(v.Size))
	}
}

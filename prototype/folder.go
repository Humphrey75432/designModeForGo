package prototype

import "fmt"

type Folder struct {
	Children []Inode
	Name     string
}

func (f Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, v := range f.Children {
		v.Print(indentation + indentation)
	}
}

func (f Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []Inode
	for _, v := range f.Children {
		clone := v.Clone()
		tempChildren = append(tempChildren, clone)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}

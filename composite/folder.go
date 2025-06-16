package composite

import "fmt"

type Folder struct {
	Components []Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in folder %s\n", keyword, f.Name)
	for _, compose := range f.Components {
		compose.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}

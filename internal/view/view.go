package view

import "fmt"

func Show(p Page) string {
	var output = ""
	p.PrintAllItems()
	return output
}

func (hp *ManagerHomePage) PrintAllItems() {
	for idx, item := range hp.Items {
		fmt.Printf("%d.%s\n", idx+1, item)
	}
}

package view

type Page interface {
	// SetItems()
	PrintAllItems()
}

type OrderPickerHomePage struct {
	Title string
	Items []string
}

type ManagerHomePage struct {
	Title string
	Items []string
}

type DispatcherHomePage struct {
	Title string
	Items []string
}

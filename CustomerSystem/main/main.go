package main

import (
	"go_code/CustomerSystem/view"
)

func main() {
	customer := view.NewCustomerView()
	customer.Menu()
}

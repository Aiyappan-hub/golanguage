package main

import "fmt"

func main() {

	type company struct {
		Name    string
		field   string
		city    string
		Pincode int
	}

	var a company
	var b company
	var c company

	fmt.Println("a, b, c")

	a = company{Name: "kloudone", field: "IT", city: "California", Pincode: 95762}

	fmt.Println("Company1:", a)

	b = company{Name: "wipro", field: "IT", city: "Banglore", Pincode: 277001}

	fmt.Println("Company2:", b)

	c = company{Name: "standard chattered", field: "Non-IT", city: "chennai", Pincode: 600003}

	fmt.Println("Company3:", c)
}

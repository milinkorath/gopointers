package main

import "fmt"

func main() {

	var p *int
	var mystring *string
	var mybool *bool
	p1 := 20
	p = &p1
	mystring1 := "string"
	mystring = &mystring1
	mybool1 := true
	mybool = &mybool1

	fmt.Println(*p)
	fmt.Println(*mystring)
	fmt.Println(*mybool)

	fmt.Println(p)
	fmt.Println(mystring)
	fmt.Println(mybool)

}

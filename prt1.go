package main

import "fmt"

func main() {

	p := 20
	mystring := "hello"
	mybool := true

	p1 := &p
	s1 := &mystring
	b1 := &mybool
	fmt.Println(p)
	fmt.Println(mystring)
	fmt.Println(mybool)

	fmt.Println(&p)
	fmt.Println(&mystring)
	fmt.Println(&mybool)

	fmt.Println(p1)
	fmt.Println(s1)
	fmt.Println(b1)

	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", b1)

	pp := &p1
	fmt.Printf("%T\n", pp)

	fmt.Println(**pp)
}

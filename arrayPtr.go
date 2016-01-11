package main

import "fmt"

func main() {

	a := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	arrayModify(&a)
	fmt.Println(a)
	d := new(Struct)
	i := 40
	myArray(d, &i)
	fmt.Println(*d)
}

func arrayModify(a *[4]int) {
	fmt.Println(a)
	// y := *a
	a[0] = 1000
	// y[0] = 1000
	// fmt.Println(y[0])
}

type Struct struct {
	x int
	y int
}

func myArray(a *Struct, ptrint *int) {
	(*a).x++

	(*a).y++
	fmt.Println(*ptrint)
}

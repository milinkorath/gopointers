package main

import "fmt"

func main() {

	u := User{"Milin"}
	// Modify(u)   //Not possible
	fmt.Println(u.name)
	Modify(&u)
	fmt.Println(u.name)
	fmt.Println("---------------------")
	Modify1(&u)
	fmt.Println(u.name)

	fmt.Println("---------------------")
	u1 := &u
	u2 := &u1

	Modify2(u2)
	fmt.Println(u)
}

type User struct {
	name string
}

func Modify(d *User) {
	d.name = "Milin-New"
}

func Modify1(d *User) {
	d = &User{"Milin_NEWNEW"}
	fmt.Println(d.name)
}

func Modify2(d **User) {
	*d = &User{"Milin_Modify2"}
}

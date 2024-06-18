package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Println("Salving user", u.name)
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user := user{"Lim", 30}
	fmt.Println(user)

	user.save()
	user.birthday()
	fmt.Println(user.age)
}

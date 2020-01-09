package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var std [10]student
	std[0] = student{"kao", 18, "kao_game@hotmail.co.th"}

	fmt.Println(std[0])
	fmt.Println(std[0].name)
	fmt.Println(std[0].age)
	fmt.Println(std[0].email)
}

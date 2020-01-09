package main

import "fmt"

func main() {
	var studentName [10]string
	var studentAge [10]int
	var studentEmail [10]string

	studentName[0] = "kao"
	studentName[1] = "kao1"
	studentName[2] = "kao2"
	studentName[3] = "kao3"
	studentAge[0] = 18
	studentEmail[0] = "kao_game@hotmail.co.th"

	fmt.Println(studentName[0], studentName[1], studentName[2], studentName[3], studentAge[0], studentEmail[0])
}

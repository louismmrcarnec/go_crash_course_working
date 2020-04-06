package main

import "fmt"

// in go when create a var it needs to be used


func main(){
	var age int32 = 27
	const isCool bool = true
	name := "Brad"

	name1, email1 := "Louis", "louis@gmail.com"

	fmt.Println(name, age, isCool)
	fmt.Println(name1, email1)
}
package main

import "fmt"

func main() {
	fmt.Println("Learning Programming Languages")
	fmt.Println("go-basics")

	var i int = 42
	var f float32 = 3.14

	firstName := "Stephen"
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(firstName)

	i1, i2 := 1, 42
	fmt.Println(i1)
	fmt.Println(i2)

	var name *string = new(string)
	*name = "Stephen"
	fmt.Println(name)
	fmt.Println(*name)

	ptr := &firstName
	fmt.Println(ptr, *ptr, firstName)
	*ptr = "Smithy"
	fmt.Println(ptr, *ptr, firstName)

}

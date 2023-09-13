package main

import "fmt"

func main() {
	age := 20
	name := "joe"

	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the str is: ", str)

}

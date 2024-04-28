package main

import "fmt"

func main() {
	name := "Sean"
	myInt := 10
	myFloat := 10.0

	fmt.Printf("hello my name is %s my int is %d my float is %f\n", name, myInt, myFloat)
	fmt.Printf("Hello, it's %s`s world!\n", name)

	var myFriendsName string
	var myBool bool
	var myOtherInt int

	myFriendsName = "Prime"

	fmt.Printf("my other friends name %s my bool %t and my other int %d\n", myFriendsName, myBool, myOtherInt)
}

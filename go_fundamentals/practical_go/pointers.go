package main

import "fmt"

func main(){
	x := 15

	a := &x 			//memory adress of x
	fmt.Println(a)

	b := *a				//Read through a (see what's on that memory adress)
	fmt.Println(b)

	*a = 5				//We also can declare values using this structure
	fmt.Println(*a**a)
	fmt.Println(x)
}
package main

import "fmt"

func main() {

	x := 10
	var p *int
	p = &x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 20
	fmt.Println(x)
}

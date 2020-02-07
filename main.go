package main

import "fmt"

func main() {
	var name string
	fmt.Print("What's your name?:")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello World %s", name )

}

package main

import "fmt"

func main() {
	var l float32
	fmt.Println("Area de un cuadrado")
	fmt.Println("Capture el lado: ")
	fmt.Scanf("%f", &l)
	area := l * l
	fmt.Println("El area es de ", area)
}

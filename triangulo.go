package main

import "fmt"

func main() {
	var b float32
	var h float32
	fmt.Println("Area de un triangulo")
	fmt.Println("Capture la base")
	fmt.Scan(&b)
	fmt.Println("Capture la altura")
	fmt.Scan(&h)
	area := (b * h) / 2
	fmt.Println("El area es de ", area)
}

package main

import (
	"fmt"
)

func main() {
	var r float32
	var pi float32 = 3.14159265359

	fmt.Println("Area de un circulo")
	fmt.Println("Ingrese el radio del circulo:")
	fmt.Scanf("%f", &r)
	area := pi * (r * r)
	fmt.Println("El area del circulo es ", area)
}

package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Введите 2 числа:")
	fmt.Scanf("%d %d", &a, &b)
	fmt.Printf("Произведение: %d\n", a*b)
}

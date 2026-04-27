package main

import (
	"fmt"
	"math"
)

func main() {
	// var greeting string = "Koda"
	// var a float32
	// var b int

	// d := true
	// isSame := greeting == strconv.Itoa(b)

	// fmt.Println(greeting)
	// fmt.Printf("%s\n", greeting)
	// fmt.Printf("%.2f\n", a)
	// fmt.Printf("%t\n", d)
	// fmt.Printf("%t\n", isSame)
	// fmt.Println("Hello World!")
	// hello()
	// greet("dwiki")
	// fmt.Println(multiplication(4, 2))
	fmt.Printf("Luas = %.2f\n", hitungLuas(5))
	fmt.Printf("Keliling = %.2f\n", hitungKeliling(5))
	luas, keliling := luasKeliling(5)
	fmt.Printf("Luas = %.2f\nKeliling = %.2f\n", luas, keliling)
}

// func hello() {
// 	fmt.Println("hello")
// }

// func greet(name string) {
// 	fmt.Printf("hello %s\n", name)
// }

// func multiplication(num1 int, num2 int) int {
// 	return num1 * num2
// }

func hitungLuas(r float32) float32 {
	return math.Pi * r * r
}

func hitungKeliling(r float32) float32 {
	return 2 * math.Pi * r
}

func luasKeliling(r float32) (float32, float32) {
	return hitungLuas(r), hitungKeliling(r)
}

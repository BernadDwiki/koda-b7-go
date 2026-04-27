package main

import (
	"fmt"

	"github.com/bernaddwiki/koda-b7-go/internals/minitask1"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask2"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask3"
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
	// fmt.Println(printToN(5))

	// var ages [5]int = [5]int{15, 15, 16, 17, 20}
	// fmt.Println(ages)

	// num := []int{50, 75, 66, 20, 32, 90}
	// sisipAngka(88, num)
	// num := make([]int, 0, 8)
	// num = append(num, 50, 75, 66, 20, 32, 90)
	// num = sisipAngka(88, num)
	// fmt.Println(num)

	fmt.Printf("Luas = %.2f\n", minitask1.HitungLuas(5))
	fmt.Printf("Keliling = %.2f\n", minitask1.HitungKeliling(5))
	luas, keliling := minitask1.LuasKeliling(5)
	fmt.Printf("Luas = %.2f\nKeliling = %.2f\n", luas, keliling)
	minitask2.SegitigaSiku(5)
	minitask3.SisipAngkaSlice(88)

	dwiki := user{
		name:        "Dwiki",
		image:       "kdsjgakdgakd",
		email:       "dwiki.adicitra6@gmail.com",
		age:         22,
		phoneNumber: "081228669806",
		isMarried:   false,
		education: []education{{
			name:  "Sanata Dharma University",
			major: "Informatics",
		},
		},
	}

	fmt.Println(dwiki.name)
	fmt.Println(dwiki.image)
	fmt.Println(dwiki.email)
	fmt.Println(dwiki.age)
	fmt.Println(dwiki.phoneNumber)
	fmt.Println(dwiki.isMarried)
	fmt.Println(dwiki.education[0].name)
	fmt.Println(dwiki.education[0].major)

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

// func printToN(n int) {
// 	for i := 0; i <= n; i++ {
// 		fmt.Println(i)
// 	}

// 	for i := range n {
// 		fmt.Println(i)
// 	}
// }

// func sisipAngka(angka int, num []int) []int {
// 	num = slices.Insert(num, 3, angka)
// 	return num
// }

type education struct {
	name  string
	major string
}

type user struct {
	name        string
	image       string
	email       string
	age         int
	phoneNumber string
	isMarried   bool
	education   []education
}

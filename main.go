package main

import (
	"fmt"

	"github.com/bernaddwiki/koda-b7-go/internals/minitask1"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask2"
	"github.com/bernaddwiki/koda-b7-go/internals/minitask3"
)

func main() {
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

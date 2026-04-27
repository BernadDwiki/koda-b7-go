package minitask3

import (
	"fmt"
	"slices"
)

func SisipAngkaSlice(angka int) {
	data := []int{50, 75, 66, 20, 32, 90}

	data = slices.Insert(data, 3, angka)

	fmt.Println(data)

	for i, v := range data {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

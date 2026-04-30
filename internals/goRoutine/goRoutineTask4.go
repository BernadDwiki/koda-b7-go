// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	chn := make(chan int, 1)
// 	angkaN := 10
// 	chn <- angkaN
// 	wg.Add(1)
// 	go generateAngka(chn, &wg)
// 	close(chn)
// 	wg.Wait()
// }

// func generateAngka(n chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	chn := make(chan []int, 1)
// 	angka := []int{}

// 	limit := <-n
// 	for i := 1; i <= limit; i++ {
// 		angka = append(angka, i)
// 	}
// 	fmt.Println("Angka yang dihasilkan:", angka)

// 	chn <- angka

// 	wg.Add(1)
// 	go angkaGenap(chn, wg)
// 	close(chn)
// }

// func angkaGenap(angkaChn chan []int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	chn := make(chan []int, 1)
// 	genap := []int{}

// 	sliceAngka := <-angkaChn
// 	for _, angka := range sliceAngka {
// 		if angka%2 == 0 {
// 			genap = append(genap, angka)
// 		}
// 	}
// 	fmt.Println("Angka genap yang dihasilkan:", genap)
// 	chn <- genap
// 	close(chn)
// 	wg.Add(1)
// 	go kuadratAngka(chn, wg)
// }

// func kuadratAngka(angkaChn chan []int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	kuadrat := []int{}
// 	sliceGenap := <-angkaChn
// 	for _, angka := range sliceGenap {
// 		kuadrat = append(kuadrat, angka*angka)
// 	}
// 	fmt.Println("Kuadrat dari angka genap:", kuadrat)
// }

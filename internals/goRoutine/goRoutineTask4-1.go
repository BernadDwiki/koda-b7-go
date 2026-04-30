package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	chanInput := make(chan int)
	chanGenap := make(chan []int)
	chanKuadrat := make(chan []int)

	wg.Add(3)
	go generateAngka(chanInput, chanGenap, &wg)
	go angkaGenap(chanGenap, chanKuadrat, &wg)
	go kuadratAngka(chanKuadrat, &wg)

	chanInput <- 10
	close(chanInput)

	wg.Wait()
}

func generateAngka(in chan int, out chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	limit := <-in
	angka := []int{}
	for i := 1; i <= limit; i++ {
		angka = append(angka, i)
	}
	fmt.Println("Angka yang dihasilkan:", angka)
	out <- angka
}

func angkaGenap(in chan []int, out chan []int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	sliceAngka := <-in
	genap := []int{}
	for _, v := range sliceAngka {
		if v%2 == 0 {
			genap = append(genap, v)
		}
	}
	fmt.Println("Angka genap yang dihasilkan:", genap)
	out <- genap
}

func kuadratAngka(in chan []int, wg *sync.WaitGroup) {
	defer wg.Done()

	sliceGenap := <-in
	kuadrat := []int{}
	for _, v := range sliceGenap {
		kuadrat = append(kuadrat, v*v)
	}
	fmt.Println("Kuadrat dari angka genap yang dihasilkan:", kuadrat)
}

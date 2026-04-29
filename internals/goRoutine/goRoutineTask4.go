package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	chn := make(chan int, 1)
	angkaN := 10
	chn <- angkaN
	wg.Add(1)
	go generateAngka(chn, &wg)
	close(chn)
	wg.Wait()

}

func generateAngka(n chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	chn := make(chan int, 1)
	for i := range n {
		chn <- i
	}
	wg.Add(1)
	go angkaGenap(chn, wg)
	close(chn)
}

func angkaGenap(angkaChn chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	chn := make(chan int, 1)
	for angka := range angkaChn {
		if angka%2 == 0 {
			chn <- angka
		}
	}
	wg.Add(1)
	go kuadratAngka(chn, wg)
	close(chn)
}

func kuadratAngka(angkaChn chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for angka := range angkaChn {
		kuadrat := angka * angka
		fmt.Println(kuadrat)
	}
}

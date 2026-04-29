// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	coffeeShop()
// }

// func coffeeShop() {
// 	var wg sync.WaitGroup
// 	var chn chan string
// 	// chn = make(chan string) // unbuffered channel
// 	chn = make(chan string, 3) // buffered channel dengan kapasitas 3
// 	coffees := []string{"Espresso", "Latte", "Cappucino", "Americano", "Mocha", "Macchiato"}

// 	baristaCount := 3
// 	for range baristaCount {
// 		wg.Add(1)
// 		go barista(chn, &wg)
// 	}

// 	for _, coffee := range coffees {
// 		chn <- coffee
// 		fmt.Printf("Pesanan diterima: %s\n", coffee)
// 		time.Sleep(1 * time.Millisecond)
// 	}
// 	close(chn)
// 	wg.Wait()
// 	fmt.Println("Semua kopi sudah dibuat!")
// 	// wg.Add(3)
// 	// go barista(coffees[i+0], &wg)
// 	// go barista(coffees[i+1], &wg)
// 	// go barista(coffees[i+2], &wg)
// 	// wg.Wait()

// }

// func barista(coffeeChn chan string, wg *sync.WaitGroup) {
// 	defer func() {
// 		fmt.Println("Barista selesai bekerja")
// 		wg.Done()
// 	}()
// 	for coffee := range coffeeChn {
// 		time.Sleep(1 * time.Millisecond)
// 		fmt.Printf("Barista membuat kopi: %s\n", coffee)
// 		time.Sleep(1 * time.Second)
// 		fmt.Printf("Barista selesai membuat kopi: %s\n", coffee)
// 	}
// }

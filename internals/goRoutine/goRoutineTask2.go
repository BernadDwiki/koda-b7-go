// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	dailyRoutine()
// }

// func dailyRoutine() {
// 	var wg sync.WaitGroup
// 	wg.Add(4)
// 	go mandi(&wg)
// 	go buatKopi(&wg)
// 	go sarapan(&wg)
// 	go rapikanKamar(&wg)
// 	wg.Wait()
// 	fmt.Println("Berangkat kerja")
// }

// func mandi(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Mulai mandi")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Selesai mandi")
// }

// func buatKopi(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Mulai membuat kopi")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Selesai membuat kopi")
// }

// func sarapan(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Mulai sarapan")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Selesai sarapan")
// }

// func rapikanKamar(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Mulai merapikan kamar")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("Selesai merapikan kamar")
// }

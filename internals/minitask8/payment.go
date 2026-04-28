package minitask8

import (
	"errors"
	"fmt"
)

type Pembayaran interface {
	Transfer(jumlah []float64) error
}

type Bank struct {
}

func (b Bank) Transfer(jumlah []float64) error {
	total := sum(jumlah)
	fmt.Println("Pembayaran Via Bank Berhasil")
	fmt.Println("Total :", total)
	return nil
}

type OnlineBank struct {
}

func (o OnlineBank) Transfer(jumlah []float64) error {
	total := sum(jumlah)
	fmt.Println("Pembayaran Via Online Bank Berhasil")
	fmt.Println("Total :", total)
	return nil
}

type MockPayment struct {
	History []string
}

func (m *MockPayment) Transfer(jumlah []float64) error {
	for _, v := range jumlah {
		if v < 0 {
			return errors.New("harga tidak boleh negatif")
		}
	}
	total := sum(jumlah)
	m.History = append(m.History, fmt.Sprintf("Transfer: %.2f", total))
	return nil
}

func sum(prices []float64) float64 {
	var total float64
	for _, value := range prices {
		total += value
	}
	return total
}

package main

import (
	"fmt"
)

type Deret struct {
	Limit int
}

func (d *Deret) AngkaPrima() {
	hasil := []int{}
	for i := 2; i <= d.Limit; i++ {
		numPrima := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				numPrima = false
				break
			}
		}
		if numPrima {
			hasil = append(hasil, i)
		}
	}
	fmt.Println("Angka Prima:", hasil)
}

func (d *Deret) AngkaGanjil() {
	hasil := []int{}
	for i := 1; i <= d.Limit; i += 2 {
		hasil = append(hasil, i)
	}
	fmt.Println("Angka Ganjil:", hasil)
}

func (d *Deret) AngkaGenap() {
	hasil := []int{}
	for i := 2; i <= d.Limit; i += 2 {
		hasil = append(hasil, i)
	}
	fmt.Println("Angka Genap:", hasil)
}

func (d *Deret) AngkaFibo() {
	hasil := []int{0, 1}
	for i := 2; hasil[len(hasil)-1]+hasil[len(hasil)-2] <= d.Limit; i++ {
		hasil = append(hasil, hasil[len(hasil)-1]+hasil[len(hasil)-2])
	}
	fmt.Println("Angka Fibonacci:", hasil)
}

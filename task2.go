package main

import (
	"fmt"
	"sync"
)

func GanjilGenap(Limit int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	hasil := []int{0, 1}
	for i := 2; hasil[len(hasil)-1]+hasil[len(hasil)-2] <= Limit; i++ {
		nextProses := hasil[len(hasil)-1] + hasil[len(hasil)-2]
		hasil = append(hasil, nextProses)

		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			if nextProses%2 == 0 {
				ch <- value
				fmt.Println("Angka Genap", nextProses)
			}
			if nextProses%2 != 0 {
				ch <- value
				fmt.Println("Angka Ganjil", nextProses)
			}
		}(nextProses)
	}
}

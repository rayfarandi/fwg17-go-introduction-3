package main

import (
	"sync"
)

func main() {
	// task1
	// deret := Deret{Limit: 40}

	// var wg sync.WaitGroup // harus di matikan jika salah satu acctive
	// wg.Add(4)

	// go func() {
	// 	defer wg.Done()
	// 	deret.AngkaPrima()
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	deret.AngkaGanjil()
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	deret.AngkaGenap()
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	deret.AngkaFibo()
	// }()

	// wg.Wait()

	// task2

	// GanjilGenap(40)
	var wg sync.WaitGroup

	ch := make(chan int, 10)

	wg.Add(1)
	go GanjilGenap(40, &wg, ch)

	wg.Wait()
	close(ch)

}

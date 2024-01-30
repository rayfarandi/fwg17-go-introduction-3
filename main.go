package main

import (
	"fmt"
	"sync"
)

func main() {
	// task1 (jika ingin menjalankan maka hilangkan coment)
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

	// task2 (jika ingin menjalankan maka hilangkan coment)

	// GanjilGenap(40)
	// var wg sync.WaitGroup

	// ch := make(chan int, 10)

	// wg.Add(1)
	// go GanjilGenap(40, &wg, ch)

	// wg.Wait()
	// close(ch)

	// task 3

	var workingDays = make([]string, 0)
	var mu sync.Mutex
	var wg sync.WaitGroup // harus di matikan jika salah satu

	ch := make(chan string)

	processDays(ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		mu.Lock()
		workingDays = append(workingDays, result)
		mu.Unlock()
	}

	// Menampilkan semua data hari
	fmt.Println("Semua Hari:", days)
	fmt.Println("Hari Kerja:", workingDays)

}
